package main

import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
	jsonWriter := json.NewEncoder(w)

	err := jsonWriter.Encode(students)
	if err != nil {
		return err
	}

	return nil
}
