package services

import (
	"context"
	"log"
	athletefunctions "medals/internal/clients/athlete/functions"
	eventfunctions "medals/internal/clients/event/functions"
	databaseservice "medals/internal/database/service"
	"medals/internal/models"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/medals"
)

type Server struct {
	medals.UnimplementedMedalsServiceServer
	S *databaseservice.DatabaseService
	A *athletefunctions.Athlete
	E *eventfunctions.Event
}

func (u *Server) GetCountryRanking(ctx context.Context, req *medals.MedalRankRequest) (*medals.MedalRankResponse, error) {
	var newReq models.MedalRankRequest
	res, err := u.S.MedalRankings(ctx, &newReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &medals.MedalRankResponse{Rankings: res.Rankings}, nil
}
func (u *Server) MedalCreate(ctx context.Context, req *medals.MedalCreateRequest) (*medals.GeneralResponseMedals, error) {
	Countryid, err := u.A.CheckAthlete(req.Atheleteid)
	if err != nil {
		log.Println(err)
		return &medals.GeneralResponseMedals{Status: "There is no any athlete like this"}, err
	}
	if Countryid != req.Countryid {
		return &medals.GeneralResponseMedals{Status: "this athlete is not from this country"}, nil
	}
	// WinnerId, err := u.E.CheckEvent(req.Eventid)
	// if err != nil {
	// 	log.Println(err)
	// 	return &medals.GeneralResponseMedals{Status: "there wern't any event"}, nil
	// }
	// if WinnerId != req.Atheleteid {
	// 	return &medals.GeneralResponseMedals{Status: "this athlete didn't win that match"}, nil
	// }
	var newReq = models.MedalCreateRequest{
		CountryID: Countryid,
		Type:      req.Type,
		EventID:   req.Eventid,
		AthleteID: req.Atheleteid,
	}
	res, err := u.S.MedalCreate(ctx, &newReq)
	if err != nil {
		log.Println("here", err)
		return nil, err
	}
	return &medals.GeneralResponseMedals{Status: res.Status}, nil
}

func (u *Server) MedalUpdate(ctx context.Context, req *medals.MedalUpdateRequest) (*medals.GeneralResponseMedals, error) {
	var newReq = models.MedalUpdateRequest{
		MedalID:   req.Medalid,
		CountryID: req.Countryid,
		Type:      req.Type,
		EventID:   req.Eventid,
		AthleteID: req.Atheleteid,
	}
	res, err := u.S.MedalUpdate(ctx, &newReq)
	if err != nil {
		return nil, err
	}
	return &medals.GeneralResponseMedals{Status: res.Status}, nil
}

func (u *Server) MedalDelete(ctx context.Context, req *medals.MedalDeleteRequest) (*medals.GeneralResponseMedals, error) {
	res, err := u.S.MedalDelete(ctx, &models.MedalDeleteRequest{MedalID: req.Medalid})
	if err != nil {
		return nil, err
	}
	return &medals.GeneralResponseMedals{Status: res.Status}, nil
}
