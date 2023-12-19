FROM golang:1.19

WORKDIR /

COPY build/user ./

# WORKDIR /app

# RUN go mod download
# RUN go mod tidy
# RUN go build cmd/main.go

# WORKDIR /

EXPOSE 8083

CMD [ "./app/user" ]