test:
	@go test -run=./app_test.go -v

build-app:
	@go build -o build

exec-app:
	./build/testCi