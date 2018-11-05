package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// Import dependencies packages
// Packages are imported via `go get packageName`

// Datum Declare data variables with types
type Datum struct {
	One   string
	Two   string
	Three string
}

// TODO: generic struct to handle n fields of type string, n declared by a param in the execution of the
// binary

// Main function where the code will be executed

// Example: http://www.cihanozhan.com/converting-csv-data-to-json-with-golang/
func main() {
	csvFile, err := os.Open("./myData.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
	// Get a reader for the csv
	reader := csv.NewReader(csvFile)
	// Mandatory option to handle double quotes
	reader.LazyQuotes = true

	csvData, csvErr := reader.ReadAll()

	if csvErr != nil {
		fmt.Println(csvErr)
		os.Exit(1)
	}
	fmt.Println(csvData)

	var datum Datum
	var data []Datum
	// TODO: get the generic struct and loop
	for _, each := range csvData {
		datum.One = each[0]
		datum.Two = each[1]
		datum.Three = each[2]

		data = append(data, datum)
	}
	fmt.Print("DATA")
	fmt.Println(data)

	// Convert to JSON
	jsonData, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1) // Exit with an error code
	}

	fmt.Print("JSON DATA")

	fmt.Println(string(jsonData))

	// Create the json file on the disk
	jsonFile, jsonFileErr := os.Create("./data.json")
	if err != nil {
		fmt.Println(jsonFileErr)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
}
