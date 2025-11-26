package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type File interface {
	ReadJsonFile() ([]byte, error)
	WriteJsonFile(content []byte) error
}

type JsonFile struct {
	name string
}

func NewFileDb(filename string) (*JsonFile, error) {
	if strings.HasSuffix(filename, ".json") == false {
		return nil, errors.New("FILE_IS_NOT_JSON")
	}
	return &JsonFile{
		name: filename,
	}, nil
}

func (fileDB *JsonFile) ReadJsonFile() ([]byte, error) {
	data, err := os.ReadFile(fileDB.name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fileDB *JsonFile) WriteJsonFile(content []byte) error {
	file, err := os.Create(fileDB.name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}
