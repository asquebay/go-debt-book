**Установка:**\
[user@nixos:~]$ git clone https://github.com/asquebay/go-debt-book.git

[user@nixos:~]$ cd go-debt-book

**Создание таблиц Users и Debts в PostgreSQL**\
(выберите только один из двух способов)

1) С помощью GUI для PostgreSQL (pgAdmin и др.):
```
CREATE USER "debt_book_owner" -- замените "debt_book_owner" на ваш вариант имени нового пользователя
LOGIN NOSUPERUSER NOCREATEROLE CREATEDB NOINHERIT REPLICATION NOBYPASSRLS
PASSWORD 'secure_password' -- замените 'secure_password' на ваш пароль (для нового пользователя, не ваш основной)
CONNECTION LIMIT -1;

SET ROLE debt_book_owner; -- замените debt_book_owner на того пользователя, которого вы указали ранее

CREATE DATABASE debt_book_db ENCODING 'UTF8'; -- замените debt_book_db на имя для новой базы данных
```

2) Через терминал:
```
sudo -u postgres psql <<-EOSQL
  CREATE USER debt_book_owner NOSUPERUSER NOCREATEROLE CREATEDB NOINHERIT REPLICATION NOBYPASSRLS CONNECTION LIMIT -1; -- замените debt_book_owner на ваш вариант имени нового пользователя
  ALTER USER debt_book_owner WITH PASSWORD 'secure_password'; -- замените 'secure_password' на ваш пароль (для нового пользователя, не ваш основной)
  CREATE DATABASE debt_book_db OWNER debt_book_owner ENCODING 'UTF8'; -- замените debt_book_db на имя для новой базы данных
EOSQL
```

**Заполнение таблиц (инициализация, нужно использовать только в первый раз)**\
go run main.go init

**Использование:**

Добавление пользователя
go run main.go add-user "Иван Иванов"

**Добавление долга (добавляет долг пользователю с данным User_ID)**

Добавление пользователю с User_ID = 2 долга на сумму 200.33:
go run main.go add-debt 2 200.33

Добавление пользователю с User_ID = 5 долга на сумму 1000:
go run main.go add-debt 5 1000

**Получение книги долгов (получает книгу долгов для пользователя с данным User_ID)**

Получение книги долгов для пользователя с User_ID = 2:
go run main.go get-book 2

Получение книги долгов для пользователя с User_ID = 5:
go run main.go get-book 5
