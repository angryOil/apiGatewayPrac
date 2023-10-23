package res

import (
	"apiGateway/internal/domain"
	"time"
)

type ListDto struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	OrderNum  int    `json:"order_num"`
	CreatedAt string `json:"created_at"`
	IsDeleted bool   `json:"is_done"`
}

func ToListDtoList(todos []domain.Todo) []ListDto {
	listDto := make([]ListDto, len(todos))
	for i, todo := range todos {
		listDto[i] = ToListDto(todo)
	}
	return listDto
}

func ToListDto(todo domain.Todo) ListDto {
	return ListDto{
		Id:        todo.Id,
		Title:     todo.Title,
		OrderNum:  todo.OrderNum,
		CreatedAt: convertTimeToString(todo.CreatedAt),
		IsDeleted: todo.IsDeleted,
	}
}

type DetailDto struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	OrderNum  int    `json:"order_num,omitempty"`
	IsDeleted bool   `json:"is_done,omitempty"`
}

func ToDetailDto(todo domain.Todo) DetailDto {
	return DetailDto{
		Id:        todo.Id,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Content:   todo.Content,
		CreatedAt: convertTimeToString(todo.CreatedAt),
		OrderNum:  todo.OrderNum,
		IsDeleted: todo.IsDeleted,
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}
