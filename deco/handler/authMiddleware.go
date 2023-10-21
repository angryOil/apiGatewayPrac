package handler

import (
	"apiGateway/jwt"
	"context"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	p jwt.Provider
}

func NewAuthMiddleware(p jwt.Provider) AuthMiddleware {
	return AuthMiddleware{p: p}
}

func (a AuthMiddleware) CheckToken(w http.ResponseWriter, r *http.Request, h http.Handler) {
	token := r.Header.Get("token")
	if !tokenCheck(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token"))
		return
	}

	result, err := a.p.ValidToken(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server err"))
		return
	}
	if !result {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	// ctx 에 token 추가
	ctx := ContextWithToken(r.Context(), token)
	h.ServeHTTP(w, r.WithContext(ctx))
}

func tokenCheck(token string) bool {
	if token == "" {
		return false
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}
	return true
}

func ContextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "token", token)
}
