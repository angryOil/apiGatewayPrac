package cafe

import (
	"apiGateway/internal/cli/cafe"
	req2 "apiGateway/internal/cli/cafe/req"
	"apiGateway/internal/controller/cafe/req"
	cafe2 "apiGateway/internal/domain/cafe"
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
