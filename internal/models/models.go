package models

type MedalRankRequest struct{}

type Medals struct {
	CountryID string `json:"countryid"`
	Medalid   string `json:"-"`
	Gold      int64  `json:"gold"`
	Silver    int64  `json:"silver"`
	Bronze    int64  `json:"bronze"`
	Score     int64  `json:"score"`
}

type MedalRankResponse struct {
	Rankings []Medals `json:"rankings"`
}

type MedalCreateRequest struct {
	CountryID string `json:"countryid"`
	Type      string `json:"type"`
	EventID   string `json:"eventid"`
	AthleteID string `json:"athleteid"`
}

type MedalUpdateRequest struct {
	MedalID   string `json:"medalid"`
	CountryID string `json:"countryid"`
	Type      string `json:"type"`
	EventID   string `json:"eventid"`
	AthleteID string `json:"athleteid"`
}

type MedalDeleteRequest struct {
	MedalID string `json:"medalid"`
}

type GeneralResponseMedals struct {
	Status string `json:"status"`
}
