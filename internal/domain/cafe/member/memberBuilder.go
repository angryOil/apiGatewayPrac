package member

var _ Builder = (*builder)(nil)

func NewMemberBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	UserId(userId int) Builder
	Nickname(nickname string) Builder
	CreatedAt(createdAt string) Builder

	Build() Member
}

type builder struct {
	id        int
	cafeId    int
	userId    int
	nickname  string
	createdAt string
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
	return b
}

func (b *builder) UserId(userId int) Builder {
	b.userId = userId
	return b
}

func (b *builder) Nickname(nickname string) Builder {
	b.nickname = nickname
	return b
}

func (b *builder) CreatedAt(createdAt string) Builder {
	b.createdAt = createdAt
	return b
}

func (b *builder) Build() Member {
	return &member{
		id:        b.id,
		cafeId:    b.cafeId,
		userId:    b.userId,
		nickname:  b.nickname,
		createdAt: b.createdAt,
	}
}
