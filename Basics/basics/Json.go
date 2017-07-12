package basics

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID int				`json:"uid"`
	Name string			`json:"nom,omitempty"`
	Colors []string
}

func Json() {

	group := ColorGroup{ ID: 1, Name:"Reds", Colors:[]string{"Crimson", "Red", "Ruby"}}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
}
