FROM golang:1.18.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o main.go

EXPOSE 8090

CMD ["./main"]