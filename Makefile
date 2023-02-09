clear:
	rm bin/*

build: 
	go build -o bin/lambda lambda/main.go

build-local:
	go build -o bin/main cmd/main.go

