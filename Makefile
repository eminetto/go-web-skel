.PHONY: all
all: format test build migrations frontend


SHELL  := env CODENATION_ENV=$(CODENATION_ENV) $(SHELL)
CODENATION_ENV ?= dev

include config/$(CODENATION_ENV).env
export $(shell sed 's/=.*//' config/$(CODENATION_ENV).env)

.PHONY: build
build:
	go build -o ./bin/codenation web/codenation/main.go
	go build -o ./bin/consumer cmd/consumer/main.go
	go build -o ./bin/api api/main.go

frontend:
	cd web/codenation && gulp production

format:
	find . -iname \*.go -exec go fmt {} \;

test:
	find . -iname \*_test.go -exec go test {} \;

migrations:
	sql-migrate up -config=config/dbconfig.yml -env=production 

clean:
	rm bin/codenation
	rm -rf data/resume/*

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

install-frontend:
	cd web/codenation
	yarn install

db:
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) -e "create database if not exists $(DATABASE_NAME)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists batch (id int primary key auto_increment, name varchar(100) not null, date_start datetime, date_end datetime, city varchar(100), url varchar(100), description text)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists applicant (id int primary key auto_increment, name varchar(100) not null, nickname varchar(100) not null, picture varchar(100) null,email varchar(100) not null, phone varchar(20) not null, birthday datetime not null, current_job varchar(100) null, course varchar(100) not null, semester varchar(100) not null, dev_years varchar(100) not null, resume_url varchar(200) , cover_letter text, orientation varchar(100) not null, aptitude varchar(100) not null)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists company (id int primary key auto_increment, name varchar(100), logo varchar(255), email varchar(255), password varchar(255), city varchar(100), url varchar(150), description text)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists applicant_batch (applicant_id int not null, batch_id int not null, salary float, company_id int, selected_by int)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "create table if not exists company_batch (company_id int not null, batch_id int not null, positions int not null)"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "insert into batch values (1, 'Florianópolis', now(), '2017-12-31 00:00:00', 'Florianópolis', 'floripa.codenation.com.br', 'Descrição do batch de Florianópolis')"
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) $(DATABASE_NAME) -e "insert into batch values (2, 'Joinville', now(), '2017-12-31 00:00:00', 'Joinville', 'joinville.codenation.com.br', 'Descrição do batch de Joinville')"
