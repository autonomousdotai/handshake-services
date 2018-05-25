package bean

type Pagination struct {
	PageSize      int
	Page          int
	Total         int
	Items         interface{}
}