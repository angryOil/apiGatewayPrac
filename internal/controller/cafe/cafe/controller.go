package cafe

import (
	req3 "apiGateway/internal/controller/cafe/cafe/req"
	res2 "apiGateway/internal/controller/cafe/cafe/res"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/cafe/cafe"
	req2 "apiGateway/internal/service/cafe/cafe/req"
	"context"
)

type Controller struct {
	s cafe.Service
}

func NewController(s cafe.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, createDto req3.CreateCafeDto) error {
	err := c.s.Create(ctx, req3.CreateCafeDto{
		Name:        createDto.Name,
		Description: createDto.Description,
	})
	return err
}

func (c Controller) GetList(ctx context.Context, reqPage page2.ReqPage) ([]res2.CafeListDto, int, error) {
	list, cnt, err := c.s.GetList(ctx, reqPage)
	if err != nil {
		return []res2.CafeListDto{}, 0, err
	}
	dto := make([]res2.CafeListDto, len(list))
	for i, l := range list {
		dto[i] = res2.CafeListDto{
			Id:   l.Id,
			Name: l.Name,
		}
	}
	return dto, cnt, nil
}

func (c Controller) GetDetail(ctx context.Context, id int) (res2.CafeDetailDto, error) {
	detail, err := c.s.GetDetail(ctx, id)
	if err != nil {
		return res2.CafeDetailDto{}, err
	}
	return res2.CafeDetailDto{
		Id:          detail.Id,
		Name:        detail.Name,
		Description: detail.Description,
	}, nil
}

func (c Controller) Patch(ctx context.Context, id int, p req3.PatchDto) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          id,
		Name:        p.Name,
		Description: p.Description,
	})
	return err
}
