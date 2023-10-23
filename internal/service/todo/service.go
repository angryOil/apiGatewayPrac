package todo

import (
	"apiGateway/internal/cli"
	"apiGateway/internal/domain"
	"context"
)

type Service struct {
	tr cli.TodoRequester
}

func NewService(tr cli.TodoRequester) Service {
	return Service{tr: tr}
}

func (s Service) CreateTodo(ctx context.Context, td domain.Todo) error {
	err := s.tr.CreateTodo(ctx, td)
	return err
}
