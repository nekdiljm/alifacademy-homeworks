UPDATE books SET
price = 100
WHERE price < 100;

SELECT title, author, price FROM books
WHERE price <= 100;