package cafe

import (
	"apiGateway/internal/controller/cafe"
	"apiGateway/internal/controller/cafe/req"
	"apiGateway/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	c cafe.Controller
}

func NewHandler(c cafe.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes", h.create).Methods(http.MethodPost)
	r.HandleFunc("/cafes", h.getList).Methods(http.MethodGet)
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

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	reqPage := page.GetPageReqByRequest(r)
	cafeList, count, err := h.c.GetList(r.Context(), reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageDto := page.GetPagination(cafeList, reqPage, count)

	data, err := json.Marshal(pageDto)
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
