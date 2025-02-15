FROM golang:1.23.4

WORKDIR /app

COPY consumer/go.mod ./
COPY consumer/go.sum ./
RUN go mod download

COPY consumer/*.go ./
RUN CGO_ENABLED=0 go build -o app .

CMD [ "/app/app" ]