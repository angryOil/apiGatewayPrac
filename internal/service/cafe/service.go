package cafe

import (
	"apiGateway/internal/cli/cafe"
	req2 "apiGateway/internal/cli/cafe/req"
	"apiGateway/internal/controller/cafe/req"
	cafe2 "apiGateway/internal/domain/cafe"
	page2 "apiGateway/internal/page"
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
