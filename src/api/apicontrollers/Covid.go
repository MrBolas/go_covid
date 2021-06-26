package apicontrollers

import (
	"encoding/json"
	"fmt"
	models "go_covid/src/api/apimodels"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURLv2 string = "https://corona.lmao.ninja/v2/countries/"
const baseURLv3 string = "https://disease.sh/v3/covid-19/countries/"
const baseHistorucURLv3 string = "https://disease.sh/v3/covid-19/historical/"
const vacineUrl string = "https://disease.sh/v3/covid-19/vaccine/coverage/countries/" //Portugal?lastdays=3&fullData=true"
const vacineUrlParams string = "?lastdays=2&fullData=true"

// GetHistoricalCountryData gets country covid history for the past x days from API.
// Returns a CountryHistory Struct
func GetHistoricalCountryData(country string, days int) models.CountryHistory {
	// Get country data
	url := fmt.Sprintf("%s%s?lastdays=%d", baseHistorucURLv3, country, days)
	resp, err := http.Get(url)

	// handle error
	if err != nil {
		log.Fatal(err)
	}

	// Verify Body not empty
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Reads body
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Creates model variable
	usedCountry := models.CountryHistory{}

	// Unmarshal Covid Data
	jsonErr := json.Unmarshal(body, &usedCountry)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return usedCountry
}

// GetVaccineData gets country covid vaccine data from API.
// Returns a VaccineCountryData Struct
func GetVaccineData(country string) models.VaccineCountryData {
	resp, _ := http.Get(vacineUrl + country + vacineUrlParams)

	// Verify Body not empty
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Reads body
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Creates model variable
	vaccineData := models.VaccineCountryData{}

	// Unmarshal Covid Data
	jsonErr := json.Unmarshal(body, &vaccineData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vaccineData
}

// GetCountryData gets generic country covid data from API.
// Returns a Country Struct
func GetCountryData(country string) models.Country {
	// Get country data
	resp, err := http.Get(baseURLv3 + country)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// Verify Body not empty
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Reads body
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Creates model variable
	usedCountry := models.Country{}

	// Unmarshal Covid Data
	jsonErr := json.Unmarshal(body, &usedCountry)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return usedCountry
}

// validateCountry validates covid data from API.
// Returns a String and error
func validateCountry(country string) (string, error) {
	return country, nil
}
