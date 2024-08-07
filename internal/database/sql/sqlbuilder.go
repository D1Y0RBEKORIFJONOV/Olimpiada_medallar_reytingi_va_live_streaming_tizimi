package sqlbuilder

import (
	"fmt"
	"log"
	"medals/internal/models"

	"github.com/Masterminds/squirrel"
)

func GetRankings() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("rank").
		OrderBy("score DESC").
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func CreateMedal(req *models.MedalCreateRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("medals").
		Columns("countryid", "type", "eventid", "athleteid").
		Values(req.CountryID, req.Type, req.EventID, req.AthleteID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println("builder", err)
		return "", nil, err
	}
	fmt.Println("Create Medal", query, args)
	return query, args, nil
}
func GetMedal(id int) (string, []interface{}, error) {
	query, args, err := squirrel.Select("type").
		From("medals").
		Where(squirrel.Eq{"medalid": id}).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func UpdateMedal(req *models.MedalUpdateRequest) (string, []interface{}, error) {
	updated := map[string]interface{}{
		"countryid": req.CountryID,
		"type":      req.Type,
		"eventid":   req.EventID,
		"athleteid": req.AthleteID,
	}
	query, args, err := squirrel.Update("medals").
		SetMap(updated).
		Where(squirrel.Eq{"medalid": req.MedalID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func DeleteMedal(req *models.MedalDeleteRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("medals").
		Where(squirrel.Eq{"medalid": req.MedalID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func ScoreCreate(req *models.Medals) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("rank").
		Columns("countryid", "gold", "silver", "bronze", "score").
		Values(req.CountryID, req.Gold, req.Silver, req.Bronze, req.Score).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func ScoreUpdate(req *models.Medals) (string, []interface{}, error) {
	queryBuilder := squirrel.Update("rank").
		Where(squirrel.Eq{"countryid": req.CountryID}).
		Set("gold", squirrel.Expr("gold+ ?", req.Gold)).
		Set("silver", squirrel.Expr("silver+ ?", req.Silver)).
		Set("bronze", squirrel.Expr("bronze+?", req.Bronze)).
		Set("score", squirrel.Expr("score+?", req.Score)).
		PlaceholderFormat(squirrel.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return "", nil, err
	}

	fmt.Println(query, args)
	return query, args, nil
}

func CheckCountryForScore(req string) (string, []interface{}, error) {
	query, args, err := squirrel.Select("EXISTS (SELECT 1").
		From("rank").
		Where(squirrel.Eq{"countryid": req}).
		Suffix(")").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	fmt.Println("CheckCountry", query, args)
	return query, args, nil
}

func CheckMedal(req string) (string, []interface{}, error) {
	query, args, err := squirrel.Select("type").
		From("medals").
		Where(squirrel.Eq{"medalid": req}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	fmt.Println("CheckCountry", query, args)
	return query, args, nil
}

func DeleteeMedal(req string) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("medals").
		Where(squirrel.Eq{"medalid": req}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return "", nil, err
	}
	fmt.Println("CheckCountry", query, args)
	return query, args, nil
}

func SpecialUpdateforrank(req *models.Medals) (string, []interface{}, error) {
	queryBuilder := squirrel.Update("rank").
		Where(squirrel.Eq{"countryid": req.CountryID}).
		Set("gold", squirrel.Expr("gold- ?", req.Gold)).
		Set("silver", squirrel.Expr("silver- ?", req.Silver)).
		Set("bronze", squirrel.Expr("bronze-?", req.Bronze)).
		Set("score", squirrel.Expr("score-?", req.Score)).
		PlaceholderFormat(squirrel.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return "", nil, err
	}

	fmt.Println(query, args)
	return query, args, nil
}
