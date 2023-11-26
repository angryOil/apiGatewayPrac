package cafe

import (
	"apiGateway/internal/controller/cafe/req"
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
