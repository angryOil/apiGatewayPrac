package user

import (
	"apiGateway/controller/user/req"
	"apiGateway/service/user"
	"context"
)

type Controller struct {
	s user.Service
}

func NewController(s user.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Login(ctx context.Context, u req.LoginDto) (string, error) {
	du := u.ToDomain()
	token, err := c.s.Login(ctx, du)
	return token, err
}

func (c Controller) CreateUser(ctx context.Context, cd req.CreateDto) error {
	u := cd.ToDomain()
	err := c.s.CreateUser(ctx, u)
	return err
}
