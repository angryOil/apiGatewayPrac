package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
)

type Handler2 struct {
}

func NewHandler2() http.Handler {
	uh := Handler2{}
	m := mux.NewRouter()
	m.HandleFunc("/users2/login", uh.login2).Methods(http.MethodPost)
	return m
}

type tempLoginUser2 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler2) login2(w http.ResponseWriter, r *http.Request) {
	var lu = tempLoginUser{}
	err := json.NewDecoder(r.Body).Decode(&lu)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token"))
		return
	}

	data, err := json.Marshal(&lu)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token"))
		return
	}

	re, err := http.NewRequest("POST", "http://localhost:8081/users/login", strings.NewReader(string(data)))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	defer resp.Body.Close()
	readBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(readBody))
	w.WriteHeader(resp.StatusCode)
	w.Write(readBody)
}
