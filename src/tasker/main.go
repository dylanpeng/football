package main

import (
	"fmt"
	"football/lib/leisu"
)

func main() {
	match, err := leisu.QueryMatch()

	if err != nil {
		fmt.Printf("QueryMatch failed. match: %s | err: %s", match, err)
		return
	}

	return
}
