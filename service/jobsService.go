package service

import (
	"gomp/domain"
	"gomp/dto"
)

type JobsService interface {
	GetAllJobs() ([]dto.JobsResponse, error)
	GetJobsByID(int) (*dto.JobsResponse, error)
	CreateJob(dto.NewJob) (*dto.JobsResponse, error)
	UpdateJob(int, dto.NewJob) (*dto.JobsResponse, error)
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

func (s DefaultJobsService) GetJobsByID(jobsID int) (*dto.JobsResponse, error) {
	j, err := s.repo.FindByID(jobsID)
	if err != nil {
		return nil, err
	}

	response := j.ToDTO()

	return &response, nil
}

func (s DefaultJobsService) CreateJob(nj dto.NewJob) (*dto.JobsResponse, error) {
	j := domain.Jobs{}
	j.Title = nj.Title
	j.City = nj.City
	j.Status = nj.Status
	j.CompanyID = nj.CompanyID

	jobs, err := s.repo.AddJob(j)
	if err != nil {
		return nil, err
	}

	res := jobs.ToDTO()

	return &res, nil
}

func (s DefaultJobsService) UpdateJob(id int, uj dto.NewJob) (*dto.JobsResponse, error) {
	j := domain.Jobs{}
	j.Title = uj.Title
	j.City = uj.City
	j.Status = uj.Status
	j.CompanyID = uj.CompanyID

	jobs, err := s.repo.UpdateJob(id, j)
	if err != nil {
		return nil, err
	}

	res := jobs.ToDTO()

	return &res, nil
}

func NewJobsService(repository domain.JobsRepository) DefaultJobsService {
	return DefaultJobsService{repository}
}
