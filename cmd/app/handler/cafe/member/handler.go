package member

import (
	"apiGateway/internal/controller/cafe/member"
	"apiGateway/internal/controller/cafe/member/req"
	"apiGateway/internal/controller/cafe/member/res"
	"apiGateway/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c member.Controller
}

func NewHandler(c member.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes/members/my", h.getMyCafeList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/members/info", h.getMemberInfo).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/members/join", h.joinCafe).Methods(http.MethodPost)
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

func (h Handler) getMemberInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	dto, err := h.c.GetMemberInfo(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(dto)
	if err != nil {
		log.Println("getMemberInfo json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", ApplicationJson)
	w.Write(data)
}

func (h Handler) joinCafe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}

	var joinDto req.JoinCafe
	err = json.NewDecoder(r.Body).Decode(&joinDto)
	if err != nil {
		log.Println("joinCafe json.NewDecoder err: ", err)
		http.Error(w, InternalServerError, http.StatusBadRequest)
		return
	}
	err = h.c.JoinCafe(r.Context(), cafeId, joinDto)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
