package domain

import (
	"gomp/dto"
	"gomp/logger"
	"math"

	"gorm.io/gorm"
)

type JobsRepositoryDB struct {
	db *gorm.DB
}

func NewJobsRepositoryDB(client *gorm.DB) JobsRepositoryDB {

	return JobsRepositoryDB{client}
}

func (jr JobsRepositoryDB) FindAll(p dto.Pagination) (dto.Pagination, error) {
	var pag dto.Pagination
	tr, totalPages, fromRow, toRow := 0, 0, 0, 0 // pagination attribute
	offset := p.Page * p.Limit

	var jobs []Jobs
	var job Jobs

	errFind := jr.db.Limit(p.Limit).Offset(offset).Find(&jobs).Error
	if errFind != nil {
		return pag, errFind
	}
	p.Rows = jobs
	// count data
	totalRows := int64(tr)
	errCount := jr.db.Model(job).Count(&totalRows).Error
	if errCount != nil {
		return p, errCount
	}

	p.TotalRows = totalRows

	totalPages = int(math.Ceil(float64(totalRows)/float64(p.Limit))) - 1

	if p.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = p.Limit
	} else {
		if p.Page <= totalPages {
			// calculate from & to row
			fromRow = p.Page*p.Limit + 1
			toRow = (p.Page + 1) * p.Limit
		}
	}

	if toRow > tr {
		// set to row with total rows
		toRow = tr
	}

	p.FromRow = fromRow
	p.ToRow = toRow
	return p, nil
}

func (jr JobsRepositoryDB) FindByID(id int) (*Jobs, error) {
	var j Jobs
	query := jr.db.First(&j, "job_id = ?", id)
	err := query.Error

	if err != nil {
		logger.Error("error fetch data to customer table " + err.Error())
		return nil, err
	}

	return &j, nil

}

func (jr JobsRepositoryDB) AddJob(j Jobs) (*Jobs, error) {
	query := jr.db.Create(&j)
	err := query.Error

	if err != nil {
		logger.Error("error creating data " + err.Error())
		return nil, err
	}

	return &j, nil
}

func (jr JobsRepositoryDB) UpdateJob(id int, j Jobs) (*Jobs, error) {
	up := jr.db.Model(&j).Where("job_id = ?", id).Updates(j).Error
	if up != nil {
		logger.Error("error updating data " + up.Error())
		return nil, up
	}
	err := jr.db.Model(&j).Where("job_id = ?", id).Take(&j).Error
	if err != nil {
		logger.Error("error displaying data " + err.Error())
		return nil, err
	}

	return &j, nil
}

func (jr JobsRepositoryDB) DeleteJob(id int) (*Jobs, error) {
	var j Jobs
	query := jr.db.Delete(&j, "job_id = ?", id)
	err := query.Error
	if err != nil {
		logger.Error("failed to delete data " + err.Error())
		return nil, err
	}

	return &j, nil
}
