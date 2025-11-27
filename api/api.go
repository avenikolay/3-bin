package api

import (
	"3-bin/bins"
	"3-bin/config"
	"fmt"
)

func SendRequest(bin *bins.Bin, key *config.Config) {
	fmt.Println("Request ready to send:")
	fmt.Println("BIN: ", bin)
	fmt.Println("KEY: ", key)
}
