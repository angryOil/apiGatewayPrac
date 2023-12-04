package res

type CafeListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CafeListTotalDto struct {
	Content []CafeListDto `json:"content"`
	Total   int           `json:"total"`
}

func NewCafeListTotalDto(list []CafeListDto, total int) CafeListTotalDto {
	return CafeListTotalDto{
		Content: list,
		Total:   total,
	}
}
