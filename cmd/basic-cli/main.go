package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/starlink-community/starlink-grpc-go/pkg/spacex.com/api/device"
	"google.golang.org/grpc"
)

var (
	addr = "192.168.100.1:9200"
	req  = "status"
)

func init() {
	flag.StringVar(&addr, "addr", addr, "grpc addr (dishy is at 192.168.100.1:9200, wifi is at 192.168.1.1:9000")
	flag.StringVar(&req, "req", req, "status or ping")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
	}
	c := pb.NewDeviceClient(conn)

	in := new(pb.Request)

	// Set the type of request you want to make
	// The list of Request_* types can be found in pkg/spacex.com/api/device/device.go
	// https://github.com/starlink-community/starlink-grpc-go/blob/fb342a752e3ba99ed88acb9d28a09dbe4c12d292/pkg/spacex.com/api/device/device.pb.go#L178
	switch req {
	case "status":
		in.Request = &pb.Request_GetStatus{}
	case "history":
		in.Request = &pb.Request_GetHistory{}
	case "ping":
		in.Request = &pb.Request_GetPing{}
	default:
		fmt.Println("unknown request", req)
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// r is the gprc response
	r, err := c.Handle(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	t := proto.TextMarshaler{}
	t.Marshal(os.Stdout, r)
}
