SELECT course_name, SUM(price) FROM courses
GROUP BY course_name
ORDER BY SUM(price) DESC
LIMIT 1;