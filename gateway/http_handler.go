package main

import (
	"log"
	"net/http"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/users/{userId}/cards", h.HandleCreateCard)
}

func (h *handler) HandleCreateCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello")
}