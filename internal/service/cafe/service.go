package cafe

import "apiGateway/internal/cli/cafe"

type Service struct {
	r cafe.Requester
}

func NewService(r cafe.Requester) Service {
	return Service{r: r}
}
