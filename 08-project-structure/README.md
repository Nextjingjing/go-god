# บทที่ 8 Project structure

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## project structure คืออะไร?
project structure คือโครงสร้างของโปรเจคช่วยให้เรารู้ว่าควรจะเอาไฟล์ไหนไปอยู่ใน `Directory` ไหน? จริงๆ มันก็ไม่มีถูกไม่มีผิด ไฟล์ไหนจะเก็บในไหนก็ได้ ดังนั้นในบทที่ 8 จึงนำเสนอโดยอ้างอิงโครงสร้างจาก [golang-standards
project-layout](https://github.com/golang-standards/project-layout) ที่มีคนให้ดาวมากกว่า 5 หมื่นแล้ว

## Go Directories

### /cmd
```
├── cmd/
│   ├── server.go
```

`cmd/` เอาไว้เก็บไฟล์ Go ที่เอาไว้รัน Application เช่น `server.go` เอาไว้เปิด Server

### /internal
```
├── internal/
│   ├── service.go
│   ├── repository.go
│   ├── pkg/
│   │   ├── ...
```

`internal/` เอาไว้เก็บ Private application และ Library code ที่คุณไม่ต้องการให้ใครก็ตามนำไปใช้นอก Project นี้ได้

### /pkg
```
├── pkg/
│   ├── math/
│   │   ├── ...
```

`pkg/` เอาไว้เก็บ Library code ที่คุณต้องการให้คนอื่นนำไปใช้ได้หรือเป็น Public

## Directory อื่นๆ
- `/api` เก็บไฟล์นิยามต่างๆ เช่น OpenAPI/Swagger specs, JSON Schema หรือ Protocol Buffers (.proto)
- `/web` ส่วนของ Web Application เช่น Static assets, Templates หรือไฟล์จากโปรเจกต์ SPA (React/Vue)
- `/configs` ไฟล์ตั้งค่าเริ่มต้นหรือ Templates สำหรับการ Config ระบบ
- `/deployments` ไฟล์ตั้งค่าการ Deploy เช่น Kubernetes (Helm), Terraform หรือ Docker
- `/test` เก็บข้อมูลสำหรับการทดสอบ (Test data) หรือแอปพลิเคชันสำหรับทดสอบภายนอก
