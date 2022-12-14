package schedule

import (
	"fmt"
	"football/common/entity"
	"football/common/model"
	"football/lib/leisu"
	"time"
)

var Match = matchLogic{}

type matchLogic struct{}

func (s *matchLogic) InitSchedule(object *leisu.TopObject) (err error) {
	if object == nil {
		return
	}

	if len(object.Stages) == 0 {
		return
	}

	if len(object.Stages[0].Matches) == 0 {
		return
	}

	for _, match := range object.Stages[0].Matches {
		_ = s.InitMatch(object, match)
	}

	return
}

func (s *matchLogic) InitMatch(object *leisu.TopObject, match *leisu.Match) (err error) {
	entityMatch, err := model.Match.GetMatchByThirdId(match.Id)

	if err != nil {
		fmt.Printf("GetMatchByThirdId failed. id: %d | err: %s\n", match.Id, err)
		return
	}

	if entityMatch == nil {
		// 不存在，初始化比赛
		err = s.AddMatch(object, match)
		return
	} else if entityMatch.HashValue != match.GetHashValue() {
		// 已存在且有修改，更新比赛信息
		err = s.UpdateMatch(object, match, entityMatch)
		return
	}

	fmt.Printf("no need update. matchId: %d | home: %s | away: %s\n", match.Id, match.HomeTeam.Name, match.AwayTeam.Name)

	return
}

func (s *matchLogic) AddMatch(object *leisu.TopObject, match *leisu.Match) (err error) {
	entityMatch := &entity.Match{
		MatchThirdId:        match.Id,
		LeagueThirdId:       82,
		StageThirdId:        object.CurStageId,
		SeasonThirdId:       object.CurSeasonId,
		Title:               match.Title,
		RoundNum:            match.RoundNum,
		GroupNum:            match.GroupNum,
		HomeTeamThirdId:     match.HomeTeam.Id,
		HomeTeamName:        match.HomeTeam.Name,
		HomeTeamLogo:        match.HomeTeam.Logo,
		AwayTeamThirdId:     match.AwayTeam.Id,
		AwayTeamName:        match.AwayTeam.Name,
		AwayTeamLogo:        match.AwayTeam.Logo,
		MatchTime:           match.MatchTime,
		StatusId:            match.StatusId,
		HomeScore:           match.HomeScore,
		HomeFirstHalfScore:  match.HomeScores[1],
		HomeSecondHalfScore: match.HomeScores[0] - match.HomeScores[1],
		HomeRedCard:         match.HomeScores[2],
		HomeYellowCard:      match.HomeScores[3],
		HomeCorner:          match.HomeScores[4],
		AwayScore:           match.AwayScore,
		AwayFirstHalfScore:  match.AwayScores[1],
		AwaySecondHalfScore: match.AwayScores[0] - match.AwayScores[1],
		AwayRedCard:         match.AwayScores[2],
		AwayYellowCard:      match.AwayScores[3],
		AwayCorner:          match.AwayScores[4],
		HashValue:           match.GetHashValue(),
		CreateTime:          time.Now().Unix(),
		UpdateTime:          time.Now().Unix(),
	}

	err = model.Match.Add(entityMatch)

	if err != nil {
		fmt.Printf("match add failed. match: %s | err: %s\n", match, err)
	} else {
		fmt.Printf("match add success. matchId: %d\n", match.Id)
	}

	return
}

func (s *matchLogic) UpdateMatch(object *leisu.TopObject, match *leisu.Match, entityMatch *entity.Match) (err error) {
	prop := make(map[string]any)
	prop["match_time"] = match.MatchTime
	prop["status_id"] = match.StatusId
	prop["home_score"] = match.HomeScore
	prop["home_first_half_score"] = match.HomeScores[1]
	prop["home_second_half_score"] = match.HomeScores[0] - match.HomeScores[1]
	prop["home_red_card"] = match.HomeScores[2]
	prop["home_yellow_card"] = match.HomeScores[3]
	prop["home_corner"] = match.HomeScores[4]
	prop["away_score"] = match.AwayScore
	prop["away_first_half_score"] = match.AwayScores[1]
	prop["away_second_half_score"] = match.AwayScores[0] - match.AwayScores[1]
	prop["away_red_card"] = match.AwayScores[3]
	prop["away_yellow_card"] = match.AwayScores[4]
	prop["away_corner"] = match.AwayScores[5]
	prop["update_time"] = time.Now().Unix()
	prop["hash_value"] = match.GetHashValue()

	err = model.Match.Update(entityMatch, prop)

	if err != nil {
		fmt.Printf("match update failed. match: %s | err: %s\n", match, err)
	} else {
		fmt.Printf("match update success. matchId: %d\n", match.Id)
	}

	return
}
