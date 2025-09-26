.PHONY: prepare build

prepare:
	go mod tidy

build: build-webapp
	go build -o dist/hsh main.go

build-webapp:
	npm run build --prefix ./assets/
	rm -rf ./server/webapp/dist/
	cp -r ./assets/dist ./server/webapp/dist/
