package basics

import "fmt"

func MapBasics() {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of ", country, " is ", capital)
	}

	capital, exists := countryCapitalMap["USA"]
	if exists {
		fmt.Println("Capital of USA is ", capital)
	} else {
		fmt.Println("Capital of USA is not present")
	}
}
