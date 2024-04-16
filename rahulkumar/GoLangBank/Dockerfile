# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


EXPOSE 8080
CMD [ "/app/main" ]