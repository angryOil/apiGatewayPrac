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
	//h := Handler{c: c}
	return r
}
