package dto

type NewJob struct {
	Title     string `json:"title"`
	City      string `json:"city"`
	Status    string `json:"status"`
	CompanyID uint   `json:"company_id"`
}
