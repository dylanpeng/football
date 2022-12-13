package leisu

import (
	"encoding/json"
	"fmt"
	"football/lib/http"
	"football/lib/reg"
	"time"
)

const (
	RegexpPatternMatchJs          = `https:\/\/static.leisu.com\/public\/askaliy\/zuqiu\/comp\/match-.*?js`
	RegexpPatternMatchContent     = `window\[_t19798\[2\]\] ='(.*?)'`
	RegexpPatternMatchTeams       = `window\[_t19798\[3\]\] ='(.*?)'`
	RegexpPatternMatchContentTeam = `".*?"`
)

func QueryMatch() (match *TopObject, err error) {
	httpClient := http.NewClient(10 * time.Second)

	url := "https://www.leisu.com/yc/fixtures"

	rspCode, rsp, err := httpClient.Get(url, nil, nil)

	if err != nil {
		fmt.Printf("get fixtures failed. httpCode: %d | rsp: %s", rspCode, rsp)
		return
	}

	js, err := reg.FindAllString(string(rsp), RegexpPatternMatchJs)

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

	matchContent, err := reg.FindStringSubMatch(jsContent, RegexpPatternMatchContent)

	if err != nil {
		fmt.Printf("get match content failed. matchContent: %s | err: %s", matchContent, err)
		return
	}

	if len(matchContent) < 2 {
		fmt.Printf("get match js content failed. jsContent: %s | err: %s", jsContent, err)
		return
	}

	match = &TopObject{}
	matchJson := Rot(matchContent[1], 13)

	err = json.Unmarshal([]byte(matchJson), match)
	if err != nil {
		fmt.Printf("match Unmarshal failed. err: %s", err)
		return
	}

	//teamsContent, err := reg.FindStringSubMatch(jsContent, RegexpPatternMatchTeams)
	//
	//if err != nil {
	//	fmt.Printf("get match content failed. teamsContent: %s | err: %s", teamsContent, err)
	//	return
	//}
	//
	//teamJson := Rot(teamsContent[1], 13)
	//
	//teams := make([]*Team, 0, 8)
	//
	//err = json.Unmarshal([]byte(teamJson), &teams)
	//if err != nil {
	//	fmt.Printf("team Unmarshal failed. err: %s", err)
	//	return
	//}

	return
}
