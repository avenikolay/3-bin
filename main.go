package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id          string
	private     bool
	createdDate time.Time
	name        string
}

func makeBin(id string, private bool, name string) (*Bin, error) {
	createdDate := time.Now()
	return &Bin{id, private, createdDate, name}, nil
}

func main() {
	var BinList []Bin
	testBin, error := makeBin("1", true, "test")

	if error != nil {
		fmt.Println(error)
	}
	BinList = append(BinList, *testBin)
	fmt.Println(BinList)
}
