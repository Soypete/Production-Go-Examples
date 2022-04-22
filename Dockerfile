FROM golang:1.18-alpine

RUN apk update
RUN apk add git

WORKDIR /app
COPY go.* ./
RUN go mod download

EXPOSE 8080

COPY . ./
RUN go build -v -o main.go

CMD ["/app/main"]
