.PHONY: prepare build

prepare:
	go mod tidy

build:
	go build -o dist/home main.go

build-webapp:
	npm run build --prefix ./assets/
	rm -rf ./server/webapp/dist/
	cp -r ./assets/dist ./server/webapp/dist/
