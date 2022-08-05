package domain

import "gomp/dto"

type Jobs struct {
	ID        int    `json:"id" gorm:"column:job_id"`
	Title     string `json:"title"`
	City      string `json:"city"`
	Status    string `json:"status"`
	CompanyID uint   `json:"company_id" gorm:"company_id"`
}

type JobsRepository interface {
	FindAll() ([]Jobs, error)
	FindByID(int) (*Jobs, error)
}

func (j Jobs) convertJobStatus() string {
	jobStatus := "open"
	if j.Status == "0" {
		jobStatus = "closed"
	}
	return jobStatus
}

func (j Jobs) ToDTO() dto.JobsResponse {
	return dto.JobsResponse{
		ID:        j.ID,
		Title:     j.Title,
		City:      j.City,
		Status:    j.convertJobStatus(),
		CompanyID: j.CompanyID,
	}
}
