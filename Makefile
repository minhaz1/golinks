all: build

clean: ./bin
	rm -rf ./bin 

build: app.go
	mkdir -p bin
	GOBIN=./bin go install 

install: app.go
	mkdir -p /export/service/bin
	GOBIN=/export/service/bin go install
