SELECT title, price, pages FROM books
WHERE  price > 100 AND pages > 300;

SELECT title, price, pages FROM books
WHERE price < 100 OR pages < 200;

SELECT title, price FROM books
WHERE author = 'Робер Кийосаки' AND price > 150;
