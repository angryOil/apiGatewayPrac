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

func (r TodoRequester) GetTodoDetail(ctx context.Context, id int) (domain.Todo, error) {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return domain.Todo{}, errors.New("invalid token")
	}
	re, err := http.NewRequest("GET", fmt.Sprintf("%s/%d", r.todoUrl, id), nil)
	if err != nil {
		log.Println("make NewRequest err: ", err)
		return domain.Todo{}, errors.New("internal server error")
	}
	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("request defaultClient do err: ", err)
		return domain.Todo{}, errors.New("internal server error")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return domain.Todo{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("read todoResponse err: ", err)
			return domain.Todo{}, errors.New("internal server error")
		}
		log.Println("response is not ok or nonFound", string(readBody))
		return domain.Todo{}, errors.New("internal server error")
	}

	var result domain.Todo
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("todo json decode err: ", err)
		return domain.Todo{}, errors.New("internal server error")
	}
	return result, nil
}

func (r TodoRequester) UpdateTodo(ctx context.Context, td domain.Todo) error {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New("invalid token")
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(td)
	if err != nil {
		log.Println("update todo json encode err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("PUT", fmt.Sprintf("%s", r.todoUrl), &buf)
	if err != nil {
		log.Println("make newRequest err: ", err)
		return errors.New("internal server error")
	}

	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("update todo client do fail:", err)
		return errors.New("internal server error")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("res update readBody fail: ", err)
			return errors.New("internal server error")
		}
		if resp.StatusCode == http.StatusBadRequest {
			return errors.New(string(readBody))
		}
		log.Println("update todo not ok, not badRequest response is: ", string(readBody))
		return errors.New("internal server error")
	}
	return nil
}

func (r TodoRequester) DeleteTodo(ctx context.Context, id int) error {
	token, ok := common.TokenFromContext(ctx)
	if !ok {
		return errors.New("invalid token")
	}

	re, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", r.todoUrl, id), nil)
	if err != nil {
		log.Println("delete todo make request fail:", err)
		return errors.New("server internal error")
	}
	re.Header.Add("token", token)

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("delete todo client do fail:", err)
		return errors.New("server internal error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("delete response readBody fail: ", err)
			return errors.New("internal server error")
		}
		if resp.StatusCode == http.StatusBadRequest {
			return errors.New(string(readBody))
		}
		log.Println("delete todo response statusCode is not ok,badRequest", string(readBody))
		return errors.New("internal server error")
	}

	return nil
}

type listResDto struct {
	Contents []domain.Todo `json:"contents"`
	TotalCnt int           `json:"total_content"`
}
