FROM golang:1.19 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main.go

FROM scratch

WORKDIR /

COPY --from=build /app /app

EXPOSE 8080

CMD ["./"]