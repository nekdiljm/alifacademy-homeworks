SELECT students.first_name, students.last_name, SUM(courses.price * courses.duration) FROM students
JOIN courses ON courses.student_id = students.id
GROUP BY students.first_name, students.last_name
HAVING SUM(courses.price * courses.duration) > 500;