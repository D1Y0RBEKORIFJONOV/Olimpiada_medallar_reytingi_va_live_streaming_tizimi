package athletefunctions

import (
	"context"
	"log"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/athlete"
)

type Athlete struct{
	C athlete.AthleteServiceClient
}

func (u *Athlete) CheckAthlete(athleteid string)(string,error){
	res,err:=u.C.GetbyIdAthlete(context.Background(),&athlete.AthleteResponse{Id: athleteid})
	if err!=nil{
		log.Println(err)
		return "",err
	}
	return res.CountryId,nil
}