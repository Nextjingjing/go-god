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
- แต่แบบนี้ยังมีข้อเสียเพราะ Go จริงๆ Compile เป็น Binary อยู่แล้ว ไม่จำเป็นต้องนำ Source code หรือพวก Go compiler/Runtime ก็ได้ ไปอยู่ใน Image ก็ได้

```bash
$ docker images go-app
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
go-app       latest    1b7fb137b16a   3 hours ago   908MB
```
- สังเกตว่าขนาดตั้ง Image 908MB!

## Best Practice (Multi-stage Builds)
```dockerfile
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
```

```bash
$ docker images go-app
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
go-app       latest    b7166f7f0aa6   35 seconds ago   19.3MB
```

- สังเกตว่าขนาดแค่ Image 19.3MB! ต่างกันเกือบๆ 50 เท่า