package core

type PaginationResult struct {
	Records    []RecordInterface
	TotalCount int64
	Page       int
	PageSize   int
}
