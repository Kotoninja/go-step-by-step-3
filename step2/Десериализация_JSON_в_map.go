package main

import (
	"encoding/json"
)

func DeserializeStringMap(data string) (map[string]string, error) {
	ans := map[string]string{}

	err := json.Unmarshal([]byte(data), &ans)
	if err != nil {
		return map[string]string{}, err
	}

	return ans, nil
}

// func main() {
// 	value := map[string]string{"1": "1", "2": "2"}

// }
