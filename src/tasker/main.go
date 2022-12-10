package main

import (
	"fmt"
	"football/common/consts"
	"football/lib/http"
	"football/lib/reg"
	"football/lib/tools"
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
	header := map[string][]string{}
	header["referer"] = []string{"https://www.leisu.com/yc/fixtures"}
	header["user-agent"] = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"}
	header["accept-encoding"] = []string{"gzip, deflate, br"}

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

	matchJson := tools.Rot(matchContent[1], 13)

	teamsContent, err := reg.FindStringSubMatch(jsContent, consts.RegexpPatternMatchTeams)

	if err != nil {
		fmt.Printf("get match content failed. teamsContent: %s | err: %s", teamsContent, err)
		return
	}

	teamJson := tools.Rot(teamsContent[1], 13)

	fmt.Println(matchJson)
	fmt.Println(teamJson)

	return
}
