package entity

import "fmt"

type Match struct {
	Id                  int64  `gorm:"primaryKey" json:"id"`
	MatchThirdId        int64  `gorm:"column:match_third_id" json:"match_third_id"`
	LeagueThirdId       int64  `json:"league_third_id"`
	StageThirdId        int64  `json:"stage_third_id"`
	SeasonThirdId       int64  `json:"season_third_id"`
	Title               string `json:"title"`
	RoundNum            int    `json:"round_num"`
	GroupNum            int    `json:"group_num"`
	HomeTeamThirdId     int64  `json:"home_team_third_id"`
	HomeTeamName        string `json:"home_team_name"`
	HomeTeamLogo        string `json:"home_team_logo"`
	AwayTeamThirdId     int64  `json:"away_team_third_id"`
	AwayTeamName        string `json:"away_team_name"`
	AwayTeamLogo        string `json:"away_team_logo"`
	MatchTime           int64  `json:"match_time"`
	StatusId            int    `json:"status_id"`
	HomeScore           int    `json:"home_score"`
	HomeFirstHalfScore  int    `json:"home_first_half_score"`
	HomeSecondHalfScore int    `json:"home_second_half_score"`
	HomeRedCard         int    `json:"home_red_card"`
	HomeYellowCard      int    `json:"home_yellow_card"`
	HomeCorner          int    `json:"home_corner"`
	AwayScore           int    `json:"away_score"`
	AwayFirstHalfScore  int    `json:"away_first_half_score"`
	AwaySecondHalfScore int    `json:"away_second_half_score"`
	AwayRedCard         int    `json:"away_red_card"`
	AwayYellowCard      int    `json:"away_yellow_card"`
	AwayCorner          int    `json:"away_corner"`
	OddList             int64  `json:"odd_list"`
	HalfOddList         int64  `json:"half_odd_list"`
	HashValue           string `json:"hash_value"`
	CreateTime          int64  `json:"create_time"`
	UpdateTime          int64  `json:"update_time"`
}

func (e *Match) TableName() string {
	return "data_match"
}

func (e *Match) PrimarySeted() bool {
	return e.Id > 0
}

func (e *Match) String() string {
	return fmt.Sprintf("%+v", *e)
}
