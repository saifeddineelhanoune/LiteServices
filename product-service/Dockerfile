FROM golang:1.23-bookworm


WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 5001

CMD [ "go", "run", "." ]
