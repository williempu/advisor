DROP TABLE IF EXISTS students;
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    student_id CHAR(12) NOT NULL,
    student_name VARCHAR(100) NOT NULL
);
