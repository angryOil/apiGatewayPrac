package main

import (
	handler2 "apiGateway/cmd/app/handler"
	"apiGateway/cmd/app/handler/user"
	user2 "apiGateway/controller/user"
	"apiGateway/deco/handler"
	"apiGateway/jwt"
	user3 "apiGateway/service/user"
	"apiGateway/service/user/req"
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

	// test 입니다
	t := handler2.NewTestHandler()

	wrappedTest := handler.NewDecoHandler(t, checkFunc)
	r.PathPrefix("/test").Handler(wrappedTest)
	return r
}

func getTokenCheckFunc(p jwt.Provider) func(http.ResponseWriter, *http.Request, http.Handler) {
	am := handler.NewAuthMiddleware(p)
	return am.CheckToken
}

func getUserHandler(p jwt.Provider) http.Handler {
	var loginUrl = "http://localhost:8081/users/login"
	var userCreateUrl = "http://localhost:8081/users"

	return user.NewHandler(user2.NewController(user3.NewService(p, req.NewRequester(loginUrl, userCreateUrl))))
}
