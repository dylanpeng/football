package main

import (
	"fmt"
	"football/lib/http"
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

	fmt.Printf("get http success\nrsp: %s", rsp)

	return
}
