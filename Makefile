.PHONY: pre-commit

pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...
