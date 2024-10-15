FROM golang:1.23.2-alpine

WORKDIR /app

COPY src/go.mod src/go.sum /app/
RUN go mod download

COPY src/ /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o main.out

EXPOSE 8080

CMD ["/app/main.out"]