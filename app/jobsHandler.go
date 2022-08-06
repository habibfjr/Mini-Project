package app

import (
	"gomp/dto"
	"gomp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	service service.JobsService
}

func (jh *JobsHandler) getAll(c *gin.Context) {
	jobs, err := jh.service.GetAllJobs()

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) getJobsByID(c *gin.Context) {
	// claims := r.Context().Value(userInfo)
	// logger.Info(fmt.Sprintf("claims: %v", claims))

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	jobs, err := jh.service.GetJobsByID(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) createJob(c *gin.Context) {
	var input dto.NewJob
	err := c.ShouldBindJSON(&input)
	jobs, _ := jh.service.CreateJob(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, jobs)
}
