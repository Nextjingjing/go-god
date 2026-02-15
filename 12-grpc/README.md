# บทที่ 12 gRPC

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## gRPC คืออะไร?
เป็นอีกรูปแบบหนึ่งของ Inter-process communication (IPC) นิยมใช้ใน Backend/Microservices โดยทำงานบน HTTP/2 และใช้ Protocol Buffers (Protobuf) ซึ่งเป็นรูปแบบไบนารีทำให้ข้อมูลมีขนาดเล็ก ส่งข้อมูลได้รวดเร็ว และรองรับการทำ Streaming 

### มันดีกว่า REST API ยังไง?
- ไบนารีทำให้ข้อมูลมีขนาดเล็กกว่า JSON
- ส่งข้อมูลได้รวดเร็วกว่า
- Streaming ได้ (REST ทำไม่ได้)

### รูปแบบการสื่อสาร
- Unary RPC สื่อสารแบบ Request/Response เหมือน REST ปกติ
- Server streaming RPC
- Client streaming RPC
- Bidirectional streaming RPC ทั้งคู่ Stream

### เรื่องต้องรู้ gRPC
![grpc stub](https://grpc.io/img/landing-2.svg)

รูปจาก https://grpc.io/img/landing-2.svg

ย้อนไป Remote Procedure Calls (RPC ไม่ใช่ gRPC)คือกลไกที่ทำให้การเรียกฟังก์ชัน (procedure) ระหว่าง process ที่อยู่ คนละเครื่องในเครือข่าย ดูเหมือนกับการเรียกฟังก์ชันธรรมดาในโปรแกรมเดียวกัน
- ปัญหาที่คนยัง REST กันอยู่ก็คือต้องสร้าง Server ไม่พอต้อง Client ให้คุยกับ Server ได้ด้วย
  - หมายความว่าถ้าเราแก้ Server เราต้องตามไปแก้ Client ด้วย
- โชคยังดีที่ gRPC เข้ามาช่วยเราเรื่องนี้ครับ
  - gRPC ช่วยเจนโค้ดทั้งสองฝั่ง โดยฝั่ง Server จะเป็น Interface เปล่าๆ ที่เราต้อง Implement แต่ Client สามารถเรียกใช้ฟังชันก์ของ Server ได้เลย
  - ใช้ไฟล์ `.proto` ในการเจนโค้ดทั้งสองฝั่ง

ไฟล์ `Protocol Buffers (Protobuf)` สามารถเจนโค้ด `gRPC Stub/Server` และโค้ดที่เจนมาจะเป็นภาษาอะไรก็ได้ (ถ้าเขาทำตัวเจนมาให้นะ) สามารถดูภาษาที่เขาลองรับได้ที่ [เอกสารนี้](https://grpc.io/docs/languages/)

```proto
syntax = "proto3";

package calculator.v1;

option go_package = "./internal/pb/";

// CalculatorService defines the gRPC service for performing basic arithmetic operations.
service CalculatorService {
    // Add number1 and number2 and return the result.
    rpc Add(CalculationRequest) returns (CalculationResponse);
}

// CalculationRequest contains the two numbers to be calculated.
message CalculationRequest {
    double number1 = 1;
    double number2 = 2;
}

// CalculationResponse contains the result of the calculation.
message CalculationResponse {
    double result = 1;
}
```

- `service` เอาไว้นิยาม `Interface` ของทั้ง Client และ Server
- `message` เอาไว้กำหนด format การรับ-ส่งข้อมูล
- `CalculationRequest`, `CalculatorService`, ... เหล่านี้ต้องเป็น `Pascal case`
- `option go_package = "./internal/pb/";` เอาไว้เป็น path ที่ไว้เจนโค้ดไปใส่

การเจนโค้ดในที่นี้ผมจะเจนเป็น `Golang` นั้นแหละ

## การเตรียมเครื่องมือที่ใช้เจนโค้ดจาก Protocol Buffers (Protobuf)
- ให้ติดตั้ง https://protobuf.dev/installation/ อันนี้คือ Compiler
- ติดตั้ง gRPC tool ของ Golang
  - มันต้องลง `go.mod` ให้โปรเจคด้วย
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## การเจนโค้ด
```bash
protoc --go_out=. --go-grpc_out=. api/proto/v1/*.proto
```
- `--go-grpc_out=. api/proto/v1/*.proto` คือที่อยู่ของไฟล์ `.proto`
- `--go_out=.` โค้ดที่เจนออกมาจะเก็บไว้ไหน ซึ่งมันจะอ้างอิงจากไฟล์ `.proto` ด้วย
```proto
syntax = "proto3";

package calculator.v1;

option go_package = "./internal/pb/";
```

### ก่อนเจน
```
.
├── Makefile
├── README.md
├── api
│   └── proto
│       └── v1
│           └── calculator.proto
├── go.mod
├── go.sum
```
### หลังเจน
```
.
├── Makefile
├── README.md
├── api
│   └── proto
│       └── v1
│           └── calculator.proto
├── go.mod
├── go.sum
└── internal
    └── pb
        ├── calculator.pb.go
        └── calculator_grpc.pb.go
```
### ผมแนะนำให้ทำ `Makefile`
```makefile
.PHONY: gen-v1

gen-v1:
	protoc --go_out=. --go-grpc_out=. api/proto/v1/*.proto
```
การรัน
```
make gen-v1
```

## มันเจนอะไรมาให้เรา
```
.
├── internal
    └── pb
        ├── calculator.pb.go
        └── calculator_grpc.pb.go
        ...
```

ไฟล์ที่สำคัญคือ `calculator_grpc.pb.go`

```go
...
type CalculatorServiceServer interface {
	// Add number1 and number2 and return the result.
	Add(context.Context, *CalculationRequest) (*CalculationResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

type UnimplementedCalculatorServiceServer struct{}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method Add not implemented")
}
...
```

อันนี้คือ `interface` ของฝั่ง Server ที่เราต้อง `Implement`
```
.
├── internal
    ├── handler
    │   └── calculator_service_server_impl.go
    └── pb
        ├── calculator.pb.go
        └── calculator_grpc.pb.go
```

ผมจะสร้างไฟล์ `./internal/handler/calculator_service_server_impl.go`
```go
package handler

type calculatorServiceServerImpl struct {
	pb.UnimplementedCalculatorServiceServer
}

func NewCalculatorServiceServer() pb.CalculatorServiceServer {
	return &calculatorServiceServerImpl{}
}

func (s *calculatorServiceServerImpl) Add(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Number1 + req.Number2
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}
```
ข้อสังเกต
```go
type calculatorServiceServerImpl struct {
	pb.UnimplementedCalculatorServiceServer
}
```
จะเป็นเหมือนการฝัง `Embed` method ทุกอย่างไปใน `calculatorServiceServerImpl` และให้สังเกตว่า

ในไฟล์ `api/v1/calculator_grpc.pb.go`
```go
func (UnimplementedCalculatorServiceServer) Add(context.Context, *CalculationRequest) (*CalculationResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method Add not implemented")
}
```
- แต่ถ้าไม่ได้ `implement` ก็จะเกิด Error ส่งให้ Client ตอน Client พยายามเรียกใช้ `Add()`

พูดง่ายๆ ถ้าเรา Implement ไม่ครบก็จะ Error

## การสร้าง Server
ไฟล์ `./cmd/server/server.go`
```go
package main

import (
	"log"
	"net"

	"github.com/Nextjingjing/go-god/12-grpc/internal/handler"
	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, handler.NewCalculatorServiceServer())
	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

## การใช้ Postman 
เพื่อการทดสอบโดยยังไม่ต้องมี Client ก็ได้

![postman-grpc1.png](/docs/images/postman-grpc1.png)
![postman-grpc2.png](/docs/images/postman-grpc2.png)

- แล้วกด `Import`

![postman-grpc3.png](/docs/images/postman-grpc3.png)

![postman-grpc4.png](/docs/images/postman-grpc4.png)

สำเร็จแล้ว !

## การสร้าง Client
ไฟล์ `./cmd/client/client.go`
```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cred := insecure.NewCredentials()

	// create a gRPC client connection
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create a gRPC client
	client := pb.NewCalculatorServiceClient(conn)

	// call the Add method
	addResp, err := client.Add(ctx, &pb.CalculationRequest{Number1: 10, Number2: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(addResp.Result)
```
ผลลัพธ์
```
15
```

## รูปแบบการ Streaming
ไฟล์ `./api/v1/calculator.proto`
```proto
syntax = "proto3";

package calculator.v1;

option go_package = "./internal/pb/";

// CalculatorService defines the gRPC service for performing basic arithmetic operations.
service CalculatorService {
    // Add number1 and number2 and return the result.
    rpc Add(CalculationRequest) returns (CalculationResponse);
    // Average streams a sequence of numbers from the client and returns the average once the client has finished sending numbers.
    rpc Average(stream AverageRequest) returns (CalculationResponse);
    // Multiplication Table takes a single number and returns a stream of results for the multiplication table of that number from 1 to 12.
    rpc MultiplicationTable(MultiplicationTableRequest) returns (stream CalculationResponse);
}

...

message AverageRequest {
    double number = 1;
}

message MultiplicationTableRequest {
    double number = 1;
}
```

จากนั้นเจนโค้ดใหม่ `make gen-v1`

### Client streaming
ที่ไฟล์ `./internal/handler/calculator_service_server_impl.go`
```go

...

func (s *calculatorServiceServerImpl) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.CalculationResponse]) error {
	sum := 0.0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
	if count == 0 {
		return errors.New("Error: No numbers provided")
	}
	average := sum / float64(count)
	return stream.SendAndClose(&pb.CalculationResponse{
		Result: average,
	})
}
```

ที่ฝั่ง Client

```go
// call the Average method
stream, err := client.Average(ctx)
if err != nil {
	panic(err)
}
numbers := []float64{10, 20, 30, 40, 50}
for _, num := range numbers {
	if err := stream.Send(&pb.AverageRequest{Number: num}); err != nil {
		panic(err)
	}
}
// close the stream and receive the average response
resp, err := stream.CloseAndRecv()
if err != nil {
	panic(err)
}
fmt.Println("Average:", resp.Result)
```
ผลลัพธ์
```
Average: 30
```

### Server streaming
ที่ไฟล์ `./internal/handler/calculator_service_server_impl.go`
```go
func (s *calculatorServiceServerImpl) MultiplicationTable(req *pb.MultiplicationTableRequest, stream grpc.ServerStreamingServer[pb.CalculationResponse]) error {
	for i := 1; i <= 12; i++ {
		result := req.Number * float64(i)
		err := stream.Send(&pb.CalculationResponse{
			Result: result,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
```

ที่ฝั่ง Client

```go
// call Multiplication table method
mulTableStream, err := client.MultiplicationTabl(ctx, &pb.MultiplicationTableRequest{Number: 5})
if err != nil {
	panic(err)
}
for {
	resp, err := mulTableStream.Recv()
	if err == io.EOF {
		break
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("Multiplication Table Result:", resp.Result)
}
```
ผลลัพธ์
```
Multiplication Table Result: 5
Multiplication Table Result: 10
Multiplication Table Result: 15
Multiplication Table Result: 20
Multiplication Table Result: 25
Multiplication Table Result: 30
Multiplication Table Result: 35
Multiplication Table Result: 40
Multiplication Table Result: 45
Multiplication Table Result: 50
Multiplication Table Result: 55
Multiplication Table Result: 60
```