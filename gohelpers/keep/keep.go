package keep

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	ghc "github.com/Dagosu/BookingApp/gohelpers/config"
)

// WaitForSignal keeps the application running until CTRL+C is pressed or system terminates it
func WaitForSignal() {
	log.Println("Server running", ghc.C.ServiceName)

	// Create a channel to receive OS signals
	c := make(chan os.Signal, 1)

	// catch CTRL+C and system termination signal
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed
	// And our main routine keeps running
	// If the main routine were to shutdown so would the child routine that is Serving the server
	<-c

	// allow process to be killed on a second CTRL+C
	signal.Stop(c)

	log.Println("Received termination signal")
}
