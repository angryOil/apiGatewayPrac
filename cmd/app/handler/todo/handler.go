package todo

import (
	"apiGateway/internal/controller/todo"
	"apiGateway/internal/controller/todo/req"
	page2 "apiGateway/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c todo.Controller
}

func NewHandler(c todo.Controller) http.Handler {
	m := mux.NewRouter()

	h := Handler{c: c}
	m.HandleFunc("/todos", h.getTodoList).Methods(http.MethodGet)
	m.HandleFunc("/todos/{id:[0-9]+}", h.getTodoDetail).Methods(http.MethodGet)
	m.HandleFunc("/todos", h.createTodo).Methods(http.MethodPost)
	m.HandleFunc("/todos/{id:[0-9]+}", h.updateTodo).Methods(http.MethodPut)
	m.HandleFunc("/todos/{id:[0-9]+}", h.deleteTodo).Methods(http.MethodDelete)
	return m
}

func (h Handler) getTodoList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	pReq := page2.NewReqPage(page, size)
	result, err := h.c.GetTodoList(r.Context(), pReq)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("json marshal err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) getTodoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if todoId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("todoId 가 없습니다."))
		return
	}
	resDto, err := h.c.GetTodoDetail(r.Context(), todoId)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(resDto)
	if err != nil {
		log.Println("marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func (h Handler) createTodo(w http.ResponseWriter, r *http.Request) {
	ct := req.CreateTodoDto{}
	err := json.NewDecoder(r.Body).Decode(&ct)
	if err != nil {
		log.Println("createTodo decode fail: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("잘못된값으로 요청했습니다."))
		return
	}

	err = h.c.CreateTodo(r.Context(), ct)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}

func (h Handler) updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ut := req.UpdateTodoDto{}
	err = json.NewDecoder(r.Body).Decode(&ut)
	if err != nil {
		log.Println("updateTodo decode fail: ", err)
		http.Error(w, "잘못된 값으로 요청했습니다.", http.StatusBadRequest)
		return
	}

	err = h.c.UpdateTodo(r.Context(), ut, todoId)
	if err != nil {
		if !strings.Contains(err.Error(), "internal") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (h Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.DeleteTodo(r.Context(), todoId)

	if err != nil {
		if !strings.Contains(err.Error(), "internal") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "server internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
