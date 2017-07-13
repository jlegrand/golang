package basics

import (
	"net/http"
	"time"
	"bytes"
)

type Handler struct {

}

func NewHandler() *Handler {
	return new(Handler)
}

func (h *Handler) ServeHTTP( w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("hello World !")
	w.Write(buffer.Bytes())
}

func Server() {

	handler := NewHandler()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(server.ListenAndServe())
	
}
func handleConnection() {

}
