package leisu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type MachTeam struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type TopObject struct {
	CompType    int      `json:"comp_type"`
	CurRound    int      `json:"cur_round"`
	CurStageId  int64    `json:"cur_stage_id"`
	CurSeasonId int64    `json:"cur_season_id"`
	Stages      []*Stage `json:"stages"`
}

type Stage struct {
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	GroupCount int      `json:"group_count"`
	RoundCount int      `json:"round_count"`
	Mode       int      `json:"mode"`
	Matches    []*Match `json:"matches"`
}

type Match struct {
	Id          int64        `json:"id"`
	HomeTeam    *MachTeam    `json:"home_team"`
	AwayTeam    *MachTeam    `json:"away_team"`
	Title       string       `json:"title"`
	Competition *Competition `json:"competition"`
	Venue       *Venue       `json:"venue"`
	MatchTime   int64        `json:"match_time"`
	StatusId    int          `json:"status_id"`
	HomeScore   int          `json:"home_score"`
	AwayScore   int          `json:"away_score"`
	RoundNum    int          `json:"round_num"`
	GroupNum    int          `json:"group_num"`
	Note        string       `json:"note"`
	HomeScores  []int        `json:"home_scores"`
	AwayScores  []int        `json:"away_scores"`
	OddList     []string     `json:"odd_list"`
	HalfOddList []string     `json:"half_odd_list"`
}

//func (m *Match) String() string {
//	return fmt.Sprintf("%+v", *m)
//}

func (m *Match) GetHashValue() (hashValue string) {
	if m == nil {
		return
	}

	matchByte, err := json.Marshal(m)

	if err != nil {
		return
	}

	mdByte := md5.Sum(matchByte)
	hashValue = hex.EncodeToString(mdByte[:])

	return
}

type Competition struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Venue struct {
	Id     int64  `json:"id"`
	NameZh string `json:"name_zh"`
	NameEn string `json:"name_en"`
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
