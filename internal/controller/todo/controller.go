package todo

import (
	"apiGateway/internal/controller/todo/req"
	"apiGateway/internal/service/todo"
	"context"
)

type Controller struct {
	s todo.Service
}

func NewController(s todo.Service) Controller {
	return Controller{s: s}
}

func (c Controller) CreateTodo(ctx context.Context, ct req.CreateTodoDto) error {
	td := ct.ToDomain()
	err := c.s.CreateTodo(ctx, td)
	return err
}
