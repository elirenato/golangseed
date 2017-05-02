package middleware

import (
	"github.com/revel/revel"
	"strings"
	"io/ioutil"
	"encoding/json"
)

var JsonParamsFilter = func(c *revel.Controller, fc []revel.Filter) {
    if strings.Contains(c.Request.ContentType, "application/json") {
        data := map[string]string{}
        content, _ := ioutil.ReadAll(c.Request.Body)
        err := json.Unmarshal(content, &data)
		if err != nil {
			panic(err)
		}
        for k, v := range data {
            c.Params.Values.Set(k, v)
        }
    }
    fc[0](c, fc[1:]) // Execute the next filter stage.
}