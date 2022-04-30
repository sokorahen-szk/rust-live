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
	docker exec -t --env-file .env app gotestsum -- -p 1 -count=1 ./...
	go mod tidy

docker-compose:
	docker-compose --env-file .env -f ./build/docker-compose.yml up --build

docker-compose-destroy:
	docker-compose --env-file .env -f ./build/docker-compose.yml down --rmi all -v \
	&& docker volume prune --force