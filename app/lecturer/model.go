package lecturer

import "database/sql"

type Lecturer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Quota int    `json:"quota"`
}

type LecturerRepo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *LecturerRepo {
	return &LecturerRepo{DB: db}
}

// GetAll returns all lecturer records
func (r *LecturerRepo) GetAll() ([]*Lecturer, error) {
	var lecturers []*Lecturer
	sql := "SELECT id, name, quota FROM lecturers"
	rows, err := r.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lecturer Lecturer
		if err := rows.Scan(&lecturer.Id, &lecturer.Name, &lecturer.Quota); err != nil {
			return nil, err
		}
		lecturers = append(lecturers, &lecturer)
	}

	return lecturers, nil
}

// GetById returns lecturer record with certain id
func (r *LecturerRepo) GetById(id int) (*Lecturer, error) {
	var lecturer Lecturer
	sql := "SELECT id, name, quota FROM lecturers WHERE id = $1"
	row := r.DB.QueryRow(sql, id)
	if err := row.Scan(&lecturer.Id, &lecturer.Name, &lecturer.Quota); err != nil {
		return nil, err
	}
	return &lecturer, nil
}
