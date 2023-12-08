package req

type Login struct {
	Email    string
	Password string
}

func (l Login) ToLoginDto() LoginDto {
	return LoginDto{
		Email:    l.Email,
		Password: l.Password,
	}
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
