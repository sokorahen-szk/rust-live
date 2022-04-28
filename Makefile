proto:
	rm -f api/proto/*.go
	protoc -I=./ --go_opt=paths=source_relative --go_out=./ \
	--go-grpc_opt=paths=source_relative --go-grpc_out=./ ./api/proto/*.proto

build:
	go mod tidy
	cd cmd/rust-live && go build -v

run: build
	cd cmd/rust-live && ./rust-live

clean:
	cd cmd/rust-live && go clean

test: proto
	go test -v ./...
	go mod tidy

docker-compose:
	docker-compose --env-file .env -f docker-compose.yml up --build