package app

import (
	"encoding/json"
	"gomp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	Service service.JobsService
}

func (jh *JobsHandler) getAll(c *gin.Context) {
	jobs, err := jh.Service.GetAllJobs()

	if err != nil {
		writeResponse(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(c.Writer, http.StatusOK, jobs)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
