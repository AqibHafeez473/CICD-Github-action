# Step 1: Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod init cicdapp
RUN go build -o app .

# Step 2: Runtime image
FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]