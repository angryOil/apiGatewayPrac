package cafe

import (
	"apiGateway/internal/controller/cafe/cafe"
	req2 "apiGateway/internal/controller/cafe/cafe/req"
	"apiGateway/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	c cafe.Controller
}

func NewHandler(c cafe.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes", h.create).Methods(http.MethodPost)
	r.HandleFunc("/cafes", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{id:[0-9]+}", h.getDetail).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	return r
}

const (
	InvalidId           = "invalid cafe id"
	InternalServerError = "internal server error"
)

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	var c req2.CreateCafeDto
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

func (h Handler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	detail, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(detail)
	if err != nil {
		log.Println("getDetail json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	var d req2.PatchDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("patch json.NewDecoder err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = h.c.Patch(r.Context(), id, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
