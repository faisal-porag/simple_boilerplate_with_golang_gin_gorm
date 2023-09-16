package main

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"net/http"
)

type Person struct {
	Name    string  `csv:"Name"`
	Age     int     `csv:"Age"`
	City    string  `csv:"City"`
	Contact Contact `csv:"contact"`
}

type Contact struct {
	Address string `csv:"address"`
	Phone   string `csv:"phone"`
}

func generateCSVData() []*Person {
	// Create sample CSV data
	data := []*Person{
		{Name: "John Doe", Age: 30, City: "New York", Contact: Contact{
			Address: "Dhaka",
			Phone:   "5787689",
		}},
		{Name: "Jane Smith", Age: 25, City: "Los Angeles", Contact: Contact{
			Address: "Khulna",
			Phone:   "678794",
		}},
		{Name: "Bob Johnson", Age: 35, City: "Chicago", Contact: Contact{
			Address: "Kishoreganj",
			Phone:   "987784",
		}},
		{Name: "Faisal", Age: 29, City: "Dhaka", Contact: Contact{
			Address: "Dhaka",
			Phone:   "79089656475",
		}},
	}

	return data
}

func generateCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.csv")

	// Create a new CSV writer
	writer := csv.NewWriter(w)
	defer writer.Flush()

	data := generateCSVData()

	// Use gocsv to write data to the CSV file
	if err := gocsv.MarshalCSV(data, writer); err != nil {
		http.Error(w, "Error writing to CSV", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/generate-csv", generateCSV)
	_ = http.ListenAndServe(":8280", nil)
}
