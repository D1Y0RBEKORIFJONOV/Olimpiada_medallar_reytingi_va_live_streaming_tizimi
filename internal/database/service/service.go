package databaseservice

import (
	"context"
	interface17 "medals/internal/interface"
	"medals/internal/models"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/medals"
)

type DatabaseService struct {
	Psql interface17.MedalsService
}

func (u *DatabaseService) MedalRankings(ctx context.Context, req *models.MedalRankRequest) (*medals.MedalRankResponse, error) {
	return u.Psql.MedalRankings(ctx, req)
}

func (u *DatabaseService) MedalCreate(ctx context.Context, req *models.MedalCreateRequest) (*models.GeneralResponseMedals, error) {
	return u.Psql.MedalCreate(ctx, req)
}

func (u *DatabaseService) MedalUpdate(ctx context.Context, req *models.MedalUpdateRequest) (*models.GeneralResponseMedals, error) {
	return u.Psql.MedalUpdate(ctx, req)
}

func (u *DatabaseService) MedalDelete(ctx context.Context, req *models.MedalDeleteRequest) (*models.GeneralResponseMedals, error) {
	return u.Psql.MedalDelete(ctx, req)
}
