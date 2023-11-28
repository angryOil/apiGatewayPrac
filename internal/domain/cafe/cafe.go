package cafe

import (
	"apiGateway/internal/domain/cafe/vo"
	"errors"
	"time"
)

var _ Cafe = (*cafe)(nil)

type Cafe interface {
	ValidCafeFiled() error
	ValidCreate() error
	ValidUpdate() error
	Update(name, description string) Cafe
	VerifyUpdate() error
	GetOwnerId() int

	ToDetail() vo.Detail
	ToCafeListInfo() vo.CafeListInfo
	UpdateCafeInfo() vo.UpdateCafe
}

type cafe struct {
	id          int
	ownerId     int
	name        string
	description string
	createdAt   time.Time
}

func (c *cafe) UpdateCafeInfo() vo.UpdateCafe {
	return vo.UpdateCafe{
		Id:          c.id,
		OwnerId:     c.ownerId,
		Name:        c.name,
		Description: c.description,
		CreatedAt:   c.createdAt,
	}
}

func (c *cafe) ToCafeListInfo() vo.CafeListInfo {
	return vo.CafeListInfo{
		Id:   c.id,
		Name: c.name,
	}
}

func (c *cafe) ToDetail() vo.Detail {
	return vo.Detail{
		Id:          c.id,
		Name:        c.name,
		Description: c.description,
	}
}

func (c *cafe) GetOwnerId() int {
	return c.ownerId
}

const (
	EmptyName = "name is empty"
)

func (c *cafe) ValidCafeFiled() error {
	if c.name == "" {
		return errors.New(EmptyName)
	}
	if c.ownerId == 0 {
		return errors.New("owner id is zero")
	}
	if c.id == 0 {
		return errors.New("id is zero")
	}
	return nil
}

func (c *cafe) ValidCreate() error {
	if c.name == "" {
		return errors.New(EmptyName)
	}
	return nil
}

func (c *cafe) Update(name, description string) Cafe {
	c.name = name
	c.description = description
	return c
}

const (
	NotVerifyId       = "cafe id is zero"
	NotVerifyOwnerId  = "owner id is zero"
	NotVerifyCafeName = "cafe name is empty"
)

func (c *cafe) ValidUpdate() error {
	if c.id < 1 {
		return errors.New(NotVerifyId)
	}
	if c.name == "" {
		return errors.New(NotVerifyCafeName)
	}
	return nil
}

func (c *cafe) VerifyUpdate() error {
	if c.id < 1 {
		return errors.New(NotVerifyId)
	}
	if c.ownerId < 1 {
		return errors.New(NotVerifyOwnerId)
	}
	if c.name == "" {
		return errors.New(NotVerifyCafeName)
	}

	return nil
}
