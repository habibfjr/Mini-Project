package dto

type Pagination struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	TotalRows int64       `json:"total_rows"`
	FromRow   int         `json:"from_row"`
	ToRow     int         `json:"to_row"`
	Rows      interface{} `json:"rows"`
}
