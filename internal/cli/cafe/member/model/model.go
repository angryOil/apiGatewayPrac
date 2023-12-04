package model

import "apiGateway/internal/domain/cafe/member"

type Member struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (m Member) ToDomain() member.Member {
	return member.NewMemberBuilder().
		Id(m.Id).
		UserId(m.UserId).
		Nickname(m.Nickname).
		CreatedAt(m.CreatedAt).
		Build()
}

type MyCafeListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MyCafeListTotalDto struct {
	Contents []MyCafeListDto `json:"contents"`
	Total    int             `json:"total"`
}
