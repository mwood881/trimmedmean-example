package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mwood881/trimmean"
)

func main() {
	// Open the integers CSV file
	intFile, err := os.Open("integers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer intFile.Close()

	// Read the integers CSV file
	intReader := csv.NewReader(intFile)
	intRecords, err := intReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Skip the first row (header)
	if len(intRecords) > 0 {
		intRecords = intRecords[1:]
	}

	// Parse the integers from the CSV
	var integers []float64
	for _, record := range intRecords {
		for _, value := range record {
			// Log the value to debug
			fmt.Println("Parsing integer value:", value)

			// Convert string to float64 and skip invalid entries
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Skipping invalid value '%s': %v\n", value, err)
				continue // Skip invalid value
			}
			integers = append(integers, num)
		}
	}

	floatFile, err := os.Open("floats.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer floatFile.Close()

	// Read the floats CSV file
	floatReader := csv.NewReader(floatFile)
	floatRecords, err := floatReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Skip the first row (header)
	if len(floatRecords) > 0 {
		floatRecords = floatRecords[1:]
	}

	// Parse the floats from the CSV
	var floats []float64
	for _, record := range floatRecords {
		for _, value := range record {
			// Log the value to debug
			fmt.Println("Parsing float value:", value)

			// Convert string to float64 and skip invalid entries
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Skipping invalid value '%s': %v\n", value, err)
				continue // Skip invalid value
			}
			floats = append(floats, num)
		}
	}

	fmt.Println("Computing trimmed mean for integers...")
	mean, err := trimmean.TrimmedMean(integers, 0.05, 0.05)
	if err != nil {
		log.Fatalf("Error computing trimmed mean for integers: %v", err)
	}
	fmt.Printf("Trimmed mean for integers: %.2f\n", mean)

	// Example with floats
	fmt.Println("Computing trimmed mean for floats...")
	mean, err = trimmean.TrimmedMean(floats, 0.05, 0.05)
	if err != nil {
		log.Fatalf("Error computing trimmed mean for floats: %v", err)
	}
	fmt.Printf("Trimmed mean for floats: %.2f\n", mean)
}
