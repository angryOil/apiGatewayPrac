package cafe

import (
	req2 "apiGateway/internal/cli/cafe/req"
	"apiGateway/internal/service/common"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

var baseUrl = "http://localhost:8083/cafes"

const (
	InvalidToken        = "invalid token"
	InternalServerError = "internal server error"
)

func (r Requester) Create(ctx context.Context, c req2.Create) error {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New(InvalidToken)
	}
	var buf bytes.Buffer
	dto := c.ToCreateDto()
	err := json.NewEncoder(&buf).Encode(dto)
	if err != nil {
		log.Println("Create json.NewEncoder err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("POST", baseUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err: ", err)
		return errors.New(InternalServerError)
	}
	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Create DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Create readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}
