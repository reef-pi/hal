.PHONY: test
test:
	go test -coverprofile=coverage.txt ./...
.PHONY: lint
lint:
	./build/lint.sh
