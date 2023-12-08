package member

import (
	"apiGateway/internal/cli/cafe/member/model"
	"apiGateway/internal/cli/cafe/member/req"
	"apiGateway/internal/domain/cafe/member"
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

func (r Requester) GetMemberInfo(ctx context.Context, cafeId int) (member.Member, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return member.NewMemberBuilder().Build(), errors.New(InvalidToken)
	}
	reqUrl := fmt.Sprintf("%s/%d/%s/info", baseUrl, cafeId, Members)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetMemberInfo NewRequest err: ", err)
		return member.NewMemberBuilder().Build(), errors.New(InternalServerError)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetMemberInfo DefaultClient.Do err: ", err)
		return member.NewMemberBuilder().Build(), errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetMemberInfo readBody err: ", err)
			return member.NewMemberBuilder().Build(), errors.New(InternalServerError)
		}
		return member.NewMemberBuilder().Build(), errors.New(string(readBody))
	}

	var m model.Member
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("GetMemberInfo json.NewDecoder err: ", err)
		return member.NewMemberBuilder().Build(), errors.New(InternalServerError)
	}
	return m.ToDomain(), nil
}

func (r Requester) JoinCafe(ctx context.Context, c req.JoinCafe) error {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New(InvalidToken)
	}
	reqUrl := fmt.Sprintf("%s/%d/%s/join", baseUrl, c.CafeId, Members)
	dto := model.JoinCafe{Nickname: c.NickName}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(dto)
	if err != nil {
		log.Println("JoinCafe json.NewDecoder err: ", err)
		return errors.New(InternalServerError)
	}
	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("JoinCafe NewRequest err: ", err)
		return errors.New(InternalServerError)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("JoinCafe DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("JoinCafe readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}

	return nil
}

func (r Requester) PatchMember(ctx context.Context, p req.PatchMember) error {
	reUrl := fmt.Sprintf("%s/%d/%s/%d", baseUrl, p.CafeId, Members, p.MemberId)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p.ToDto())
	if err != nil {
		log.Println("PatchMember json.NewEncode err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("PATCH", reUrl, &buf)
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New(InvalidToken)
	}
	re.Header.Add(Token, token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("PatchMember DefaultClient do err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("PatchMember readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}
