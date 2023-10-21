package user

import (
	"apiGateway/domain"
	"apiGateway/jwt"
	"apiGateway/service/user/req"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type Service struct {
	p  jwt.Provider
	ur req.UserRequester
}

func NewService(p jwt.Provider, ur req.UserRequester) Service {
	return Service{p: p, ur: ur}
}

func (s Service) Login(ctx context.Context, u domain.User) (string, error) {
	resToken, err := s.ur.Login(ctx, u)
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

func tokenToDomain(token string) (domain.User, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		//
		log.Printf("로그인 결과가 토큰이아닙니다. token:%s ", token)
		return domain.User{}, errors.New("internal server error")
	}
	u, err := payloadToDomain(parts[1])
	return u, err
}

func payloadToDomain(payload string) (domain.User, error) {
	data, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}
