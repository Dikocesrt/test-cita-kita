package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type CountryCityCount struct {
	CountryName string `json:"country_name"`
	CityCount   int    `json:"city_count"`
}

func main() {
	csvURL := "https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/refs/heads/master/csv/cities.csv"

	resp, err := http.Get(csvURL)
	if err != nil {
		fmt.Println("failed download file:", err)
		return
	}
	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)
	r.LazyQuotes = true

	header, err := r.Read()
	if err != nil {
		fmt.Println("unable to read csv header:", err)
		return
	}

	idxCountry := -1
	for i, col := range header {
		if strings.ToLower(col) == "country_name" {
			idxCountry = i
			break
		}
	}
	if idxCountry == -1 {
		fmt.Println("country_name column not found in csv")
		return
	}

	counter := make(map[string]int)
	rowNum := 1
	for {
		row, err := r.Read()
		rowNum++
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error read row %d: %v\n", rowNum, err)
			continue
		}
		if len(row) <= idxCountry {
			continue
		}
		country := row[idxCountry]
		counter[country]++
	}

	var results []CountryCityCount
	for country, count := range counter {
		results = append(results, CountryCityCount{
			CountryName: country,
			CityCount:   count,
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].CityCount > results[j].CityCount
	})

	saveToJSON(filepath.Join("soal_b", "output.json"), results)
    saveToCSV(filepath.Join("soal_b", "output.csv"), results)

	fmt.Println("output.csv and output.json have been created successfully.")
}

func saveToJSON(filename string, data []CountryCityCount) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed write json:", err)
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func saveToCSV(filename string, data []CountryCityCount) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed write csv:", err)
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"country_name", "city_count"})

	for _, d := range data {
		w.Write([]string{d.CountryName, strconv.Itoa(d.CityCount)})
	}
}