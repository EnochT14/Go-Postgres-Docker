package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	host          = "192.168.100.10" // Host IP
	port          = 5432
	user          = "postgres" // PostgreSQL username
	password      = "password" // PostgreSQL password
	dbname        = "postgres" // Database name
	listenAddress = ":8080"
)

func main() {
	// Create a database connection
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// CORS with allowed origins
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	router.Use(corsMiddleware.Handler)

	// HTTP endpoint to receive sensor data
	router.HandleFunc("/collect-sensor-data", func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON data from the request body
		var sensorData struct {
			Temperature float64 `json:"temperature"`
			Humidity    float64 `json:"humidity"`
			Pressure    float64 `json:"pressure"`
			CO2PPM      int     `json:"co2_ppm"`
			TVOCPpb     int     `json:"tvoc_ppb"`
		}

		if err := json.NewDecoder(r.Body).Decode(&sensorData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		timestamp := time.Now()

		insertStatement := `
            INSERT INTO sensor_data (timestamp, temperature, humidity, pressure, co2_ppm, tvoc_ppb)
            VALUES ($1, $2, $3, $4, $5, $6)
        `
		_, err := db.Exec(insertStatement, timestamp, sensorData.Temperature, sensorData.Humidity, sensorData.Pressure, sensorData.CO2PPM, sensorData.TVOCPpb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Sensor data inserted successfully.")
		w.WriteHeader(http.StatusNoContent)
	})

	log.Printf("Server listening on %s...", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, router))
}
