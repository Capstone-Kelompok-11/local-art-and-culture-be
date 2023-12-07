FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN go build -o /app/main .
EXPOSE 80
CMD ["/app/main"]