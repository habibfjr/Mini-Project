package domain

import (
	"gomp/logger"

	"gorm.io/gorm"
)

type JobsRepositoryDB struct {
	db *gorm.DB
}

func NewJobsRepositoryDB(client *gorm.DB) JobsRepositoryDB {

	return JobsRepositoryDB{client}
}

func (jr JobsRepositoryDB) FindAll() ([]Jobs, error) {

	var jobs []Jobs

	result := jr.db.Find(&jobs)
	err := result.Error

	if err != nil {
		logger.Error("error fetch data to customer table " + err.Error())
		return nil, err
	}

	return jobs, nil

}
