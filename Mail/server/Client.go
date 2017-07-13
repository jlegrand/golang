package server

import (
	"net/http"
	"bytes"
	"fmt"
)

func Client() {

	var c http.Client
	r, err := c.Get("http://localhost:8080/1")
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(r.Body)
	fmt.Println(buffer.String())
}
