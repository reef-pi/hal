.PHONY: test
test:
	go test ./...
.PHONY: lint
lint:
	./build/lint.sh
