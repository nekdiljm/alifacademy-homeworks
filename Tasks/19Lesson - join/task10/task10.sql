SELECT students.*, COUNT(courses.student_id) FROM students
LEFT JOIN courses ON courses.student_id = students.id
GROUP BY students.id, students.first_name, students.last_name, students.age;