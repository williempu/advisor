package transaction

import (
	"advisor/app/lecturer"
	"advisor/app/student"
	"advisor/lib/response"
	"encoding/json"
	"net/http"
	"strconv"
)

type Controller struct {
	StudentRepo  *student.StudentRepo
	LecturerRepo *lecturer.LecturerRepo
	Repo         *TransactionRepo
	JSON         response.JSONResponse
}

func NewController(studentRepo *student.StudentRepo, lecturerRepo *lecturer.LecturerRepo, transactionRepo *TransactionRepo) *Controller {
	return &Controller{
		StudentRepo:  studentRepo,
		LecturerRepo: lecturerRepo,
		Repo:         transactionRepo,
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	// check student and lecturer existence first
	// then validate if student has already selected the lecturer
	// proceed if all good

	// parse request -> check -> done
	var body struct {
		StudentId  string `json:"studentId"`
		LecturerId string `json:"lecturerId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	// check if student id is real, get the student
	student, err := c.StudentRepo.GetByStdId(body.StudentId)
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	lecturerId, err := strconv.Atoi(body.LecturerId)
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}
	lecturer, err := c.LecturerRepo.GetById(lecturerId)
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	// before really creating, finalize again
	// check for transaction, if any exists, return it as OK status
	transaction, err := c.Repo.GetByStudentId(student.Id)
	if transaction != nil || err == nil {
		c.JSON.HttpStatus(w, http.StatusOK, transaction)
		return
	}

	tData := TransactionData{
		StudentId:  student.Id,
		LecturerId: lecturer.Id,
	}

	if err := c.Repo.Create(tData); err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON.HttpStatus(w, http.StatusOK, "Success")
}

func (c *Controller) GetLecturerQuota(w http.ResponseWriter, r *http.Request) {
	quotas, err := c.Repo.GetLecturerQuota()
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON.HttpStatus(w, http.StatusOK, quotas)
}

func (c *Controller) GetByStudentId(w http.ResponseWriter, r *http.Request) {
	// get studentId, put inside a struct
	var transactionInput TransactionInput
	if err := json.NewDecoder(r.Body).Decode(&transactionInput); err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	student, err := c.StudentRepo.GetByStdId(transactionInput.StudentId)
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	// unique case: if no transaction found, return as success
	transaction, err := c.Repo.GetByStudentId(student.Id)
	if err != nil {
		c.JSON.HttpStatus(w, http.StatusOK, student)
		return
	}

	c.JSON.HttpStatus(w, http.StatusOK, transaction)
}
