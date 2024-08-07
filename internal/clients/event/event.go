package event

import (
	"log"

	event1 "github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func EventClient() event1.EventServiceClient {
	conn, err := grpc.NewClient("userservice:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := event1.NewEventServiceClient(conn)
	return client
}
