package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	dt "github.com/Dagosu/BookingApp/datatypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// Setup MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/booking-api"))
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	// Create a ticker to call updateFlightStatus every 5 minutes
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})

	fmt.Println("Starting sync api")

	go func() {
		for {
			select {
			case <-ticker.C:
				updateFlightStatus(client)
				fmt.Println("Flights updated!")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for a termination signal
	<-quit
}

func updateFlightStatus(client *mongo.Client) {
	// Get a handle for collection
	collection := client.Database("booking-api").Collection("flights")

	// Create a cursor for all documents in the collection
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Failed to find documents:", err)
		return
	}
	defer cur.Close(ctx)

	// Iterate over the cursor and update each document's status field
	for cur.Next(ctx) {
		// Decode the document into a map
		var doc map[string]bson.RawValue
		err := cur.Decode(&doc)
		if err != nil {
			fmt.Println("Failed to decode document:", err)
			continue
		}

		id := getStringField(doc, "_id")

		// Extract the departure_time and arrival_time fields
		departureTimeRaw, ok := doc["departure_time"]
		if !ok {
			fmt.Println("Missing departure_time field")
			continue
		}
		arrivalTimeRaw, ok := doc["arrival_time"]
		if !ok {
			fmt.Println("Missing arrival_time field")
			continue
		}

		// Convert the departure_time and arrival_time fields to timestamppb.Timestamp
		departureTime, err := extractTimestamp(departureTimeRaw)
		if err != nil {
			fmt.Println("Failed to extract departure_time:", err)
			continue
		}
		arrivalTime, err := extractTimestamp(arrivalTimeRaw)
		if err != nil {
			fmt.Println("Failed to extract arrival_time:", err)
			continue
		}

		// Create a Flight instance and set the extracted fields
		flight := &dt.Flight{
			Id:            id,
			DepartureTime: departureTime,
			ArrivalTime:   arrivalTime,
		}

		now := time.Now()
		var newStatus string

		if now.Before(flight.GetDepartureTime().AsTime()) {
			newStatus = "scheduled"
		} else if now.After(flight.GetArrivalTime().AsTime()) {
			newStatus = "arrived"
		} else {
			newStatus = "active"
		}

		update := bson.D{{"$set", bson.D{{"status", newStatus}}}}
		_, err = collection.UpdateOne(ctx, bson.D{{"_id", flight.GetId()}}, update)
		if err != nil {
			fmt.Println("Failed to update document:", err)
			continue
		}
	}

	if err := cur.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return
	}
}

func extractTimestamp(rawValue bson.RawValue) (*timestamppb.Timestamp, error) {
	data := rawValue.Value

	if len(data) != 8 {
		return nil, fmt.Errorf("invalid binary timestamp format")
	}

	// Convert the binary timestamp to a little-endian int64 value
	timestamp := int64(binary.LittleEndian.Uint64(data))

	seconds := timestamp / 1000
	nanos := (timestamp % 1000) * 1000000

	return &timestamppb.Timestamp{
		Seconds: seconds,
		Nanos:   int32(nanos),
	}, nil
}

func getStringField(doc map[string]bson.RawValue, fieldName string) string {
	rawValue, ok := doc[fieldName]
	if !ok {
		fmt.Println("Missing field:", fieldName)
		return ""
	}

	var fieldValue string
	err := rawValue.Unmarshal(&fieldValue)
	if err != nil {
		fmt.Println("Failed to extract field:", fieldName, "-", err)
		return ""
	}

	return fieldValue
}
