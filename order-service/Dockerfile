FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install golang -y\
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . /app

EXPOSE 5002

CMD [ "go", "run", "." ]
