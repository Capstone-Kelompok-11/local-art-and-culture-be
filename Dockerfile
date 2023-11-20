FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . . 

RUN go build -o /app/main .
RUN go run main.go
EXPOSE 8080
CMD ["/app/main"]