package dto

// Paging be used when paging
type PaginationRequestDTO struct {
	Offset int `json:"offset,default=10" form:"offset,default=10"`
	Page   int `json:"page,default=1" form:"page,default=1"`
}

// Paging be used when paging
type PaginationResponseDTO struct {
	Offset int   `json:"offset" form:"offset"`
	Page   int   `json:"page" form:"page"`
	Total  int64 `json:"total" form:"total"`
}

type ListDTO[T any] struct {
	List []T                   `json:"list" form:"list"`
	Page PaginationResponseDTO `json:"page" form:"page"`
}
