FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["server/go.mod", "./"]

RUN go mod download

COPY server ./

RUN go build -o ./bin/server ./cmd/server.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/server /

ENV PORT 3333

EXPOSE ${PORT}

CMD ["/server", "0.0.0.0"]