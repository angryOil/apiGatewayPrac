package user

import "time"

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	Email(email string) Builder
	Password(password string) Builder
	Role(role []string) Builder
	CreatedAt(createdAt time.Time) Builder

	Build() User
}

type builder struct {
	id        int
	email     string
	password  string
	role      []string
	createdAt time.Time
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) Email(email string) Builder {
	b.email = email
	return b
}

func (b *builder) Password(password string) Builder {
	b.password = password
	return b
}

func (b *builder) Role(role []string) Builder {
	b.role = role
	return b
}

func (b *builder) CreatedAt(createdAt time.Time) Builder {
	b.createdAt = createdAt
	return b
}

func (b *builder) Build() User {
	return &user{
		id:        b.id,
		email:     b.email,
		password:  b.password,
		role:      b.role,
		createdAt: b.createdAt,
	}
}
