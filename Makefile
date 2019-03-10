BUILD_CONTEXT := ./build

.PHONY: go-test
go-test:
	@echo "Run all project tests..."
	go test -p 1 ./...

.PHONY: bin
bin: clean
	@echo "Build project binaries..."
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/tmm_linux_386
	GOOS=darwin GOARCH=386 go build -v -o $(BUILD_CONTEXT)/tmm_darwin_386
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/tmm_windows_386

.PHONY: clean
clean:
	rm -rf $(BUILD_CONTEXT)/

.PHONY: test
test: go-test

.PHONY: run
run:
	go run morris.go
