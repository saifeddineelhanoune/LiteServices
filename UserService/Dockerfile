FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    apt-get install golang \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . /app

EXPOSE 8080

RUN go build UserModels.go
