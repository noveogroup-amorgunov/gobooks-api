# gobooks-api

Books REST API using [golang](https://golang.org).

## Installation

Create `.env` file and add database (postgres) connection string and port, which web server will listen:

```env
POSTGRES_URI=postgres://root:root@localhost:5432/books-store
PORT=:8080
```

You can create a free database on https://www.elephantsql.com.

Create table `books` and insert some rows:

```sql
create table books (id serial, title varchar, author varchar, year varchar);
insert into books (title, author, year) values('Golang is great', 'Mr. Go', '2012');
insert into books (title, author, year) values('How works Javascript', 'Douglas Crockford', '2018');
select * from books;
```

## Usage

Now you can run web server:

```bash
go run main.go
```

To test API endpoint I recommend https://www.getpostman.com.
