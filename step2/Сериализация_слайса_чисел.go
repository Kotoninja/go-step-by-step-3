package main

import (
	"encoding/json"
)

func SerializeIntSlice(nums []int) ([]byte, error) {
	result, err := json.Marshal(nums)
	if err != nil {
		return []byte{}, nil
	}

	return result, nil
}
