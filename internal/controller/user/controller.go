package user

import (
	"apiGateway/internal/controller/user/req"
	"apiGateway/internal/service/user"
	req2 "apiGateway/internal/service/user/req"
	"context"
)

type Controller struct {
	s user.Service
}

func NewController(s user.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Login(ctx context.Context, u req.LoginDto) (string, error) {
	token, err := c.s.Login(ctx, req2.Login{
		Email:    u.Email,
		Password: u.Password,
	})
	return token, err
}

func (c Controller) CreateUser(ctx context.Context, cd req.CreateDto) error {
	err := c.s.CreateUser(ctx, req2.CreateUser{
		Email:    cd.Email,
		Password: cd.Password,
	})
	return err
}
