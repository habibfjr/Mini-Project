package service

import (
	"gomp/domain"
	"gomp/dto"
)

type JobsService interface {
	GetAllJobs() ([]dto.JobsResponse, error)
	// GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultJobsService struct {
	repo domain.JobsRepository
}

func (s DefaultJobsService) GetAllJobs() ([]dto.JobsResponse, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.JobsResponse
	for _, customer := range customers {
		response = append(response, customer.ToDTO())
	}

	return response, nil
}

func NewJobsService(repository domain.JobsRepository) DefaultJobsService {
	return DefaultJobsService{repository}
}
