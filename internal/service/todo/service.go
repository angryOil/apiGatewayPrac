package todo

import (
	"apiGateway/internal/cli"
	"apiGateway/internal/domain"
	page2 "apiGateway/internal/page"
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

func (s Service) GetTodoList(ctx context.Context, reqPage page2.ReqPage) ([]domain.Todo, int, error) {
	return s.tr.GetTodoList(ctx, reqPage)
}
