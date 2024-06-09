package define

type SortOrderType string

const (
	OrderAsc  SortOrderType = "asc"
	OrderDesc SortOrderType = "desc"
)

type Status int

const (
	ActiveStatus   Status = 1
	InActiveStatus Status = 2
)
