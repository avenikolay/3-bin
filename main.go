package main

import (
	_ "3-bin/api"
	_ "3-bin/bins"
	"3-bin/storage"
	"fmt"
)

func main() {
	//var BinList []storage.Bin
	//BinList = append(BinList, *testBin)

	testBin := storage.MakeBin("1", true, "test")
	storage.SaveBin(testBin)

	loadedBin, err := storage.ReadBin()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loadedBin)
}
