package grpcserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// RegisterCallback function type for grpc init callback.
type RegisterCallback func(*grpc.Server)

// GRPCServer contains the required settings for running a grpcServer
type GRPCServer struct {
	grpcPort           int
	metricsPort        int
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	enableReflection   bool
	tcpListener        net.Listener
	httpServer         *http.Server
	grpcServer         *grpc.Server
	gracefullyStopped  bool
}

// isValidPortNum tests that the argument is a valid, non-zero port number.
func isValidPortNum(p int) error {
	low := 1
	high := 65535

	if low <= p && p <= high {
		return nil
	}

	return fmt.Errorf("Port must be between %d and %d, inclusive", low, high)
}

// WithReflection adds the ability to enable the grpc reflection.
func WithReflection(b bool) func(*GRPCServer) error {
	return func(s *GRPCServer) error {
		s.enableReflection = b

		return nil
	}
}

// WithUnaryInterceptors activates the passed slice
// of UnaryServerInterceptor(s) on the provided server
func WithUnaryInterceptors(i []grpc.UnaryServerInterceptor) func(*GRPCServer) error {
	return func(s *GRPCServer) error {
		s.unaryInterceptors = i

		return nil
	}
}

// WithStreamInterceptors activates the passed slice
// of StreamServerInterceptor(s) on the provided server
func WithStreamInterceptors(i []grpc.StreamServerInterceptor) func(*GRPCServer) error {
	return func(s *GRPCServer) error {
		s.streamInterceptors = i

		return nil
	}
}

// WithGRPCPort adds ability to override default
// port on which the GRPC services are served
func WithGRPCPort(port int) func(*GRPCServer) error {
	return func(s *GRPCServer) error {
		if err := isValidPortNum(port); err != nil {
			return err
		}

		s.grpcPort = port

		return nil
	}
}

// WithMetricsPort adds ability to override default
// port on which the HTTP prometheus metrics are served
func WithMetricsPort(port int) func(*GRPCServer) error {
	return func(s *GRPCServer) error {
		if err := isValidPortNum(port); err != nil {
			return err
		}

		s.metricsPort = port

		return nil
	}
}

// New returns a grpcserver struct
func New(options ...func(*GRPCServer) error) (*GRPCServer, error) {
	s := &GRPCServer{
		grpcPort:    50050,
		metricsPort: 9090,
	}

	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, fmt.Errorf("New grpcserver option error %#v", err)
		}
	}

	return s, nil
}

// Init starts the grpc stream, it receives a callback function, an array of unary stream interceptors and an array of stream server interceptors.
func (s *GRPCServer) Init(callback RegisterCallback) {
	log.Println("Starting the GRPC server")

	grpcMetrics := grpc_prometheus.NewServerMetrics()

	// we're defering the start of the GRPC server listening
	// so that we first register all required services
	defer s.startGRPCServer(grpcMetrics)()

	callback(s.grpcServer)

	if s.enableReflection {
		// Register reflection service on gRPC server.
		reflection.Register(s.grpcServer)
	}

	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(s.grpcServer)

	s.startPrometheusHTTPServer(grpcMetrics)
}

func (s *GRPCServer) startPrometheusHTTPServer(grpcMetrics *grpc_prometheus.ServerMetrics) {
	// Register standard server metrics to registry. (can add any customized metrics )
	// Create a metrics registry.
	reg := prometheus.NewRegistry()
	reg.MustRegister(grpcMetrics)

	s.httpServer = &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    fmt.Sprintf("0.0.0.0:%d", s.metricsPort),
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if s.gracefullyStopped {
				return
			}

			log.Println("Warning: Unable to start the prometheus http server:", err)
		}
	}()

	log.Println("Prometheus metrics exported on port", s.metricsPort)
}

// Initializes grpcServer, returns function that trigger the grpcServer serving
func (s *GRPCServer) startGRPCServer(grpcMetrics *grpc_prometheus.ServerMetrics) func() {
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				otelgrpc.StreamServerInterceptor(),
				grpcMetrics.StreamServerInterceptor(),
				grpc_middleware.ChainStreamServer(s.streamInterceptors...),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				otelgrpc.UnaryServerInterceptor(),
				grpcMetrics.UnaryServerInterceptor(),
				grpc_middleware.ChainUnaryServer(s.unaryInterceptors...),
			),
		),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
	}

	s.grpcServer = grpc.NewServer(opts...)

	log.Println("Starting server on port", s.grpcPort)
	var err error
	s.tcpListener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		log.Fatalf("Unable to listen on port %v: %v", s.grpcPort, err)
	}

	return func() {
		go func() {
			if err := s.grpcServer.Serve(s.tcpListener); err != nil {
				if s.gracefullyStopped {
					return
				}

				log.Fatalf("Failed to serve: %v", err)
			}
		}()

		log.Println("Server successfully started on port", s.grpcPort)
	}
}

// Close closes the opened grpcServer and tcpListener
func (s *GRPCServer) Close() {
	log.Println("Stopping the GRPC server")
	log.Println("Stopping the HTTP server")

	s.gracefullyStopped = true
	s.grpcServer.GracefulStop()
	s.httpServer.Close()
	s.tcpListener.Close()
}
