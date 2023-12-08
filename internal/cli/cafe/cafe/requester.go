package cafe

import (
	"apiGateway/internal/cli/cafe/cafe/model"
	"apiGateway/internal/cli/cafe/cafe/req"
	cafe2 "apiGateway/internal/domain/cafe/cafe"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/common"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	Token               = "token"
	InvalidToken        = "invalid token"
	InternalServerError = "internal server error"
)

func (r Requester) Create(ctx context.Context, c req.Create) error {
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
	re.Header.Add(Token, token)

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

func (r Requester) GetList(ctx context.Context, reqPage page2.ReqPage) ([]cafe2.Cafe, int, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return []cafe2.Cafe{}, 0, errors.New(InvalidToken)
	}
	reqUrl := fmt.Sprintf("%s?page=%d&size=%d", baseUrl, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []cafe2.Cafe{}, 0, errors.New(InternalServerError)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient err: ", err)
		return []cafe2.Cafe{}, 0, errors.New(InternalServerError)
	}
	var cafePage model.CafePage
	err = json.NewDecoder(resp.Body).Decode(&cafePage)
	if err != nil {
		log.Println("GetList NewDecoder err: ", err)
		return []cafe2.Cafe{}, 0, errors.New(InternalServerError)
	}

	return model.ToDomainList(cafePage.Contents), cafePage.Total, nil
}

func (r Requester) GetDetail(ctx context.Context, id int) (cafe2.Cafe, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return cafe2.NewCafeBuilder().Build(), errors.New(InvalidToken)
	}
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, id)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetDetail NewRequest err: ", err)
		return cafe2.NewCafeBuilder().Build(), errors.New(InternalServerError)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetDetail DefaultClient.Do err: ", err)
		return cafe2.NewCafeBuilder().Build(), errors.New(InternalServerError)
	}

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetDetail readBody err: ", err)
			return cafe2.NewCafeBuilder().Build(), errors.New(InternalServerError)
		}
		return cafe2.NewCafeBuilder().Build(), errors.New(string(readBody))
	}
	var m model.Cafe
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("GetDetail NewDecoder err: ", err)
		return cafe2.NewCafeBuilder().Build(), errors.New(InternalServerError)
	}
	return m.ToDomain(), nil
}

func (r Requester) Patch(ctx context.Context, p req.Patch) error {
	id := p.Id
	patchDto := p.ToPatchDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(patchDto)
	if err != nil {
		log.Println("Patch json.NewEncoder err: ", err)
		return errors.New(InternalServerError)
	}
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, id)
	re, err := http.NewRequest("PUT", reqUrl, &buf)
	if err != nil {
		log.Println("Patch NewRequest err: ", err)
		return errors.New(InternalServerError)
	}
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New(InvalidToken)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Patch DefaultClient.Do err: ", err)
		return errors.New(InternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Patch readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}
