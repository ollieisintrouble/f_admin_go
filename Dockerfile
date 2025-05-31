FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o f_admin ./cmd/f_admin

EXPOSE 8080

CMD ["./f_admin"]
