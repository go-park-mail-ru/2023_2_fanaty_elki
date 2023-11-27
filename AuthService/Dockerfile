FROM golang:1.20-alpine AS builder

WORKDIR /

COPY . ./app

WORKDIR /app

COPY ["./go.mod", "/"]

RUN apk --no-cache add bash git make gcc gettext musl-dev

RUN go mod tidy

RUN go build cmd/main.go

WORKDIR /

EXPOSE 8081

CMD [ "./app/main" ]