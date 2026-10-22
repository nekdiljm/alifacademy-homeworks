UPDATE books SET
price = 80
WHERE id = 1;

SELECT title, price FROM books
WHERE id = 1;

UPDATE books SET
author = 'George Orel'
WHERE id = 2;

SELECT title, author FROM books
WHERE id = 2;

UPDATE books SET
title = 'Crime and Punishment',
price = 300,
pages = 630
WHERE id = 3;

SELECT title, author, price, pages FROM books
WHERE id = 3;
