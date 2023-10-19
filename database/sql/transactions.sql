DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    date_selected TIMESTAMP NOT NULL,
    student_id INT,
    lecturer_id INT,
    CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES students(id),
    CONSTRAINT fk_lecture FOREIGN KEY(lecturer_id) REFERENCES lecturers(id)
);
