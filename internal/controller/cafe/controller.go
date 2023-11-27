package cafe

import (
	"apiGateway/internal/controller/cafe/req"
	"apiGateway/internal/controller/cafe/res"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/cafe"
	"context"
)

type Controller struct {
	s cafe.Service
}

func NewController(s cafe.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, createDto req.CreateCafeDto) error {
	err := c.s.Create(ctx, req.CreateCafeDto{
		Name:        createDto.Name,
		Description: createDto.Description,
	})
	return err
}

func (c Controller) GetList(ctx context.Context, reqPage page2.ReqPage) ([]res.CafeListDto, int, error) {
	list, cnt, err := c.s.GetList(ctx, reqPage)
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
	return dto, cnt, nil
}
