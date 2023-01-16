package model

type Page struct {
	Offset   int `json:"offset,omitempty" binding:"required" form:"offset"`
	PageSize int `json:"pageSize,omitempty" binding:"required" form:"pageSize"`
}
