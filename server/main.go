package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	x "server/interfaces"
	xy "server/models"
	"server/service"
	"server/utils"
	"time"
)

var router x.RouterInterface

func init() {
	// Switch between GorillaMuxRouter and Nethttp based on requirement
	useGorilla := true // Change this flag to switch routers
	if useGorilla {
		router = x.NewGorillaMuxRouter()
	} else {
		router = x.NewNethttp()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fetch the data
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("We encounterd error ", err)
		return
	}
	// we need to Unmarshal it
	p1 := xy.Payload{}

	err = json.Unmarshal(res, &p1)
	if err != nil {
		fmt.Println("We Encountered Error in Unmarshal", err)
		return
	}
	// Connect to MongoDB and save the Payload
	mongoClient := utils.GetMongoClient() // Ensure this retrieves the singleton MongoDB client

	// Check MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mongoClient.Ping(ctx); err != nil {
		http.Error(w, "failed to connect to MongoDB", http.StatusInternalServerError)
		fmt.Println("Error pinging MongoDB:", err)
		return
	}

	// Get the collection where data will be stored
	collectionService, err := mongoClient.GetCollection("metricsDB", "payloads")
	if err != nil {
		http.Error(w, "failed to get MongoDB collection", http.StatusInternalServerError)
		fmt.Println("Error retrieving collection:", err)
		return
	}

	// Insert the payload into the collection
	insertResult, err := collectionService.InsertOne(ctx, p1)
	if err != nil {
		http.Error(w, "failed to insert payload into MongoDB", http.StatusInternalServerError)
		fmt.Println("Error inserting payload into MongoDB:", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payload successfully stored with ID: %v", insertResult)
	fmt.Println("Payload stored in MongoDB with ID:", insertResult)
	service.PrintMetrics(p1)

}

func main() {
	port := "8080"
	// router created

	router.HandleFunc(" ", handler)

	// Register handler
	router.HandleFunc("/", handler)

	// Listen and Serve at port number
	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("We Encountered Error in Listen And Serve:", err)
		return
	}
}
