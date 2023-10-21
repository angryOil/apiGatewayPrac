package req

import (
	"apiGateway/domain"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type UserRequester struct {
	loginUrl string
}

func NewRequester(loginUrl string) UserRequester {
	return UserRequester{loginUrl: loginUrl}
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
