# FROM ubuntu:latest

# WORKDIR /

# COPY ["build/auth", "./"]

# # WORKDIR /app

# # RUN go mod download
# # RUN go mod tidy
# # RUN go build cmd/main.go

# WORKDIR /

# EXPOSE 8081

# CMD [ "./auth" ]

#######################

FROM golang:1.19

WORKDIR /

COPY . ./app

WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go

WORKDIR /

EXPOSE 8081

CMD [ "./app/main" ]