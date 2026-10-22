CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_name TEXT,
    price INT NOT NULL,
    duration INT NOT NULL,

    CONSTRAINT user_id_fk FOREIGN KEY (student_id) REFERENCES students(id)
);

INSERT INTO courses (student_id, course_name, price, duration)
VALUES (1, 'Go', 500, 3),
       (2, 'Python', 400, 2),
       (3, 'Java', 600, 4),
       (4, 'Go', 500, 3),
       (5, 'Python', 400, 2);
