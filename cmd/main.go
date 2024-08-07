package main

import (
	"fmt"
	"log"
	"medals/internal/connections"
	"net"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/medals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}
	server, err := connections.NewServer()
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	medals.RegisterMedalsServiceServer(s, server)
	reflection.Register(s)
	fmt.Println("server started on port 8080")
	log.Fatal(s.Serve(lis))
}
