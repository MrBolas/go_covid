package apicontrollers

import (
	"github.com/pariz/gountries"
)

func SearchCountry(input string) (gountries.Country, error) {
	query := gountries.New()
	result, err := query.FindCountryByName(input)
	if err != nil {
		result, err = query.FindCountryByAlpha(input)
	}
	if err != nil {
		searchTLD := "." + input
		countries := query.FindAllCountries()
		for _, country := range countries {
			for _, tld := range country.TLDs {
				if tld == searchTLD {
					return country, nil
				}
			}
		}
	}
	return result, err
}
