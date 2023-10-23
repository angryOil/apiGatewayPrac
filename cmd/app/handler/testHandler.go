package handler

import (
	"apiGateway/internal/service/common"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() http.Handler {
	h := TestHandler{}
	m := mux.NewRouter()
	m.HandleFunc("/test", h.THandling).Methods(http.MethodGet)

	return m
}

func (h TestHandler) THandling(w http.ResponseWriter, r *http.Request) {
	token, ok := common.TokenFromContext(r.Context())
	fmt.Println("get token", token, ok)
}
