CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT
);

INSERT INTO students (first_name, last_name, age)
VALUES
    ('Ali', 'Karimov', 20),
    ('Umed', 'Sharipov', 19),
    ('Amir', 'Rahimov', 21),
    ('Farzona', 'Saidova', 18),
    ('Muhammad', 'Azizov', 22);