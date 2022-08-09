package service

import (
	"errors"
	"fmt"
	"gomp/domain"
	"gomp/dto"
	"gomp/logger"
	"gomp/users"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JobsService interface {
	GetAllJobs(dto.Pagination, int) (dto.Pagination, error)
	GetJobsByID(int, int) (*dto.JobsResponse, error)
	CreateJob(dto.NewJob, int) (*dto.JobsResponse, error)
	UpdateJob(int, dto.NewJob, int) (*dto.JobsResponse, error)
	DeleteJob(int, int) (*dto.JobsResponse, error)
}

type DefaultJobsService struct {
	repo domain.JobsRepository
}

func GetClientDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("success connect to database...")

	return db
}

func (s DefaultJobsService) GetAllJobs(p dto.Pagination, userID int) (dto.Pagination, error) {
	db := GetClientDB()
	userRepo := users.NewUserRepositoryDB(db)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return p, err
	}
	if user.Username == "" {
		return p, errors.New("user not found")
	} else {
		jobs, err := s.repo.FindAll(p)
		if err != nil {
			return jobs, err
		}

		return jobs, nil
	}
}

func (s DefaultJobsService) GetJobsByID(jobsID int, userID int) (*dto.JobsResponse, error) {
	var j *dto.JobsResponse
	db := GetClientDB()
	userRepo := users.NewUserRepositoryDB(db)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return j, err
	}
	if user.Username == "" {
		return j, errors.New("user not found")
	} else {
		j, err := s.repo.FindByID(jobsID)
		if err != nil {
			return nil, err
		}

		response := j.ToDTO()

		return &response, nil
	}
}

func (s DefaultJobsService) CreateJob(nj dto.NewJob, userID int) (*dto.JobsResponse, error) {
	var j *dto.JobsResponse
	db := GetClientDB()
	userRepo := users.NewUserRepositoryDB(db)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return j, err
	}
	if user.Username == "" {
		return j, errors.New("user not found")
	} else {
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
}

func (s DefaultJobsService) UpdateJob(id int, uj dto.NewJob, userID int) (*dto.JobsResponse, error) {
	var j *dto.JobsResponse
	db := GetClientDB()
	userRepo := users.NewUserRepositoryDB(db)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return j, err
	}
	if user.Username == "" {
		return j, errors.New("user not found")
	} else {
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
}

func (s DefaultJobsService) DeleteJob(id int, userID int) (*dto.JobsResponse, error) {
	var j *dto.JobsResponse
	db := GetClientDB()
	userRepo := users.NewUserRepositoryDB(db)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return j, err
	}
	if user.Username == "" {
		return j, errors.New("user not found")
	} else {
		jobs, err := s.repo.DeleteJob(id)
		if err != nil {
			return nil, err
		}

		res := jobs.ToDTO()

		return &res, nil
	}
}

func NewJobsService(repository domain.JobsRepository) DefaultJobsService {
	return DefaultJobsService{repository}
}
