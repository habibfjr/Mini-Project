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
	getId, _ := strconv.Atoi(id)
	jobs, err := jh.service.GetJobsByID(getId)

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

func (jh *JobsHandler) updateJob(c *gin.Context) {
	id := c.Param("id")
	getId, _ := strconv.Atoi(id)
	var input dto.NewJob
	err := c.ShouldBindJSON(&input)
	jobs, _ := jh.service.UpdateJob(getId, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) deleteJob(c *gin.Context) {
	id := c.Param("id")
	getId, _ := strconv.Atoi(id)
	_, err := jh.service.DeleteJob(getId)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, "successfully deleted job")
}
