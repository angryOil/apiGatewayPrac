package user

import (
	"apiGateway/controller/user"
	"apiGateway/controller/user/req"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	c user.Controller
}

func NewHandler(c user.Controller) http.Handler {
	uh := Handler{c: c}
	m := mux.NewRouter()
	m.HandleFunc("/users/login", uh.login).Methods(http.MethodPost)
	return m
}

func (h Handler) login(w http.ResponseWriter, r *http.Request) {
	var ld = req.LoginDto{}
	err := json.NewDecoder(r.Body).Decode(&ld)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	token, err := h.c.Login(r.Context(), ld)
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("login fail"))
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
