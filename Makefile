.PHONY: fmt test vet check

fmt:
	gofmt -w ./cmd ./internal

test:
	go test ./...

vet:
	go vet ./...

check: fmt test vet
