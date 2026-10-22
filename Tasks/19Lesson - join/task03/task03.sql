SELECT students.first_name,
       students.last_name,
       courses.course_name
FROM students
JOIN courses ON courses.student_id = students.id;
