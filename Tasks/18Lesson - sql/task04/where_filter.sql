SELECT title FROM books
WHERE author = 'Роберт Мартин';

SELECT title, price FROM books
WHERE  price > 100;

SELECT title, price FROM books
WHERE price < 50;

SELECT title, pages FROM books
WHERE pages > 500;

SELECT title, price FROM books
WHERE price >= 100 AND price <= 200;