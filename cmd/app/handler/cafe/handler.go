package cafe

import (
	"apiGateway/internal/controller/cafe"
	"apiGateway/internal/controller/cafe/req"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	c cafe.Controller
}

func NewHandler(c cafe.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes", h.create).Methods(http.MethodPost)
	return r
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	var c req.CreateCafeDto
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.c.Create(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
