SELECT course_name, COUNT(student_id) FROM courses
GROUP BY course_name
HAVING COUNT(student_id) > 1;