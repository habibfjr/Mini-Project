package dto

type JobsResponse struct {
	ID        int    `json:"id" gorm:"job_id"`
	Title     string `json:"title"`
	City      string `json:"city"`
	Status    string `json:"status"`
	CompanyID uint   `json:"company_id" gorm:"company_id"`
}
