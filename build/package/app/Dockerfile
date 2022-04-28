FROM golang:1.18.1-alpine

ARG APP_SERVER_PORT
ARG APP_ROOT_PATH
ARG APP_EXEC_PATH

ENV CGO_ENABLED=0
ENV EXEC_FULL_PATH=${APP_ROOT_PATH}${APP_EXEC_PATH}

WORKDIR ${APP_ROOT_PATH}

RUN apk update
RUN set -eux && go install gotest.tools/gotestsum@v1.8.0

COPY go.mod go.sum ./
RUN go mod download

COPY . ${APP_ROOT_PATH}
RUN go build -v ${EXEC_FULL_PATH}/main.go

EXPOSE ${APP_SERVER_PORT}
CMD go run -v ${EXEC_FULL_PATH}/main.go
