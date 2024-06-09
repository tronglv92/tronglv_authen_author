package request

type SortOrderReq struct {
	SortBy    string `form:"sort_by,optional"`
	SortOrder string `form:"sort_order,optional"`
}

type PaginationReq struct {
	Limit int `form:"limit,optional"`
	Page  int `form:"page,optional"`
}
