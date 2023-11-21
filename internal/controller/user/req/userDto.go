package req

import (
	"apiGateway/internal/domain/user"
)

type CreateDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cd CreateDto) ToDomain() user.User {
	return user.NewBuilder().
		Email(cd.Email).
		Password(cd.Password).
		Role([]string{"user"}).
		Build()
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ld LoginDto) ToDomain() user.User {
	return user.NewBuilder().
		Email(ld.Email).
		Password(ld.Password).
		Build()
}
