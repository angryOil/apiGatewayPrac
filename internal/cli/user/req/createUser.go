package req

type CreateUser struct {
	Email    string
	Password string
}

func (c CreateUser) ToCreateUserDto() CreateUserDto {
	return CreateUserDto{
		Email:    c.Email,
		Password: c.Password,
	}
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
