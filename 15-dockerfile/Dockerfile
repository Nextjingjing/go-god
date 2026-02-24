# ---------- Stage 1: Build ----------
FROM golang:tip-alpine3.22 AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

# ---------- Stage 2: Final ----------
FROM scratch

WORKDIR /app
COPY --from=0 /go-app /bin/go-app

CMD [ "/bin/go-app" ]