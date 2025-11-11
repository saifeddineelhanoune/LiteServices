FROM golang:1.23-bookworm


WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 5002

CMD [ "go", "run", "." ]
