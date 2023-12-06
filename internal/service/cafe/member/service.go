package member

import (
	"apiGateway/internal/cli/cafe/member"
	"apiGateway/internal/cli/cafe/member/req"
	member2 "apiGateway/internal/domain/cafe/member"
	page2 "apiGateway/internal/page"
	res2 "apiGateway/internal/service/cafe/cafe/res"
	req2 "apiGateway/internal/service/cafe/member/req"
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

func (s Service) GetMemberInfo(ctx context.Context, cafeId int) (res2.GetMemberInfo, error) {
	d, err := s.r.GetMemberInfo(ctx, cafeId)
	if err != nil {
		return res2.GetMemberInfo{}, err
	}
	v := d.ToInfo()
	return res2.GetMemberInfo{
		Id:        v.Id,
		UserId:    v.UserId,
		Nickname:  v.Nickname,
		CreatedAt: v.CreatedAt,
	}, nil
}

func (s Service) JoinCafe(ctx context.Context, c req2.JoinCafe) error {
	cafeId := c.CafeId
	nickName := c.NickName
	err := member2.NewMemberBuilder().
		CafeId(cafeId).
		Nickname(nickName).
		Build().ValidCreate()
	if err != nil {
		return err
	}
	err = s.r.JoinCafe(ctx, req.JoinCafe{
		CafeId:   cafeId,
		NickName: nickName,
	})
	return err
}
