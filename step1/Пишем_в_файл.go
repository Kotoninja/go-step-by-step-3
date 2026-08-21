package main

import (
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {

		return err
	}
	defer file.Close()

	file.Write([]byte(message))

	return nil
}
