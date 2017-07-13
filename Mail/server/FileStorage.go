package server

import (
	"github.com/jlegrand/golang/Mail/mail"
	"encoding/json"
	"os"
	"strconv"
	"bytes"
)

type provider struct {
	mail_directory_path string
}

func NewProvider(mail_directory_path string) *provider {

	p := new(provider)
	p.mail_directory_path = mail_directory_path
	return p
}

func (p *provider) Get(id uint64) *mail.Message {

	h, err := os.Open(p.mail_directory_path + string(os.PathSeparator) + strconv.Itoa(int(id)) + ".json")
	if err != nil {
		panic(err)
	}
	defer h.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(h)

	// Unmarshalling
	var msg *mail.Message
	err = json.Unmarshal(b.Bytes(), &msg)
	if err != nil {
		panic(err)
	}

	return msg
}

func (p *provider) Set(id uint64, msg *mail.Message) {

	// Marshalling
	x, err := json.Marshal(msg)
	if err != nil {
		panic("Cannot marshall message")
	}

	// Write XML file
	fileWrite, err := os.Create(p.mail_directory_path + string(os.PathSeparator) + strconv.Itoa(int(id)) + ".json")
	if err != nil {
		panic(err)
	}
	defer fileWrite.Close()

	_, err = fileWrite.Write(x)
	if err != nil {
		panic(err)
	}
}
