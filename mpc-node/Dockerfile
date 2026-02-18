FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod tidy

EXPOSE 8001
CMD ["sh", "-c", "go run *.go"]