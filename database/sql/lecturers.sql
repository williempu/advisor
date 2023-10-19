DROP TABLE IF EXISTS lecturers;
CREATE TABLE lecturers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    quota INT NOT NULL
);
