.PHONY: prepare build

prepare:
	go mod tidy

build:
	go build -o dist/codesearch main.go

build-webapp:
	npm run build --prefix ./assets/
	rm -rf ./server/webapp/assets/
	cp -r ./assets/dist ./server/webapp/assets/
