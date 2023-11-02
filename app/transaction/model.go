package transaction

import (
	"database/sql"
	"time"
)

type Transaction struct {
	Id           string    `json:"id"`
	Date         time.Time `json:"dateSelected"`
	StudentId    int       `json:"studentId"`
	StudentName  string    `json:"studentName"`
	LecturerId   int       `json:"lecturerId"`
	LecturerName string    `json:"lecturerName"`
	DateString   string    `json:"dateString"`
}

type TransactionInput struct {
	StudentId string `json:"studentId"`
}

type LecturerQuota struct {
	LecturerName string `json:"name"`
	LecturerId   int    `json:"id"`
	Quota        int    `json:"quota"`
	Remaining    int    `json:"remaining"`
}

type TransactionData struct {
	StudentId  int
	LecturerId int
}

type TransactionRepo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}

// GetByStudentId returns a Transaction by studentId
func (r *TransactionRepo) GetByStudentId(studentId int) (*Transaction, error) {
	var transaction Transaction
	sql := "SELECT trx.id, trx.date_selected, trx.student_id, std.student_name, trx.lecturer_id, lect.name " +
		"FROM transactions trx " +
		"JOIN lecturers lect ON lect.id = trx.lecturer_id " +
		"LEFT JOIN students std ON std.id = trx.student_id " +
		"WHERE trx.student_id = $1"
	row := r.DB.QueryRow(sql, studentId)
	if err := row.Scan(&transaction.Id, &transaction.Date, &transaction.StudentId, &transaction.StudentName, &transaction.LecturerId, &transaction.LecturerName); err != nil {
		return nil, err
	}

	transaction.DateString = transaction.Date.Format("02 Jan 2006 15:04")
	return &transaction, nil
}

// GetLecturerQuota returns the list of lecturers data with the quota and remaining slot
func (r *TransactionRepo) GetLecturerQuota() ([]*LecturerQuota, error) {
	var quotas []*LecturerQuota
	sql := "SELECT lect.name AS name, lect.id AS id, lect.quota AS quota, COALESCE(COUNT(trx.id), 0) AS remaining FROM lecturers lect LEFT JOIN transactions trx ON lect.id = trx.lecturer_id GROUP BY lect.name, lect.id ORDER BY lect.name"

	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var quota LecturerQuota
		if err := rows.Scan(&quota.LecturerName, &quota.LecturerId, &quota.Quota, &quota.Remaining); err != nil {
			return nil, err
		}
		quotas = append(quotas, &quota)
	}

	return quotas, nil
}

func (r *TransactionRepo) Create(tData TransactionData) error {
	now := time.Now()
	sql := "INSERT INTO transactions (student_id, lecturer_id, date_selected) VALUES ($1, $2, $3)"
	if _, err := r.DB.Exec(sql, tData.StudentId, tData.LecturerId, now); err != nil {
		return err
	}
	return nil
}
