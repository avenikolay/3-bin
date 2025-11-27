package main

import (
	"3-bin/api"
	_ "3-bin/api"
	"3-bin/bins"
	"3-bin/config"
	"3-bin/file"
	"3-bin/storage"
	"fmt"
)

func main() {
	fileDb, err := file.NewFileDb("bin.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	bin, err := bins.ReadBin(fileDb)
	if err != nil {
		fmt.Println(err)
		return
	}
	key, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	binStorage := storage.NewStorage()
	binStorage.PutBin(bin)

	bin = bins.MakeBin("1", true, "test")
	binStorage.PutBin(bin)

	binsList := binStorage.GetBins()
	fmt.Println(*binsList)

	api.SendRequest(bin, key)
}
