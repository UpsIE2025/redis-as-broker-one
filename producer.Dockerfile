FROM golang:1.23.4

WORKDIR /app

COPY producer/go.mod ./
COPY producer/go.sum ./
RUN go mod download

COPY producer/*.go ./
RUN CGO_ENABLED=0 go build -o app .

CMD [ "/app/app" ]

