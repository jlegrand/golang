package server

import (
	"net/http"
	"time"
	"encoding/json"
	"strconv"
	"fmt"
)

type Handler struct {
	provider *provider
}

func NewHandler() *Handler {
	h := new(Handler)
	h.provider = NewProvider("C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Mail/server/_store")
	return h
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

func (h *Handler) ServeHTTP( w http.ResponseWriter, r *http.Request) {

	go func() {
		if err := recover(); err != nil {
			errorHandler(w, r, http.StatusNotFound)
		}
	}()

	id := r.URL.String()[1:]
	id_int, _ := strconv.Atoi(id)
	msg := h.provider.Get(uint64(id_int))

	x, err := json.Marshal(msg)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
	}

	w.Write(x)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

