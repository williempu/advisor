package lecturer

import (
	"advisor/lib/response"
	"net/http"
)

type Controller struct {
	Repo *LecturerRepo
	JSON response.JSONResponse
}

func NewController(repo *LecturerRepo) *Controller {
	return &Controller{Repo: repo}
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	lecturers, err := c.Repo.GetAll()
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON.HttpStatus(w, http.StatusOK, lecturers)
}
