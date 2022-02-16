FROM golang:1.17-alpine AS builder

WORKDIR /app/like-it

COPY go.mod .
COPY go.sum .
COPY ./infrastructure/internal/fonts/ARIAL.TTF .
RUN go mod download
COPY . .

# build
RUN go build -o main

# smaller image
FROM alpine:3.14
RUN apk add ca-certificates

COPY --from=builder /app/like-it/main /app/like-it
EXPOSE 8080
CMD ["/app/like-it"]