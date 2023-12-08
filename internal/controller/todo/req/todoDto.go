package req

import (
	"apiGateway/internal/domain/todo"
	"time"
)

type CreateTodoDto struct {
	Title    string `json:"title" example:"제목"`
	Content  string `json:"content" example:"내용"`
	OrderNum int    `json:"order_num" example:"2"`
}

func (d CreateTodoDto) ToDomain() todo.Todo {
	return todo.Todo{
		Title:    d.Title,
		Content:  d.Content,
		OrderNum: d.OrderNum,
	}
}

type UpdateTodoDto struct {
	Title     string `json:"title" example:"제목"`
	Content   string `json:"content" example:"내용"`
	OrderNum  int    `json:"order_num" example:"1"`
	IsDeleted bool   `json:"is_done" example:"false"`
}

func (d UpdateTodoDto) ToDomain(todoId int) todo.Todo {
	return todo.Todo{
		Id:            todoId,
		Title:         d.Title,
		Content:       d.Content,
		OrderNum:      d.OrderNum,
		LastUpdatedAt: time.Now(),
		IsDeleted:     d.IsDeleted,
	}
}
