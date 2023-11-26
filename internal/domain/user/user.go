package user

import (
	"apiGateway/internal/domain/user/vo"
	"errors"
	"net/mail"
	"time"
	"unicode/utf8"
)

var _ User = (*user)(nil)

type User interface {
	ValidCreateUser() error

	ToInfo() vo.Info
}

type user struct {
	id        int
	email     string
	password  string
	role      []string
	createdAt time.Time
}

const (
	InvalidEmail     = "invalid email"
	PasswordTooShort = "password is too short"
	PasswordTooLong  = "password is too long"
)

func (u *user) ValidCreateUser() error {
	_, err := mail.ParseAddress(u.email)
	if err != nil {
		return errors.New(InvalidEmail)
	}
	if u.password == "" || utf8.RuneCountInString(u.password) < 4 {
		return errors.New(PasswordTooShort)
	} else if utf8.RuneCountInString(u.password) > 72 {
		return errors.New(PasswordTooLong)
	}
	return nil
}

func (u *user) ToInfo() vo.Info {
	return vo.Info{
		UserId:   u.id,
		Email:    u.email,
		Password: u.password,
		Role:     u.role,
	}
}
