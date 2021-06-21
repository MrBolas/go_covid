package controllers

import (
	"encoding/json"
	"fmt"
	"go_covid/models"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURLv2 string = "https://corona.lmao.ninja/v2/countries/"
const baseURLv3 string = "https://disease.sh/v3/covid-19/countries/"
const baseHistorucURLv3 string = "https://disease.sh/v3/covid-19/historical/"
const vacineUrl string = "https://disease.sh/v3/covid-19/vaccine/coverage/countries/" //Portugal?lastdays=3&fullData=true"
const vacineUrlParams string = "?lastdays=2&fullData=true"

func GetHistoricalCountryData(country string) models.CountryHistory {
	// Get country data
	url := baseHistorucURLv3 + country + "?lastdays=3"
	fmt.Println(url)
	resp, err := http.Get(url)
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
	usedCountry := models.CountryHistory{}

	// Unmarshal Covid Data
	jsonErr := json.Unmarshal(body, &usedCountry)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(usedCountry)

	return usedCountry
}

func GetVacineData(country string) {
	resp, _ := http.Get(vacineUrl + country + vacineUrlParams)

	// Verify Body not empty
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Reads body
	_, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
}

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

func validateCountry(country string) (string, error) {
	return country, nil
}
