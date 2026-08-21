package main

import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return Student{}, err
	}

	var student Student

	err = json.Unmarshal(data, &student)
	if err != nil {
		return Student{}, err
	}

	return student, nil

}
