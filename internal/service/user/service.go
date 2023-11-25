package user

import (
	req2 "apiGateway/internal/cli/user/req"
	"apiGateway/internal/domain/user"
	"apiGateway/internal/jwt"
	"apiGateway/internal/service/user/req"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type Service struct {
	p  jwt.Provider
	ur req2.UserRequester
}

func NewService(p jwt.Provider, ur req2.UserRequester) Service {
	return Service{p: p, ur: ur}
}

const (
	InternalServerError = "internal server error"
)

func (s Service) Login(ctx context.Context, l req.Login) (string, error) {
	resToken, err := s.ur.Login(ctx, req2.Login{
		Email:    l.Email,
		Password: l.Password,
	})
	if err != nil {
		if strings.Contains(err.Error(), "login") {
			return "", err
		}
		log.Println(err)
		return "", err
	}

	resUserDomain, err := tokenToDomain(resToken)
	if err != nil {
		return "", err
	}

	token, err := s.p.CreateToken(resUserDomain)
	return token, err
}

func tokenToDomain(token string) (user.User, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		//
		log.Printf("로그인 결과가 토큰이아닙니다. token:%s ", token)
		return user.NewBuilder().Build(), errors.New(InternalServerError)
	}
	u, err := payloadToDomain(parts[1])
	return u, err
}

func payloadToDomain(payload string) (user.User, error) {
	data, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return user.NewBuilder().Build(), err
	}
	var u user.User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return user.NewBuilder().Build(), err
	}
	return u, nil
}

func (s Service) CreateUser(ctx context.Context, c req.CreateUser) error {
	email, password := c.Email, c.Password
	err := user.NewBuilder().
		Email(email).
		Password(password).
		Build().ValidCreateUser()
	if err != nil {
		return err
	}

	err = s.ur.CreateUser(ctx, req2.CreateUser{
		Email:    email,
		Password: password,
	})
	return err
}
