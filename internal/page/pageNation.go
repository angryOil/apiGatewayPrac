package page

import (
	"net/http"
	"strconv"
)

var minPage = 0
var minSize = 10
var maxSize = 50

type ReqPage struct {
	Page   int
	Size   int
	Offset int
}

func NewReqPage(page int, size int) ReqPage {
	rp := ReqPage{}
	if page <= minPage {
		rp.Page = 0
	} else {
		rp.Page = page
	}
	if size <= minSize {
		rp.Size = minSize
	} else if size >= maxSize {
		rp.Size = maxSize
	} else {
		rp.Size = size
	}
	rp.Offset = rp.Page * rp.Size
	return rp
}

type Pagination[T any] struct {
	Contents    []T `json:"contents"`
	Total       int `json:"total_content"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

// page 는 0번째 page 부터 시작합니다.

func GetPagination[T any](contents []T, rp ReqPage, totalCount int) Pagination[T] {
	return Pagination[T]{
		Contents:    contents,
		Total:       totalCount,
		CurrentPage: rp.Page,
		LastPage:    (totalCount - 1) / rp.Size,
	}
}

func GetPageReqByRequest(r *http.Request) ReqPage {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	return NewReqPage(page, size)
}
