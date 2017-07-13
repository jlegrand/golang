package basics

import (
	"net/http"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	h := new(Handler)
	return h
}

func (h *Handler) ServeHTTP( w http.ResponseWriter, r *http.Request) {
	//var buffer bytes.Buffer

	w.Write([]byte(r.URL.String()[1:]))



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

