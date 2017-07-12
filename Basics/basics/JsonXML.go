package basics

import (
	"encoding/json"
	"fmt"
	"os"
	"encoding/xml"
)

type ColorGroup struct {
	XMLName xml.Name	`xml:"GroupeDeCouleur"`
	ID int				`json:"uid" xml:"id,attr"`
	Name string			`json:"nom,omitempty"`
	Colors []string		`xml:"colors>color"`
	comment string  // non exporté donc non marshallé
}

func Json() {

	group := ColorGroup{ ID: 1, Name:"Reds", Colors:[]string{"Crimson", "Red", "Ruby"}}

	// Marshall JSON

	b, err := json.MarshalIndent(group, "", "\t")
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	fmt.Println()

	// Unmarshall JSON

	var colorGroup ColorGroup
	err = json.Unmarshal(b, &colorGroup)
	if err == nil {
		fmt.Printf("%+v\n", colorGroup)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

	// Marshall XML

	b, err = xml.MarshalIndent(colorGroup, "", "\t")
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	fmt.Println()

	// Unmarshall XML

	err = xml.Unmarshal(b, &colorGroup)
	if err == nil {
		fmt.Printf("%+v\n", colorGroup)
	} else {
		fmt.Println(err)
	}

}
