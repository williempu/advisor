package student

import (
	"database/sql"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StudentRepo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{DB: db}
}

// GetByStdId returns a valid student record from database
func (s *StudentRepo) GetByStdId(studentId string) (*Student, error) {
	var student Student
	sql := "SELECT id, student_name FROM students WHERE student_id = $1"
	row := s.DB.QueryRow(sql, studentId)
	if err := row.Scan(&student.Id, &student.Name); err != nil {
		return nil, err
	}
	return &student, nil
}
