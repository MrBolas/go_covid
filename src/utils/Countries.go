package utils

import (
	s "strings"

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
			if s.EqualFold(country.Alpha2, input) {
				return country, nil
			}
			if s.EqualFold(country.Alpha3, input) {
				return country, nil
			}
			for _, tld := range country.TLDs {
				if s.EqualFold(tld, searchTLD) {
					return country, nil
				}
			}
		}
	}
	return result, err
}
