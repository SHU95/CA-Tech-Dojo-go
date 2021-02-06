FROM golang:1.15.2-alpine
COPY ./go-app /app/go
WORKDIR /app/go
RUN go build -o main .
CMD ["/app/go/main"]
