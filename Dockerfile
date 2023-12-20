# FROM golang:1.20-alpine AS builder

# WORKDIR /usr/local/src

# RUN apk --no-cache add bash git make gcc gettext musl-dev

# COPY ["server/go.mod", "./"]

# RUN go mod download

# COPY server ./

# RUN go build -o ./bin/server ./cmd/server.go

# FROM alpine AS runner

# COPY --from=builder /usr/local/src/bin/server /

FROM ubuntu:latest

WORKDIR /

COPY ["build/server", "./"]

# WORKDIR /app

# RUN go mod download
# RUN go mod tidy
# RUN go build cmd/main.go

# WORKDIR /

ENV PORT 3333

EXPOSE ${PORT}

CMD ["/server", "0.0.0.0"]