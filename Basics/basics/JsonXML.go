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
	Comment string		`xml:",comment"`
}

func Json() {

	group := ColorGroup{ ID: 1, Name:"Reds", Colors:[]string{"Crimson", "Red", "Ruby"}, Comment:"commentaire"}

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

type Person struct {
	Name string					`xml:"FullName"`
	Company string
	Email []Email				`xml:""`
	Group []string				`xml:"Group>Value"`
	City string
	State string
}

type Email struct {
	Where string				`xml:"where,attr"`
	Addr string
}

func UnmarshallData() {

	data := []byte(`
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
			`)

	var personne Person
	err := xml.Unmarshal(data, &personne)
	if err == nil {
		fmt.Printf("%+v\n", personne)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

}
