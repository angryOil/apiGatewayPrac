package main

import (
	"apiGateway/cmd/app/handler/cafe/cafe"
	"apiGateway/cmd/app/handler/cafe/member"
	"apiGateway/cmd/app/handler/todo"
	"apiGateway/cmd/app/handler/user"
	cafe4 "apiGateway/internal/cli/cafe/cafe"
	member4 "apiGateway/internal/cli/cafe/member"
	todo4 "apiGateway/internal/cli/todo"
	"apiGateway/internal/cli/user/req"
	cafe2 "apiGateway/internal/controller/cafe/cafe"
	member2 "apiGateway/internal/controller/cafe/member"
	todo2 "apiGateway/internal/controller/todo"
	user2 "apiGateway/internal/controller/user"
	handler3 "apiGateway/internal/deco/handler"
	"apiGateway/internal/jwt"
	cafe3 "apiGateway/internal/service/cafe/cafe"
	member3 "apiGateway/internal/service/cafe/member"
	todo3 "apiGateway/internal/service/todo"
	user3 "apiGateway/internal/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := getHandler()
	http.ListenAndServe(":8080", r)
}

func getHandler() http.Handler {
	r := mux.NewRouter()
	// jwt 관련
	p := jwt.NewProvider("this_is_my_jwt_token_secret_key")

	// token 검사 func
	//checkFunc := getTokenCheckFunc(p)
	// user 관련
	uh := getUserHandler(p)
	r.PathPrefix("/users").Handler(uh)

	s := r.PathPrefix("/").Subrouter()
	s.Use(tokenMiddleWare)
	// todos 관련
	th := getTodoHandler()
	//wrappedTodoHandler := handler3.NewDecoHandler(th, checkFunc)
	s.PathPrefix("/todos").Handler(th)
	registrationCafeHandler(s)
	return r
}

func tokenMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authMiddle := handler3.NewAuthMiddleware(jwt.NewProvider("this_is_my_jwt_token_secret_key"))
		authMiddle.CheckToken(w, r, next)
	})
}

func getMemberHandler() http.Handler {
	return member.NewHandler(member2.NewController(member3.NewService(member4.NewRequester())))
}

func registrationCafeHandler(r *mux.Router) {
	memberH := getMemberHandler()
	r.PathPrefix("/cafes/members").Handler(memberH)
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(memberH)

	r.PathPrefix("/cafes").Handler(cafe.NewHandler(cafe2.NewController(cafe3.NewService(cafe4.NewRequester()))))
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
