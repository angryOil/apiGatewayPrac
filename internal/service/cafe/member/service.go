package member

import (
	"apiGateway/internal/cli/cafe/member"
	page2 "apiGateway/internal/page"
	"apiGateway/internal/service/cafe/member/res"
	"context"
)

type Service struct {
	r member.Requester
}

func NewService(r member.Requester) Service {
	return Service{r: r}
}

func (s Service) GetMyCafeList(ctx context.Context, reqPage page2.ReqPage) ([]res.GetMyCafeList, int, error) {
	list, total, err := s.r.GetMyCafeList(ctx, reqPage)
	if err != nil {
		return []res.GetMyCafeList{}, 0, err
	}
	dto := make([]res.GetMyCafeList, len(list))
	for i, l := range list {
		dto[i] = res.GetMyCafeList{
			Id:   l.Id,
			Name: l.Name,
		}
	}
	return dto, total, nil
}
