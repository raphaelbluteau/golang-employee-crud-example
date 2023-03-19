FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main api/main.go

EXPOSE 1323

CMD ["./main"]
