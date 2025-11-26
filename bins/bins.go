package bins

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

func ReadBin(file file.File) (*Bin, error) {
	fileData, err := file.ReadJsonFile()
	if err != nil {
		return nil, err
	}
	var bin *Bin
	err = json.Unmarshal(fileData, &bin)
	return bin, err
}
