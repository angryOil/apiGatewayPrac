package todo

import (
	"apiGateway/internal/cli/todo"
	todo2 "apiGateway/internal/domain/todo"
	page2 "apiGateway/internal/page"
	"context"
)

type Service struct {
	tr todo.TodoRequester
}

func NewService(tr todo.TodoRequester) Service {
	return Service{tr: tr}
}

func (s Service) CreateTodo(ctx context.Context, td todo2.Todo) error {
	err := s.tr.CreateTodo(ctx, td)
	return err
}

func (s Service) GetTodoList(ctx context.Context, reqPage page2.ReqPage) ([]todo2.Todo, int, error) {
	return s.tr.GetTodoList(ctx, reqPage)
}

func (s Service) GetTodoDetail(ctx context.Context, id int) (todo2.Todo, error) {
	resTodoDomain, err := s.tr.GetTodoDetail(ctx, id)
	return resTodoDomain, err
}

func (s Service) UpdateTodo(ctx context.Context, td todo2.Todo) error {
	err := s.tr.UpdateTodo(ctx, td)
	return err
}

func (s Service) DeleteTodo(ctx context.Context, id int) error {
	err := s.tr.DeleteTodo(ctx, id)
	return err
}
