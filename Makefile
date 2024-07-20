unit-tests:
	go test ./...
#functional-tests:
#	go test ./functional_tests/transformer_test.go
build:
	docker-compose build
run:
	docker-compose up
