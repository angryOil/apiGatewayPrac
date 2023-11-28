package cafe

import (
	"apiGateway/internal/cli/cafe"
	req2 "apiGateway/internal/cli/cafe/req"
	"apiGateway/internal/controller/cafe/req"
	cafe2 "apiGateway/internal/domain/cafe"
	page2 "apiGateway/internal/page"
	req3 "apiGateway/internal/service/cafe/req"
	"apiGateway/internal/service/cafe/res"
	"context"
)

type Service struct {
	r cafe.Requester
}

func NewService(r cafe.Requester) Service {
	return Service{r: r}
}

func (s Service) Create(ctx context.Context, c req.CreateCafeDto) error {
	name, description := c.Name, c.Description
	err := cafe2.NewCafeBuilder().
		Name(name).
		Description(description).
		Build().ValidCreate()

	if err != nil {
		return err
	}
	err = s.r.Create(ctx, req2.Create{
		Name:        name,
		Description: description,
	})
	return err
}

func (s Service) GetList(ctx context.Context, reqPage page2.ReqPage) ([]res.GetCafes, int, error) {
	domains, cnt, err := s.r.GetList(ctx, reqPage)
	if err != nil {
		return []res.GetCafes{}, 0, err
	}
	dto := make([]res.GetCafes, len(domains))
	for i, d := range domains {
		v := d.ToCafeListInfo()
		dto[i] = res.GetCafes{
			Id:   v.Id,
			Name: v.Name,
		}
	}
	return dto, cnt, nil
}

func (s Service) GetDetail(ctx context.Context, id int) (res.GetDetail, error) {
	d, err := s.r.GetDetail(ctx, id)
	if err != nil {
		return res.GetDetail{}, err
	}
	v := d.ToDetail()
	return res.GetDetail{
		Id:          v.Id,
		Name:        v.Name,
		Description: v.Description,
	}, nil
}

func (s Service) Patch(ctx context.Context, p req3.Patch) error {
	id := p.Id
	name, description := p.Name, p.Description

	err := cafe2.NewCafeBuilder().
		Id(id).
		Name(name).
		Description(description).
		Build().ValidUpdate()
	if err != nil {
		return err
	}
	err = s.r.Patch(ctx, req2.Patch{
		Id:          id,
		Name:        name,
		Description: description,
	})
	return err
}
