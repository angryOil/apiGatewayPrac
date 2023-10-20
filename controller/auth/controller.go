package auth

import "apiGateway/jwt"

type AuthController struct {
	p jwt.Provider
}

func NewAuthController(p jwt.Provider) AuthController {
	return AuthController{p: p}
}

func (ac AuthController) Login(user) {

}
