unit-tests:
	go test ./...
#functional-tests:
#	go test ./functional_tests/transformer_test.go
build:
	#docker build . -t go-jenkins-simple:go-micro
	docker-compose build
run:
	docker-compose up
