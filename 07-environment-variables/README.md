# บทที่ 7 Environment Variables

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- OS Environment Variables
- .env

## การขอค่า Environment variables ของ OS
```go
package main

import (
	"os"
)

func main() {
	// Access environment variables
	data := os.Getenv("PATH")
	println(data[:100])
}

```

ผลลัพธ์
```bash
C:\Program Files\Go\bin;c:\Users\snext\AppData\Roaming\Code\User\globalStorage\github.copilot-chat\d
```

## การอ่านค่าจากไฟล์ .env
```go
package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load(".env.example")

	// Access .env variable
	data2 := os.Getenv("SECRET_DATA")
	println(data2)
}
```
ที่ไฟล์ .env.example ผมมี
```bash
# This is a sample environment variable file.
SECRET_DATA = NEXTJINGJING
```
ผลลัพธ์จะได้เป็น
```bash
NEXTJINGJING
```