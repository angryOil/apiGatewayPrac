package cafe

import "time"

var _ Builder = (*cafeBuilder)(nil)

func NewCafeBuilder() Builder {
	return &cafeBuilder{}
}

type Builder interface {
	Id(id int) Builder
	OwnerId(ownerId int) Builder
	Name(name string) Builder
	Description(description string) Builder
	CreatedAt(createdAt time.Time) Builder

	Build() Cafe
}
type cafeBuilder struct {
	id          int
	ownerId     int
	name        string
	description string
	createdAt   time.Time
}

func (c *cafeBuilder) Id(id int) Builder {
	c.id = id
	return c
}

func (c *cafeBuilder) OwnerId(ownerId int) Builder {
	c.ownerId = ownerId
	return c
}

func (c *cafeBuilder) Name(name string) Builder {
	c.name = name
	return c
}

func (c *cafeBuilder) Description(description string) Builder {
	c.description = description
	return c
}

func (c *cafeBuilder) CreatedAt(createdAt time.Time) Builder {
	c.createdAt = createdAt
	return c
}

func (c *cafeBuilder) Build() Cafe {
	return &cafe{
		id:          c.id,
		ownerId:     c.ownerId,
		name:        c.name,
		description: c.description,
		createdAt:   c.createdAt,
	}
}
