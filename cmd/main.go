package main

import (
	"apiGateway/cmd/app/handler/cafe"
	"apiGateway/cmd/app/handler/todo"
	"apiGateway/cmd/app/handler/user"
	cafe4 "apiGateway/internal/cli/cafe"
	todo4 "apiGateway/internal/cli/todo"
	"apiGateway/internal/cli/user/req"
	cafe2 "apiGateway/internal/controller/cafe"
	todo2 "apiGateway/internal/controller/todo"
	user2 "apiGateway/internal/controller/user"
	handler3 "apiGateway/internal/deco/handler"
	"apiGateway/internal/jwt"
	cafe3 "apiGateway/internal/service/cafe"
	todo3 "apiGateway/internal/service/todo"
	user3 "apiGateway/internal/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	cafeH := getCafeHandler()
	r.PathPrefix("/cafes").Handler(cafeH)
	http.ListenAndServe(":8080", r)
}

func getCafeHandler() http.Handler {
	return cafe.NewHandler(cafe2.NewController(cafe3.NewService(cafe4.NewRequester())))
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

	return user.NewHandler(user2.NewController(user3.NewService(p, req.NewUserRequester(loginUrl, userCreateUrl))))
}

func getTodoHandler() http.Handler {
	var todoUrl = "http://localhost:8082/todos"
	return todo.NewHandler(todo2.NewController(todo3.NewService(todo4.NewTodoRequester(todoUrl))))
}
