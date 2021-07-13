.PHONY : build run

proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
	protoc --proto_path=api/proto --go_out=plugins=grpc:api/proto \
	--go_opt=paths=source_relative $(proto_file);)

migration:
	@go run cmd/migration/migration.go

build:
	go build -o bin

run: build
	./bin

docker:
	docker build -t be-shark:latest .

run-container:
	docker run --name=be-shark --network="host" -d be-shark

# unit test & calculate code coverage
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt

clear:
	rm bin be-shark
