package todo

import (
	"apiGateway/internal/controller/todo/req"
	"apiGateway/internal/controller/todo/res"
	page2 "apiGateway/internal/page"
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

func (c Controller) GetTodoList(ctx context.Context, reqPage page2.ReqPage) (page2.Pagination[res.ListDto], error) {

	contents, totalCnt, err := c.s.GetTodoList(ctx, reqPage)
	if err != nil {
		return page2.Pagination[res.ListDto]{}, err
	}
	contentDtoList := res.ToListDtoList(contents)
	return page2.GetPagination(contentDtoList, reqPage, totalCnt), nil
}
