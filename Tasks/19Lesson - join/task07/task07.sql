SELECT course_name, SUM(price) FROM courses
GROUP BY course_name;