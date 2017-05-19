package tests

import (
	"encoding/json"
	"strings"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/null"
	"fmt"
	"net/http"
)

type GroupTest struct {
	BaseTest
}

const (
	defaultGroupName = "First group name"
)

var (
	lastCreatedGroup *models.Group
)

func (t *GroupTest) Test001Create() {
	lastCreatedGroup = &models.Group{
			Name: null.NewString(defaultGroupName, true),
		}
	jsonByte, _ := json.Marshal(lastCreatedGroup)
	jsonString := string(jsonByte)
	req := t.PostCustom(t.Url("/group"), 
		commons.ApplicationJsonContentType, strings.NewReader(jsonString))
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", defaultGroupName))
	//update the default model to had the id		
	commons.DecodeJson(t.ResponseBody, &lastCreatedGroup)
			fmt.Println("### 1");
	fmt.Println(lastCreatedGroup.Id.Int64);

}

func (t *GroupTest) Test002CreateDuplicated() {
	newGroup := &models.Group{
			Name: null.NewString(defaultGroupName, true),
		}
	jsonByte, _ := json.Marshal(newGroup)
	jsonString := string(jsonByte)
	req := t.PostCustom(t.Url("/group"), 
		commons.ApplicationJsonContentType, strings.NewReader(jsonString))
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"error\":\"error.group.alreadyExists\"")

}

func (t *GroupTest) Test200GetExistent() {
	url := fmt.Sprintf("/group/%d", lastCreatedGroup.Id.Int64)
	fmt.Println(url)
	req := t.GetCustom(t.Url(url))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", defaultGroupName))
}