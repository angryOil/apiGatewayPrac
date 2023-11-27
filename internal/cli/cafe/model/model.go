package model

import "apiGateway/internal/domain/cafe"

type Cafe struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Cafe) ToDomain() cafe.Cafe {
	return cafe.NewCafeBuilder().
		Id(c.Id).
		Name(c.Name).
		Description(c.Description).
		Build()
}

func ToDomainList(list []Cafe) []cafe.Cafe {
	domains := make([]cafe.Cafe, len(list))
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
