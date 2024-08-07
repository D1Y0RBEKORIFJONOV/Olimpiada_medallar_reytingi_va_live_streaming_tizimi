package interface17

import (
	"context"
	"medals/internal/models"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/medals"
)


type MedalsService interface {
	MedalRankings(ctx context.Context,req *models.MedalRankRequest) (*medals.MedalRankResponse, error)
	MedalCreate(ctx context.Context,req *models.MedalCreateRequest) (*models.GeneralResponseMedals, error)
	MedalUpdate(ctx context.Context,req *models.MedalUpdateRequest) (*models.GeneralResponseMedals, error)
	MedalDelete(ctx context.Context,req *models.MedalDeleteRequest) (*models.GeneralResponseMedals, error)
}
