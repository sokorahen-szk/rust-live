proto:
	rm -f api/proto/*.go
	protoc -I=./ --go_opt=paths=source_relative --go_out=./ \
	--go-grpc_opt=paths=source_relative --go-grpc_out=./ ./api/proto/*.proto

build:
	go mod tidy
	cd cmd/rust-live && go build -v

clean:
	cd cmd/rust-live && go clean