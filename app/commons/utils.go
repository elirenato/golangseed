package commons

import (
	 "strings"
	 "encoding/json"
	"io"
	"io/ioutil"
)

const (
	ApplicationJsonContentType = "application/json; charset=utf-8"
)

func IsBlank(value string) bool {
	value = strings.Trim(value, " ")
	return len(value) <= 0
}

func DecodeJson(data []byte, resultPtr interface{}) error {
	err := json.Unmarshal(data, resultPtr)
	if err != nil {
		return err
	}
	return nil
}

func DecodeJsonBody(r io.ReadCloser, resultPtr interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err;
	}
	return DecodeJson(data, resultPtr)
}
