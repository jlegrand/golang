package basics

import (
	"encoding/json"
	"fmt"
	"os"
	"encoding/xml"
)

type ColorGroup struct {
	ID int				`json:"uid"`
	Name string			`json:"nom,omitempty"`
	Colors []string
	comment string  // non exporté donc non marshallé
}

func Json() {

	group := ColorGroup{ ID: 1, Name:"Reds", Colors:[]string{"Crimson", "Red", "Ruby"}}

	// Marshall JSON

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()

	// Unmarshall JSON

	var colorGroup ColorGroup
	err = json.Unmarshal(b, &colorGroup)
	if err == nil {
		fmt.Printf("%+v\n", colorGroup)
	} else {
		fmt.Println(err)
	}

	// Marshall XML

	b, err = xml.Marshal(colorGroup)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()

	// Unmarshall XML

	err = xml.Unmarshal(b, &colorGroup)
	if err == nil {
		fmt.Printf("%+v\n", colorGroup)
	} else {
		fmt.Println(err)
	}

}
