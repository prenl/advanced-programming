## Introduction
Hi, this is a project called "Snippetbox" initially made by Alex Edwards in his book called "Let's Go".
This is the full code that is inside the book. Also, I made additional tasks such as account view, password change, snippet create test and etc.

## Used dependencies
```github.com/alexedwards/scs/mysqlstore v0.0.0-20231113091146-cef4b05350c8
github.com/alexedwards/scs/v2 v2.5.0
github.com/go-playground/form/v4 v4.2.0
github.com/go-sql-driver/mysql v1.7.1
github.com/julienschmidt/httprouter v1.3.0
github.com/justinas/alice v1.2.0
github.com/justinas/nosurf v1.1.1
golang.org/x/crypto v0.17.0
```
## How to install

First of all you should download go 1.21 from the official website of Go Language\
Also you have to install mysql community server\

Then you can use this command to enter the mysql `mysql -u root -p`
Use `source path/to/database.sql` to create necessary tables

Execute `go mod download` to install all the dependencies

- To execute tests use `go test ./cmd/web/` or `go test ./...` (flags available)
- To run the server use `go run ./cmd/web`

## Enjoy!

