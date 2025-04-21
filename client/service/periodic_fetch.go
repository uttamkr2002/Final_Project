package service

import (
	"client/infrastructure"
	//infrastructure "client/db"
	"sync"

	models "client/modelsWithInterface"
	"fmt"
	"log"
	"time"
)

// Periodic function to fetch and store metrics every 60 seconds
func PeriodicFetchMetrics(wg *sync.WaitGroup) {
	fmt.Println("We are In PeriodicFetchMetrics")
	ticker := time.NewTicker(60 * time.Second) // Run every 60 seconds
	defer ticker.Stop()
	defer wg.Done()

	for {

		<-ticker.C
		// Collect metrics
		// create an instance of payload(variable)
		temp := models.Payload{}        // this struct is empty now
		temp.CollectMetricsforPayload() // make it full

		// Print collected metrics
		PrintMetrics(temp) // payload to temp

		dbClient, err := infrastructure.InitDb()
		if err != nil {
			log.Fatalf(" Error initializing database: %v", err)
			return
		}

		// Store the fetched metrics in PostgreSQL
		if _, err := infrastructure.InsertMetrics(dbClient, temp); err != nil { // add another parameter sqlClient
			log.Println("Error storing metrics in DB:", err)
		}

	}
}
