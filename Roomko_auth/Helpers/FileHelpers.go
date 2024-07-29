package Helpers

import (
	"fmt"
	"os"
)

func SaveToFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Error creating file\n", err)
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
