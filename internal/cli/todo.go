package cli

import (
	"apiGateway/internal/domain"
	"apiGateway/internal/service/common"
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
