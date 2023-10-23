package todo

import (
	"apiGateway/internal/controller/todo"
	"apiGateway/internal/controller/todo/req"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	c todo.Controller
}

func NewHandler(c todo.Controller) http.Handler {
	m := mux.NewRouter()

	h := Handler{c: c}
	m.HandleFunc("todos", h.getTodoList).Methods(http.MethodGet)
	m.HandleFunc("todos/{id:[0-9]+}", h.getTodoDetail).Methods(http.MethodGet)
	m.HandleFunc("todos", h.createTodo).Methods(http.MethodPost)
	m.HandleFunc("todos", h.updateTodo).Methods(http.MethodPut)
	m.HandleFunc("todos/{id:[0-9]+}", h.deleteTodo).Methods(http.MethodDelete)
	return m
}

func (h Handler) getTodoList(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) getTodoDetail(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) createTodo(w http.ResponseWriter, r *http.Request) {
	ct := req.CreateTodoDto{}
	err := json.NewDecoder(r.Body).Decode(&ct)
	if err != nil {
		log.Println("createTodo decode fail", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.c.CreateTodo(r.Context(), ct)
	if err != nil {
		// todo 에러에맞게 핸들링하기
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}

func (h Handler) updateTodo(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {

}
