package athlete

import (
	"log"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/athlete"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet()athlete.AthleteServiceClient{
	conn, err := grpc.NewClient("10.10.1.208:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client:=athlete.NewAthleteServiceClient(conn)
	return client
}
