package commons

import (
	 "strings"
	 "encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"fmt"
	"github.com/revel/revel"
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

func MapDbFieldByJsonName(structType reflect.Type) (map[string]string) {
	params := map[string]string{}
    	for i := 0; i < structType.NumField(); i++ {
        // prints empty line if there is no json tag for the field
		field := structType.Field(i)
		jsonTag := field.Tag.Get("json")
		tags := strings.Split(jsonTag, ",")
		if tags[0] != "-" {
			//if the field is not ommited
			if tags[0] == "" {
				jsonTag = field.Name
			} else {
				jsonTag = tags[0]
			}
		} else {
			jsonTag = ""
		}
		if jsonTag != "" {
			dbTag := field.Tag.Get("db")
			tags = strings.Split(dbTag, ",")
			if tags[0] != "-" {
				if tags[0] == "" {
					params[jsonTag] = field.Name
				} else {
					params[jsonTag] = tags[0]
				}
			}
		}
	}
	return params;
}

func GenerateUri(endpoint string, params *map[string]string) string {
	var protocol string
	if revel.Config.BoolDefault("http.ssl", false) {
		protocol = "https"
	} else {
		protocol = "http"
	}
	address := revel.Config.StringDefault("http.addr", "")
	if IsBlank(address) {
		address = "localhost"
	}
	port := revel.Config.StringDefault("http.port", "9000")
	baseUrl := fmt.Sprintf("%s://%s:%s", protocol, address, port)
	endpoint = strings.Replace(endpoint, baseUrl, "?", -1)
	uri := fmt.Sprintf("%s%s", baseUrl, endpoint)
	if params != nil {
		first := true
		for k, v := range *params {
			if first {
				uri = uri + "?" + k + "=" + v
				first = false
			} else {
				uri = uri + "&" + k + "=" + v
			}
		}
	}
	return uri
}
