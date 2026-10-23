SELECT students.first_name, students.last_name, SUM(courses.student_id) FROM students
JOIN courses ON courses.student_id = students.id
GROUP BY students.first_name, students.last_name
HAVING COUNT(student_id) > 1;