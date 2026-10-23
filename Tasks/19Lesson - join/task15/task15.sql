SELECT students.first_name,
       students.last_name,
       COUNT(courses.student_id),
       SUM(courses.price * courses.duration),
       AVG(price)
FROM students
JOIN courses ON courses.student_id = students.id
GROUP BY students.first_name, students.last_name
ORDER BY SUM(courses.price * courses.duration) DESC;
