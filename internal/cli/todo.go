package cli

import (
	"apiGateway/internal/domain"
	"apiGateway/internal/page"
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

type TodoRequester struct {
	todoUrl string
}

func NewTodoRequester(todoUrl string) TodoRequester {
	return TodoRequester{todoUrl: todoUrl}
}

func (r TodoRequester) CreateTodo(ctx context.Context, td domain.Todo) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(td)
	if err != nil {
		return err
	}

	re, err := http.NewRequest("POST", r.todoUrl, &buf)
	if err != nil {
		return err
	}

	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New("invalid token")
	}
	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("todo request decode err: ", err)
			return errors.New("internal server error")
		}
		if 400 <= resp.StatusCode && resp.StatusCode < 500 {
			return errors.New(string(data))
		}
		log.Println("todo create err:", err)
		return errors.New("internal server error")
	}
	return nil
}

func (r TodoRequester) GetTodoList(ctx context.Context, reqPage page.ReqPage) ([]domain.Todo, int, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return []domain.Todo{}, 0, errors.New("invalid token")
	}
	re, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d&size=%d", r.todoUrl, reqPage.Page, reqPage.Size), nil)
	if err != nil {
		log.Println("requestMake err: " + err.Error())
		return []domain.Todo{}, 0, errors.New("internal server error")
	}
	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("request err: ", err)
		return []domain.Todo{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("readBody err", err)
		}
		log.Println("getTodoListNotOk:", readBody)
		return []domain.Todo{}, 0, errors.New("todo server error")
	}
	var results listResDto
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		log.Println("json decode err: ", err)
		return []domain.Todo{}, 0, errors.New("internal server error")
	}
	return results.Contents, results.TotalCnt, nil
}

type listResDto struct {
	Contents []domain.Todo `json:"contents"`
	TotalCnt int           `json:"total_content"`
}
