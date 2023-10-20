package main

import (
	"apiGateway/cmd/app/handler/user"
	"apiGateway/deco/handler"
	"apiGateway/jwt"
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
	checkFunc := getTokenCheckFunc(p)

	uh2 := user.NewHandler2()
	wraped := handler.NewDecoHandler(uh2, checkFunc)
	r.PathPrefix("/users2").Handler(wraped)

	uh := user.NewHandler()
	handler.NewDecoHandler(uh, checkFunc)
	r.PathPrefix("/users").Handler(uh)
	return r
}

func getTokenCheckFunc(p jwt.Provider) func(http.ResponseWriter, *http.Request, http.Handler) {
	am := handler.NewAuthMiddleware(p)
	return am.CheckToken
}
