FROM golang:tip-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

EXPOSE 8080

# Run
CMD ["/go-app"]