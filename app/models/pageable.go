package models

const (
	DirectionAsc  = "asc"
	DirectionDesc = "desc"
	MaxPageSize   = 1000
)

type Pageable struct {
	Page        int64
	Size        int64
	Sort        []SortOption
}

type SortOption struct {
	Property  string
	Direction string
}

type Page struct {
	TotalPages int64
	TotalElements int64
	PageNumber int64
	PageSize int64
	NumberOfElements int64
	Items interface{}
}