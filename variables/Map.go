package variables

import "fmt"

func Basics() {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, " is ", countryCapitalMap[country])
	}
}
