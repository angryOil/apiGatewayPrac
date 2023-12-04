package cafe

import (
	"apiGateway/internal/cli/cafe/cafe"
	req4 "apiGateway/internal/cli/cafe/cafe/req"
	"apiGateway/internal/controller/cafe/cafe/req"
	cafe2 "apiGateway/internal/domain/cafe/cafe"
	page2 "apiGateway/internal/page"
	req3 "apiGateway/internal/service/cafe/cafe/req"
	res2 "apiGateway/internal/service/cafe/cafe/res"
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
	err = s.r.Create(ctx, req4.Create{
		Name:        name,
		Description: description,
	})
	return err
}

func (s Service) GetList(ctx context.Context, reqPage page2.ReqPage) ([]res2.GetCafes, int, error) {
	domains, cnt, err := s.r.GetList(ctx, reqPage)
	if err != nil {
		return []res2.GetCafes{}, 0, err
	}
	dto := make([]res2.GetCafes, len(domains))
	for i, d := range domains {
		v := d.ToCafeListInfo()
		dto[i] = res2.GetCafes{
			Id:   v.Id,
			Name: v.Name,
		}
	}
	return dto, cnt, nil
}

func (s Service) GetDetail(ctx context.Context, id int) (res2.GetDetail, error) {
	d, err := s.r.GetDetail(ctx, id)
	if err != nil {
		return res2.GetDetail{}, err
	}
	v := d.ToDetail()
	return res2.GetDetail{
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
	err = s.r.Patch(ctx, req4.Patch{
		Id:          id,
		Name:        name,
		Description: description,
	})
	return err
}
