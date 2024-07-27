package Helpers

import "os"

func SaveToFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
