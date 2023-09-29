package types

// QueryID 查询ID
type QueryID struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}
