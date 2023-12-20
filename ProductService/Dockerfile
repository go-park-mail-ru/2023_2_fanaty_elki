FROM ubuntu:latest

WORKDIR /

COPY ["build/product", "./"]

# WORKDIR /app

# RUN go mod download
# RUN go mod tidy
# RUN go build cmd/main.go

# WORKDIR /

EXPOSE 8082

CMD [ "./product" ]