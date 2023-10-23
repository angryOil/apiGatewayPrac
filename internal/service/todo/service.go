package todo

import (
	"apiGateway/internal/cli"
	"apiGateway/internal/domain"
	"context"
)

type Service struct {
	tr cli.TodoRequester
}

func NewService() Service {
	return Service{}
}

func (s Service) CreateTodo(ctx context.Context, td domain.Todo) error {
	return nil
}
