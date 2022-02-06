FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /ha-support-service

EXPOSE 3000

CMD ["/ha-support-service"]