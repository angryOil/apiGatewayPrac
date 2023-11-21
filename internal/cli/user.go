package cli

import (
	"apiGateway/internal/cli/user/req"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type UserRequester struct {
	loginUrl      string
	userCreateUrl string
}

func NewUserRequester(loginUrl string, createUrl string) UserRequester {
	return UserRequester{loginUrl: loginUrl, userCreateUrl: createUrl}
}

const (
	InternalServerError = "internal server error"
	LoginFail           = "login fail"
)

func (ur UserRequester) Login(cxt context.Context, l req.Login) (string, error) {
	dto := l.ToLoginDto()
	data, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	re, err := http.NewRequest("POST", ur.loginUrl, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	readBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(LoginFail)
	}
	return string(readBody), nil
}

func (ur UserRequester) CreateUser(ctx context.Context, c req.CreateUser) error {
	dto := c.ToCreateUserDto()
	data, err := json.Marshal(dto)
	if err != nil {

		return err
	}

	re, err := http.NewRequest("POST", ur.userCreateUrl, bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(resBody))
	}
	return nil
}
