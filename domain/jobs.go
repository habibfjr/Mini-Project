package domain

import "gomp/dto"

type Jobs struct {
	ID        int    `json:"id" gorm:"column:job_id"`
	Title     string `json:"title" gorm:"column:title"`
	City      string `json:"city" gorm:"column:city"`
	Status    string `json:"status" gorm:"column:status"`
	CompanyID uint   `json:"company_id" gorm:"column:company_id"`
}

type JobsRepository interface {
	FindAll() ([]Jobs, error)
	FindByID(int) (*Jobs, error)
	AddJob(Jobs) (*Jobs, error)
	UpdateJob(int, Jobs) (*Jobs, error)
	DeleteJob(int) (*Jobs, error)
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
