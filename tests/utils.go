package tests

import (
	"encoding/json"
)

func MapToJson(o interface{}) string {
	result, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(result)
}
