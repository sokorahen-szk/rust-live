FROM golang:1.18.1-alpine

ARG APP_SERVER_PORT
ARG APP_ROOT_PATH
ARG APP_EXEC_PATH

ENV PORT=${APP_SERVER_PORT}
ENV ROOT_PATH=${APP_ROOT_PATH}
ENV EXEC_PATH=${APP_EXEC_PATH}
ENV CGO_ENABLED 0
ENV EXEC_FULL_PATH=${ROOT_PATH}/${EXEC_PATH}

WORKDIR ${ROOT_PATH}

RUN echo ${PORT}

RUN apk update && apk add git

RUN set -eux && go install gotest.tools/gotestsum@v1.8.0

COPY go.mod go.sum ./
RUN go mod download

EXPOSE ${PORT}

WORKDIR ${EXEC_FULL_PATH}
RUN go build -v
CMD ["go", "run", "-v", "main.go"]