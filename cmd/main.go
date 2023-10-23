package main

import (
	"apiGateway/cmd/app/handler/todo"
	"apiGateway/cmd/app/handler/user"
	"apiGateway/internal/cli"
	todo2 "apiGateway/internal/controller/todo"
	user2 "apiGateway/internal/controller/user"
	handler3 "apiGateway/internal/deco/handler"
	"apiGateway/internal/jwt"
	todo3 "apiGateway/internal/service/todo"
	user3 "apiGateway/internal/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := newHandler()
	http.ListenAndServe(":8080", r)
}

func newHandler() http.Handler {
	r := mux.NewRouter()
	// jwt 관련
	p := jwt.NewProvider("this_is_my_jwt_token_secret_key")

	// token 검사 func
	checkFunc := getTokenCheckFunc(p)

	// user 관련
	uh := getUserHandler(p)
	r.PathPrefix("/users").Handler(uh)

	// todos 관련
	th := getTodoHandler()
	wrappedTodoHandler := handler3.NewDecoHandler(th, checkFunc)
	r.PathPrefix("/todos").Handler(wrappedTodoHandler)

	return r
}

func getTokenCheckFunc(p jwt.Provider) func(http.ResponseWriter, *http.Request, http.Handler) {
	am := handler3.NewAuthMiddleware(p)
	return am.CheckToken
}

func getUserHandler(p jwt.Provider) http.Handler {
	var loginUrl = "http://localhost:8081/users/login"
	var userCreateUrl = "http://localhost:8081/users"

	return user.NewHandler(user2.NewController(user3.NewService(p, cli.NewUserRequester(loginUrl, userCreateUrl))))
}

func getTodoHandler() http.Handler {
	var todoUrl = "http://localhost:8082/todos"
	return todo.NewHandler(todo2.NewController(todo3.NewService(cli.NewTodoRequester(todoUrl))))
}
