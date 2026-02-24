# บทที่ 14 Unit Testing

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- Docker

```dockerfile
FROM golang:tip-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

EXPOSE 8080

# Run
CMD ["/go-app"]
```

- `FROM golang:tip-alpine3.22` เป็น Linux ดังนั้นเราจึงต้อง Compile เป็น Linux คือ `RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app`