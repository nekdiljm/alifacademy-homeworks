SELECT students.*, courses.course_name FROM students
JOIN courses ON courses.student_id = students.id
WHERE courses.course_name = 'Go';