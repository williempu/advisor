package student

import "net/http"

type Controller struct {
	Repo *StudentRepo
}

func NewController(repo *StudentRepo) *Controller {
	return &Controller{Repo: repo}
}

// checks if student is registered in DB
func (c *Controller) CheckStudent(w http.ResponseWriter, r *http.Request) {
	// TODO check student existence -> transaction alongside lecturer
}
