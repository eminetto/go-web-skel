.PHONY: all
all: format test build


SHELL  := env SKEL_ENV=$(SKEL_ENV) $(SHELL)
SKEL_ENV ?= dev

include config/$(SKEL_ENV).env
export $(shell sed 's/=.*//' config/$(SKEL_ENV).env)

.PHONY: build migrations
doc-build:
	cd doc; aglio -i api.apib --theme-full-width --no-theme-condense -o index.html

doc-install:
	npm install -g aglio drakov dredd

doc-serve:
	cd doc; aglio -i api.apib --theme-full-width --no-theme-condense -s

doc-mock:
	cd doc; drakov -f api.apib -p 4000

doc-test:
	cd doc; dredd api.apib http://localhost:4000

build:
	dep ensure
	go build -o ./bin/cli cmd/main.go
	go build -o ./bin/api api/main.go

db:
	$(DATABASE_DRIVER) -u$(DATABASE_USER) -p$(DATABASE_PASSWORD) -h $(DATABASE_HOST) -e "create database if not exists $(DATABASE_NAME)"

new-migration: ## create a new migration, use make new-migration m=message to set the message
	sql-migrate new -config=./config/dbconfig.yml -env=production "$(m)"

migrations:
	sql-migrate up -config=config/dbconfig.yml -env=production 

test:
	./cli/go.test.sh