package usecase

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/Dagosu/BookingApp/booking-api/app/domain"
	dt "github.com/Dagosu/BookingApp/datatypes"
	"go.mongodb.org/mongo-driver/mongo"
)

type userFlightsMappingUsecase struct {
	ufmr domain.UserFlightsMappingRepository
	ur   domain.UserRepository
	fr   domain.FlightRepository
}

func newUserFlightsMappingUsecase(ufmr domain.UserFlightsMappingRepository, ur domain.UserRepository, fr domain.FlightRepository) *userFlightsMappingUsecase {
	return &userFlightsMappingUsecase{ufmr, ur, fr}
}

func (ufmu *userFlightsMappingUsecase) PurchaseFlight(ctx context.Context, userId, flightId string) (*dt.Flight, error) {
	user, err := ufmu.ur.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	flight, err := ufmu.fr.Get(ctx, flightId)
	if err != nil {
		return nil, err
	}

	userFlights, err := ufmu.ufmr.GetByUser(ctx, user.GetId())
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Check for purchased flights
	for _, f := range userFlights.GetPurchasedFlights() {
		if f.GetId() == flightId {
			return nil, fmt.Errorf("You already purchased this flight!")
		}
	}

	if err == mongo.ErrNoDocuments {
		_, err := ufmu.ufmr.Create(ctx, user, flight)
		if err != nil {
			return nil, err
		}

		return flight, nil
	}

	err = ufmu.ufmr.AddPurchasedFlight(ctx, userFlights.GetId(), flight)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (ufmu *userFlightsMappingUsecase) FavoriteFlight(ctx context.Context, userId, flightId string) (*dt.Flight, error) {
	user, err := ufmu.ur.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	flight, err := ufmu.fr.Get(ctx, flightId)
	if err != nil {
		return nil, err
	}

	userFlights, err := ufmu.ufmr.GetByUser(ctx, user.GetId())
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		_, err := ufmu.ufmr.Create(ctx, user, flight)
		if err != nil {
			return nil, err
		}

		return flight, nil
	}

	err = ufmu.ufmr.AddFavoritedFlight(ctx, userFlights.GetId(), flight)
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (ufmu *userFlightsMappingUsecase) GetPurchasedFlights(ctx context.Context, userId string) ([]*dt.Flight, error) {
	userFlights, err := ufmu.ufmr.GetByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	flights := []*dt.Flight{}
	for _, f := range userFlights.PurchasedFlights {
		flights = append(flights, f)
	}

	return flights, nil
}

func (ufmu *userFlightsMappingUsecase) GetFavoritedFlights(ctx context.Context, userId string) ([]*dt.Flight, error) {
	userFlights, err := ufmu.ufmr.GetByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	flights := []*dt.Flight{}
	for _, f := range userFlights.FavoritedFlights {
		flights = append(flights, f)
	}

	return flights, nil
}

func (ufmu *userFlightsMappingUsecase) CheckFlightPurchase(ctx context.Context, flightId, userId string) (*dt.Flight, error) {
	flight, err := ufmu.ufmr.CheckFlightPurchase(ctx, flightId, userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return flight, nil
}

func (ufmu *userFlightsMappingUsecase) RecommendFlight(ctx context.Context, userId string) ([]*dt.Flight, error) {
	userFlights, err := ufmu.ufmr.GetByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	futureFlights, err := ufmu.fr.GetFutureFlights(ctx)
	if err != nil {
		return nil, err
	}

	// Exclude flights that the user has already purchased or favorited
	existingFlights := make(map[string]bool)
	for _, flight := range append(userFlights.PurchasedFlights, userFlights.FavoritedFlights...) {
		existingFlights[flight.GetId()] = true
	}

	flightsToRecommend := make([]*dt.Flight, 0)
	for _, flight := range futureFlights {
		if !existingFlights[flight.GetId()] {
			flightsToRecommend = append(flightsToRecommend, flight)
		}
	}

	// Build user profile
	userProfile := buildUserProfile(userFlights)

	// Score remaining flights
	scores := make(map[*dt.Flight]float64)
	for _, flight := range flightsToRecommend {
		scores[flight] = calculateScore(userProfile, flight)
	}
	sort.Slice(flightsToRecommend, func(i, j int) bool {
		return scores[flightsToRecommend[i]] > scores[flightsToRecommend[j]]
	})

	return flightsToRecommend[:3], nil
}

func buildUserProfile(userFlights *dt.UserFlightsMapping) map[string]float64 {
	userProfile := make(map[string]float64)
	total := 0.0
	for _, flight := range append(userFlights.GetPurchasedFlights(), userFlights.GetFavoritedFlights()...) {
		userProfile[flight.GetDeparture()]++
		userProfile[flight.GetArrival()]++
		userProfile[timeOfDay(flight.GetDepartureTime().AsTime())]++
		userProfile[timeOfDay(flight.GetArrivalTime().AsTime())]++
		userProfile[flight.GetAirline()]++
		userProfile[priceBucket(float64(flight.GetPrice()))]++
		total += 6
	}
	// Normalize counts
	for key, count := range userProfile {
		userProfile[key] = count / total
	}

	return userProfile
}

func timeOfDay(t time.Time) string {
	hour := t.Hour()
	switch {
	case hour < 6:
		return "night"
	case hour < 12:
		return "morning"
	case hour < 18:
		return "afternoon"
	default:
		return "evening"
	}
}

func priceBucket(price float64) string {
	switch {
	case price < 100:
		return "cheap"
	case price < 200:
		return "moderate"
	default:
		return "expensive"
	}
}

func calculateScore(userProfile map[string]float64, flight *dt.Flight) float64 {
	score := 0.0
	// Weights for different features
	weights := map[string]float64{
		flight.Departure:                         0.8,
		flight.Arrival:                           1,
		timeOfDay(flight.DepartureTime.AsTime()): 0.2,
		timeOfDay(flight.ArrivalTime.AsTime()):   0.2,
		flight.Arrival:                           0.5,
		priceBucket(float64(flight.Price)):       0.8,
	}
	for feature, weight := range weights {
		if count, ok := userProfile[feature]; ok {
			score += weight * count
		}
	}
	return score
}
