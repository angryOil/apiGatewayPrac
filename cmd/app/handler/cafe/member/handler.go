package member

import (
	"apiGateway/internal/controller/cafe/member"
	"apiGateway/internal/controller/cafe/member/res"
	"apiGateway/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	c member.Controller
}

func NewHandler(c member.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes/members/my", h.getMyCafeList).Methods(http.MethodGet)
	return r
}

const (
	InvalidCafeId       = "invalid cafe id"
	InternalServerError = "internal server error"
	ApplicationJson     = "application/json"
)

func (h Handler) getMyCafeList(w http.ResponseWriter, r *http.Request) {
	reqPage := page.GetPageReqByRequest(r)

	cafes, total, err := h.c.GetMyCafeList(r.Context(), reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	listTotalDto := res.NewCafeListTotalDto(cafes, total)
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		log.Println("getMyCafeList json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", ApplicationJson)
	w.Write(data)
}
