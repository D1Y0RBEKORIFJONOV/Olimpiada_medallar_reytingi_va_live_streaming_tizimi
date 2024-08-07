package eventfunctions

import (
	"context"
	"log"

	event1 "github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/event"
)

type Event struct {
	E event1.EventServiceClient
}

func (u *Event) CheckEvent(id string) (string, error) {
	res, err := u.E.GetEvent(context.Background(), &event1.GetEventReq{EventId: id})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.WinnerId, nil
}
