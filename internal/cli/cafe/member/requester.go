package member

import (
	"apiGateway/internal/cli/cafe/member/model"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/common"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8083/cafes"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

const (
	Token               = "token"
	Members             = "members"
	InvalidToken        = "invalid token"
	InternalServerError = "internal server error"
)

func (r Requester) GetMyCafeList(ctx context.Context, reqPage page2.ReqPage) ([]model.MyCafeListDto, int, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return []model.MyCafeListDto{}, 0, errors.New(InvalidToken)
	}
	reUrl := fmt.Sprintf("%s/%s/my?page=%d&size=%d", baseUrl, Members, reqPage.Page, reqPage.Size)

	re, err := http.NewRequest("GET", reUrl, nil)
	if err != nil {
		log.Println("GetMyCafeList NewRequester err: ", err)
		return []model.MyCafeListDto{}, 0, errors.New(InternalServerError)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetMyCafeList DefaultClient.Do err: ", err)
		return []model.MyCafeListDto{}, 0, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetMyCafeList readBody err: ", err)
			return []model.MyCafeListDto{}, 0, errors.New(InternalServerError)
		}
		return []model.MyCafeListDto{}, 0, errors.New(string(readBody))
	}
	var dto model.MyCafeListTotalDto
	err = json.NewDecoder(resp.Body).Decode(&dto)

	if err != nil {
		log.Println("GetMyCafeList json.NewDecode err: ", err)
		return []model.MyCafeListDto{}, 0, errors.New(InternalServerError)
	}
	return dto.Contents, dto.Total, nil
}
