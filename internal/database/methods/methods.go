package dbmethods

import (
	"context"
	"database/sql"
	"log"
	sqlbuilder "medals/internal/database/sql"
	"medals/internal/models"
	"strings"

	"github.com/D1Y0RBEKORIFJONOV/Olimpiada_medallar_reytingi_va_live_streaming_tizimi_protos/gen/go/medals"
)

type Database struct {
	Db *sql.DB
}

func (u *Database) MedalRankings(ctx context.Context, req *models.MedalRankRequest) (*medals.MedalRankResponse, error) {
	query, args, err := sqlbuilder.GetRankings()
	if err != nil {
		return nil, err
	}
	rows, err := u.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	var res []*medals.Medals

	for rows.Next() {
		var all medals.Medals
		if err := rows.Scan(&all.Countryid, &all.Gold, &all.Silver, &all.Bronze, &all.Score); err != nil {
			return nil, err
		}
		res = append(res, &all)
	}

	return &medals.MedalRankResponse{Rankings: res}, nil
}
func (u *Database) MedalCreate(ctx context.Context, req *models.MedalCreateRequest) (*models.GeneralResponseMedals, error) {
	ok := u.CheckCountry(req)
	if !ok {
		query, args, err := sqlbuilder.CreateMedal(req)
		if err != nil {
			return nil, err
		}
		if err := u.InsertToRank(req); err != nil {
			return nil, err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			log.Println("i'm here", err)
			return nil, err
		}
		return &models.GeneralResponseMedals{Status: "done"}, nil
	}
	query, args, err := sqlbuilder.CreateMedal(req)
	if err != nil {
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println("i'm here", err)
		return nil, err
	}
	if err := u.UpdateRank(req); err != nil {
		return nil, err
	}
	return &models.GeneralResponseMedals{Status: "done"}, nil
}

func (u *Database) MedalUpdate(ctx context.Context, req *models.MedalUpdateRequest) (*models.GeneralResponseMedals, error) {
	if err := u.UpdateScore(req); err != nil {
		log.Println(err)
		return nil, err
	}
	query, args, err := sqlbuilder.UpdateMedal(req)
	if err != nil {
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	if err := u.UpdateScorees(req); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralResponseMedals{Status: "done"}, nil
}

func (u *Database) MedalDelete(ctx context.Context, req *models.MedalDeleteRequest) (*models.GeneralResponseMedals, error) {
	if err := u.DeleteScore(req); err != nil {
		log.Println(err)
		return nil, err
	}
	query, args, err := sqlbuilder.DeleteMedal(req)
	if err != nil {
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &models.GeneralResponseMedals{Status: "done"}, nil
}

func (u *Database) CheckCountry(req *models.MedalCreateRequest) bool {
	var check bool
	query, args, err := sqlbuilder.CheckCountryForScore(req.CountryID)
	if err != nil {
		log.Println(err)
		return false
	}
	if err := u.Db.QueryRow(query, args...).Scan(&check); err != nil {
		log.Println(err)
		return false
	}
	if !check {
		return false
	}
	return true
}
func (u *Database) InsertToRank(req *models.MedalCreateRequest) error {
	req.Type = strings.ToLower(req.Type)
	switch req.Type {
	case "gold":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Gold:      1,
			Score:     3,
		}
		query, args, err := sqlbuilder.ScoreCreate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "silver":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Silver:    1,
			Score:     2,
		}
		query, args, err := sqlbuilder.ScoreCreate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "bronze":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Bronze:    1,
			Score:     1,
		}
		query, args, err := sqlbuilder.ScoreCreate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (u *Database) UpdateRank(req *models.MedalCreateRequest) error {
	req.Type = strings.ToLower(req.Type)
	switch req.Type {
	case "gold":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Gold:      1,
			Score:     3,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "silver":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Silver:    1,
			Score:     2,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "bronze":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Bronze:    1,
			Score:     1,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
func (u *Database) UpdateScorees(req *models.MedalUpdateRequest) error {
	req.Type = strings.ToLower(req.Type)
	switch req.Type {
	case "gold":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Gold:      1,
			Score:     3,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "silver":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Silver:    1,
			Score:     2,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "bronze":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Bronze:    1,
			Score:     1,
		}
		query, args, err := sqlbuilder.ScoreUpdate(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
func (u *Database) BeforeUpdateRank(type1 string, req *models.MedalUpdateRequest) error {
	type1 = strings.ToLower(type1)
	switch type1 {
	case "gold":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Gold:      1,
			Score:     3,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "silver":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Silver:    1,
			Score:     2,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "bronze":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Bronze:    1,
			Score:     1,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (u *Database) CheckMedalExists(req *medals.MedalUpdateRequest) bool {
	query, args, err := sqlbuilder.CheckMedal(req.Medalid)
	if err != nil {
		log.Fatal(err)
		return false
	}
	var check string
	if err := u.Db.QueryRow(query, args...).Scan(&check); err != nil {
		log.Println(err)
		return false
	}
	if check == `` {
		return false
	}
	return true
}
func (u *Database) UpdateScore(req *models.MedalUpdateRequest) error {
	query, args, err := sqlbuilder.CheckMedal(req.MedalID)
	if err != nil {
		log.Println(err)
		return err
	}
	var type1 string
	if err := u.Db.QueryRow(query, args...).Scan(&type1); err != nil {
		log.Println(err)
		return err
	}
	if err := u.BeforeUpdateRank(type1, req); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (u *Database) DeleteScore(req *models.MedalDeleteRequest) error {
	query, args, err := sqlbuilder.DeleteeMedal(req.MedalID)
	if err != nil {
		log.Println(err)
		return err
	}
	var medailid string
	var newReq models.MedalCreateRequest
	if err := u.Db.QueryRow(query, args...).Scan(&medailid,&newReq.CountryID, &newReq.Type, &newReq.EventID, &newReq.AthleteID); err != nil {
		log.Println(err)
		return err
	}
	if err := u.deleterank(newReq.Type, &newReq); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (u *Database) deleterank(type1 string, req *models.MedalCreateRequest) error {
	type1 = strings.ToLower(type1)
	switch type1 {
	case "gold":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Gold:      1,
			Score:     3,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "silver":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Silver:    1,
			Score:     2,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	case "bronze":
		var newreq = models.Medals{
			CountryID: req.CountryID,
			Bronze:    1,
			Score:     1,
		}
		query, args, err := sqlbuilder.SpecialUpdateforrank(&newreq)
		if err != nil {
			return err
		}
		_, err = u.Db.Exec(query, args...)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
