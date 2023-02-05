package dto

// Paging be used when paging
type PaginationRequestDTO struct {
	Offset int `json:"offset" form:"offset"`
	Page   int `json:"page" form:"page"`
}

// Paging be used when paging
type PaginationResponseDTO struct {
	Offset int `json:"offset" form:"offset"`
	Page   int `json:"page" form:"page"`
	Total  int `json:"total" form:"total"`
}

type ListDTO[T any] struct {
	List []T                   `json:"list" form:"list"`
	Page PaginationResponseDTO `json:"page" form:"page"`
}
