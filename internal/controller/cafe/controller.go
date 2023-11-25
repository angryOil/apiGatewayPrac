package cafe

import "apiGateway/internal/service/cafe"

type Controller struct {
	s cafe.Service
}

func NewController(s cafe.Service) Controller {
	return Controller{s: s}
}
