package member

import (
	"apiGateway/internal/controller/cafe/member/res"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/cafe/member"
	"context"
)

type Controller struct {
	s member.Service
}

func NewController(s member.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetMyCafeList(ctx context.Context, reqPage page2.ReqPage) ([]res.CafeListDto, int, error) {
	list, total, err := c.s.GetMyCafeList(ctx, reqPage)
	if err != nil {
		return []res.CafeListDto{}, 0, err
	}
	dto := make([]res.CafeListDto, len(list))
	for i, l := range list {
		dto[i] = res.CafeListDto{
			Id:   l.Id,
			Name: l.Name,
		}
	}
	return dto, total, nil
}
