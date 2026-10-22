SELECT author, title FROM books
WHERE author = 'Роберт Мартин'
ORDER BY title;

SELECT title FROM books
ORDER BY title;

SELECT title, price FROM books
ORDER BY price DESC
LIMIT 3;

SELECT title, price FROM books
WHERE price > 1500;

SELECT title, pages FROM books
WHERE pages > 500;

SELECT title, price FROM books
WHERE price >= 500 AND price <= 1000;

UPDATE books SET
price = 750
WHERE id = 1;

SELECT price FROM books
WHERE id = 1;

DELETE FROM books
WHERE pages < 100;

SELECT * FROM books;


