package req

import (
	"apiGateway/domain"
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

func NewRequester(loginUrl string, createUrl string) UserRequester {
	return UserRequester{loginUrl: loginUrl, userCreateUrl: createUrl}
}

func (ur UserRequester) Login(cxt context.Context, u domain.User) (string, error) {
	data, err := json.Marshal(u)
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
		return "", errors.New("login fail")
	}
	return string(readBody), nil
}

func (ur UserRequester) CreateUser(ctx context.Context, u domain.User) error {
	data, err := json.Marshal(u)
	if err != nil {

		return err
	}

	re, err := http.NewRequest("POST", ur.userCreateUrl, bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		return errors.New("internal server error")
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusBadRequest {
			return errors.New("bad request: " + string(resBody))
		}
		if resp.StatusCode == http.StatusConflict {
			return errors.New("conflict:" + string(resBody))
		}
		return errors.New(string(resBody))
	}
	return nil
}
