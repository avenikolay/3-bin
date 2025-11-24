package storage

import (
	"3-bin/file"
	"encoding/json"
	"time"
)

type Bin struct {
	Id          string    `json:"id"`
	Private     bool      `json:"private"`
	CreatedDate time.Time `json:"createdDate"`
	Name        string    `json:"name"`
}

func MakeBin(id string, private bool, name string) *Bin {
	createdDate := time.Now()
	return &Bin{Id: id, Private: private, CreatedDate: createdDate, Name: name}
}

func SaveBin(bin *Bin) error {
	jsonStruct, err := json.Marshal(bin)
	err = file.WriteFile(jsonStruct, "bin.json")
	return err
}

func ReadBin() (*Bin, error) {
	file, err := file.ReadJsonFile("bin.json")
	if err != nil {
		return nil, err
	}
	var bin Bin
	err = json.Unmarshal(file, &bin)
	return &bin, err
}
