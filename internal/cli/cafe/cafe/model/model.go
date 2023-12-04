package model

import (
	cafe2 "apiGateway/internal/domain/cafe/cafe"
)

type Cafe struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Cafe) ToDomain() cafe2.Cafe {
	return cafe2.NewCafeBuilder().
		Id(c.Id).
		Name(c.Name).
		Description(c.Description).
		Build()
}

func ToDomainList(list []Cafe) []cafe2.Cafe {
	domains := make([]cafe2.Cafe, len(list))
	for i, l := range list {
		domains[i] = l.ToDomain()
	}
	return domains
}

type CafePage struct {
	Contents    []Cafe `json:"contents"`
	Total       int    `json:"total_content"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
}
