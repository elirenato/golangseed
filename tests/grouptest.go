package tests

import (
	"encoding/json"
	"strings"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/models"
	"fmt"
)

type GroupTest struct {
	BaseTest
}

func (t *GroupTest) Test001Create() {
	model := &models.Group{
			Name: commons.InitializeString("Group 1"),
		}
	jsonByte, _ := json.Marshal(model)
	jsonString := string(jsonByte)
	fmt.Printf("sending %s ",jsonString)
	req := t.PostCustom(t.Url("/group"), commons.ApplicationJsonContentType, strings.NewReader(jsonString))
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", *model.Name))
}