.PHONY: all
all: build


.PHONY: build
build: 
	aglio -i api.apib --theme-full-width --no-theme-condense -o index.html
install:
	npm install -g aglio drakov

serve:
	aglio -i api.apib --theme-full-width --no-theme-condense -s

mock: 
	drakov -f api.apib -p 4000
