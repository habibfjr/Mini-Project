package app

import (
	"gomp/dto"
	"gomp/service"
	"gomp/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	service service.JobsService
}

func PaginationReq(c *gin.Context) *dto.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	return &dto.Pagination{Limit: limit, Page: page}
}

func getCurrentJWT(c *gin.Context) int {
	cUser := c.MustGet("currentUser").(users.Users)
	userID := cUser.ID
	return userID
}

func (jh *JobsHandler) getAll(c *gin.Context) {
	userID := getCurrentJWT(c)
	pagination := PaginationReq(c)
	jobs, err := jh.service.GetAllJobs(*pagination, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) getJobsByID(c *gin.Context) {

	id := c.Param("id")
	getId, _ := strconv.Atoi(id)
	userID := getCurrentJWT(c)
	jobs, err := jh.service.GetJobsByID(getId, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) createJob(c *gin.Context) {
	var input dto.NewJob
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	userID := getCurrentJWT(c)
	jobs, err := jh.service.CreateJob(input, userID)
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
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userID := getCurrentJWT(c)
	jobs, err := jh.service.UpdateJob(getId, input, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (jh *JobsHandler) deleteJob(c *gin.Context) {
	id := c.Param("id")
	getId, _ := strconv.Atoi(id)
	userID := getCurrentJWT(c)
	_, err := jh.service.DeleteJob(getId, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, "successfully deleted job")
}
