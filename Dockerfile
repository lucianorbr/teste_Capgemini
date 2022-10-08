FROM golang:1.19-alpine

RUN apk add --no-cache git

WORKDIR /app/api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./api .

EXPOSE 8080

CMD ["./api"]