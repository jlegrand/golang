package basics

import (
	"os"
	"bytes"
	"encoding/json"
	"fmt"
	"encoding/xml"
	"bufio"
)

const FICHIER_TO_READ string = "C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Basics/basics/file/colors.json"
const FICHIER_TO_WRITE string = "C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Basics/basics/file/colors_generated.json"

func FileRead() {

	h, err := os.Open(FICHIER_TO_READ)
	if err != nil {
		panic(err)
	}

	defer h.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(h)

	var colorGroup ColorGroup
	err = json.Unmarshal(b.Bytes(), &colorGroup)
	if err == nil {
		fmt.Printf("%+v\n", colorGroup)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

}

func FileReadJsonWriteXML() {

	// Read Json file
	fileRead, err := os.Open(FICHIER_TO_READ)
	if os.IsNotExist(err) {
		panic("File " + FICHIER_TO_READ  + " doesn't exist")
	}
	defer fileRead.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(fileRead)

	// Unmarshall json
	var colorGroup ColorGroup
	err = json.Unmarshal(b.Bytes(), &colorGroup)
	if err != nil {
		panic("Cannot unmarshall Json file " + FICHIER_TO_READ)
	}

	// Marshall XML
	x, err := xml.MarshalIndent(colorGroup, "", "\t")
	if err != nil {
		panic("Cannot marshall XML file " + FICHIER_TO_READ)
	}

	// Write XML file
	fileWrite, err := os.Create(FICHIER_TO_WRITE)
	if err != nil {
		panic(err)
	}
	defer fileWrite.Close()

	_, err = fileWrite.Write(x)
	if err != nil {
		panic(err)
	}
}

func FileOtherWriteMethod() {

	// Read Json file
	fileWrite, err := os.Create(FICHIER_TO_WRITE)
	if os.IsNotExist(err) {
		panic("File " + FICHIER_TO_WRITE  + " doesn't exist")
	}
	defer fileWrite.Close()

	// method 1
	_, err = fileWrite.WriteString("Hello World ! ")
	fileWrite.Sync()

	// method 2
	w := bufio.NewWriter(fileWrite)
	_, err = w.WriteString("buffered ! ")
	w.Flush()

	//method 3
	buf := new(bytes.Buffer)
	buf.WriteString("Hello")
	buf.WriteTo(fileWrite)


}


