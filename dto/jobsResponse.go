package dto

type JobsResponse struct {
	ID        int    `json:"id" gorm:"column:job_id"`
	Title     string `json:"title" gorm:"column:title"`
	City      string `json:"city" gorm:"column:city"`
	Status    string `json:"status" gorm:"column:status"`
	CompanyID int32  `json:"company_id" gorm:"column:company_id"`
}
