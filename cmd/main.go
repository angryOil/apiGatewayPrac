package main

import (
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

var loginUrl = "http://localhost:8081/users/login"

func newHandler() http.Handler {
	r := mux.NewRouter()
	// jwt 관련
	p := jwt.NewProvider("this_is_my_jwt_token_secret_key")

	// token 검사 func
	checkFunc := getTokenCheckFunc(p)

	// user 관련
	uh := getUserHandler(p)
	handler.NewDecoHandler(uh, checkFunc)
	r.PathPrefix("/users").Handler(uh)
	return r
}

func getTokenCheckFunc(p jwt.Provider) func(http.ResponseWriter, *http.Request, http.Handler) {
	am := handler.NewAuthMiddleware(p)
	return am.CheckToken
}

func getUserHandler(p jwt.Provider) http.Handler {
	return user.NewHandler(user2.NewController(user3.NewService(p, req.NewRequester(loginUrl))))
}