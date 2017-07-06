.PHONY: all
all: format test build


SHELL  := env SKEL_ENV=$(SKEL_ENV) $(SHELL)
SKEL_ENV ?= dev

include config/$(SKEL_ENV).env
export $(shell sed 's/=.*//' config/$(SKEL_ENV).env)

.PHONY: build
build:
	go build -o ./bin/api api/main.go

format:
	find . -iname \*.go -exec go fmt {} \;

test:
	find . -iname \*_test.go -exec go test {} \;

install:
	go get github.com/codegangsta/negroni
	go get github.com/gorilla/mux
	go get github.com/joho/godotenv
	go get github.com/go-sql-driver/mysql
	go get github.com/asaskevich/govalidator
	go get github.com/gorilla/context
	go get github.com/rubenv/sql-migrate/...
	go get github.com/jmoiron/sqlx
database:
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) -e "create database if not exists $(DATABASE_NAME)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists company (id int primary key auto_increment, name varchar(100), email varchar(255), url varchar(150))"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists user (id int primary key auto_increment, name varchar(100), email varchar(255), password varchar(255), picture varchar(150))"
