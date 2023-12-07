package req

type PatchMember struct {
	MemberId int
	CafeId   int
	Nickname string
}

func (p PatchMember) ToDto() PatchMemberDto {
	return PatchMemberDto{Nickname: p.Nickname}
}

type PatchMemberDto struct {
	Nickname string `json:"nickname"`
}
