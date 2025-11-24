package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadJsonFile(name string) ([]byte, error) {

	if strings.HasSuffix(name, ".json") == false {
		return nil, errors.New("FILE_IS_NOT_JSON")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) error {

	if strings.HasSuffix(name, ".json") == false {
		err := errors.New("FILE_IS_NOT_JSON")
		return err
	}

	file, err := os.Create(name)
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
