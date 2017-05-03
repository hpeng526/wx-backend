export PATH := $(GOPATH)/bin:$(PATH)
export GO15VENDOREXPERIMENT := 1

all: build

build: app

app:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/wx-backend_darwin_amd64 ./
	env GOOS=linux GOARCH=386 go build -o ./build/wx-backend_linux_386 ./
	env GOOS=linux GOARCH=amd64 go build -o ./build/wx-backend_linux_amd64 ./
	env GOOS=linux GOARCH=arm go build -o ./build/wx-backend_linux_arm ./
	env GOOS=windows GOARCH=386 go build -o ./build/wx-backend_windows_386.exe ./
	env GOOS=windows GOARCH=amd64 go build -o ./build/wx-backend_windows_amd64.exe ./
	env GOOS=linux GOARCH=mips64 go build -o ./build/wx-backend_linux_mips64 ./
	env GOOS=linux GOARCH=mips64le go build -o ./build/wx-backend_linux_mips64le ./

clean:
	rm -rf ./build/
