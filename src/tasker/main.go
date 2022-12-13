package main

import (
	"encoding/json"
	"fmt"
	"football/common/consts"
	"football/lib/http"
	"football/lib/leisu"
	"football/lib/reg"
	"time"
)

func main() {
	httpClient := http.NewClient(10 * time.Second)

	url := "https://www.leisu.com/yc/fixtures"

	rspCode, rsp, err := httpClient.Get(url, nil, nil)

	if err != nil {
		fmt.Printf("get fixtures failed. httpCode: %d | rsp: %s", rspCode, rsp)
		return
	}

	js, err := reg.FindAllString(string(rsp), consts.RegexpPatternMatchJs)

	if err != nil {
		fmt.Printf("FindAllString failed. err: %s", err)
		return
	}

	if len(js) == 0 {
		fmt.Printf("no js url. result: %s", js)
		return
	}
	header := make(map[string]string)
	//header["Content-Type"] = "text/html; charset=UTF-8"
	header["referer"] = "https://www.leisu.com/yc/fixtures"
	header["user-agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	header["accept-encoding"] = "gzip, deflate, br"

	rspCode, rsp, err = httpClient.Get(js[0], header, nil)

	if err != nil {
		fmt.Printf("get match js failed. httpCode: %d | rsp: %s", rspCode, rsp)
		return
	}

	jsContent := string(rsp)

	matchContent, err := reg.FindStringSubMatch(jsContent, consts.RegexpPatternMatchContent)

	if err != nil {
		fmt.Printf("get match content failed. matchContent: %s | err: %s", matchContent, err)
		return
	}

	matchJson := leisu.Rot(matchContent[1], 13)

	teamsContent, err := reg.FindStringSubMatch(jsContent, consts.RegexpPatternMatchTeams)

	if err != nil {
		fmt.Printf("get match content failed. teamsContent: %s | err: %s", teamsContent, err)
		return
	}

	teamJson := leisu.Rot(teamsContent[1], 13)

	match := &leisu.TopObject{}
	teams := make([]*leisu.Team, 0, 8)

	err = json.Unmarshal([]byte(matchJson), match)
	if err != nil {
		fmt.Printf("match Unmarshal failed. err: %s", err)
		return
	}

	err = json.Unmarshal([]byte(teamJson), &teams)
	if err != nil {
		fmt.Printf("team Unmarshal failed. err: %s", err)
		return
	}

	return
}
