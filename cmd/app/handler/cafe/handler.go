package cafe

import (
	"apiGateway/internal/controller/cafe"
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

}
