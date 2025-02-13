FROM --platform=linux/amd64 golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/


COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
