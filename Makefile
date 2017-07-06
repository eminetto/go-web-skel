.PHONY: all
all: format test build migrations frontend


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
	go get github.com/gorilla/sessions
	go get golang.org/x/oauth2
	go get github.com/go-sql-driver/mysql
	go get github.com/graymeta/stow
	go get github.com/graymeta/stow/google
	go get github.com/graymeta/stow/s3
	go get github.com/graymeta/stow/local
	go get github.com/asaskevich/govalidator
	go get github.com/getsentry/raven-go
	go get github.com/gorilla/context
	go get github.com/nats-io/go-nats
	go get github.com/sendgrid/sendgrid-go
	go get github.com/sendgrid/sendgrid-go/helpers/mail
	go get github.com/nats-io/gnatsd
	go get github.com/rubenv/sql-migrate/...
	go get gopkg.in/DATA-DOG/go-sqlmock.v1
	go get github.com/dgrijalva/jwt-go
	go get golang.org/x/crypto/bcrypt

db:
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) -e "create database if not exists $(DATABASE_NAME)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists company (id int primary key auto_increment, name varchar(100), email varchar(255), url varchar(150))"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists USER (id int primary key auto_increment, name varchar(100), email varchar(255), password varchar(255), picture varchar(150))"
