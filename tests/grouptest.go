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

func (t *GroupTest) Test001CreateSuccess() {
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

func (t *GroupTest) Test003CreateInvalid() {
	newGroup := &models.Group{
			Id: null.NewInt(1, true),
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
	t.AssertContains("\"error\":\"error.generic.invalidCreateId\"")
}

func (t *GroupTest) Test200GetSuccess() {
	url := fmt.Sprintf("/group/%d", lastCreatedGroup.Id.Int64)
	req := t.GetCustom(t.Url(url))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", defaultGroupName))
}

func (t *GroupTest) Test300UpdateSucess() {
	updatedGroupName := "Second group name"
	updateGroup := &models.Group{
			Id: lastCreatedGroup.Id,
			Name: null.NewString(updatedGroupName, true),
		}
	jsonByte, _ := json.Marshal(updateGroup)
	jsonString := string(jsonByte)
	req := t.PutCustom(t.Url("/group"),
			commons.ApplicationJsonContentType, strings.NewReader(jsonString))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", updatedGroupName))
	commons.DecodeJson(t.ResponseBody, &lastCreatedGroup)
}

func (t *GroupTest) Test301UpdateNotFound() {
	updatedGroupName := "Second group name"
	updateGroup := &models.Group{
			Id: null.NewInt(999999, true),
			Name: null.NewString(updatedGroupName, true),
		}
	jsonByte, _ := json.Marshal(updateGroup)
	jsonString := string(jsonByte)
	req := t.PutCustom(t.Url("/group"),
			commons.ApplicationJsonContentType, strings.NewReader(jsonString))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertStatus(http.StatusNotFound)
	t.AssertContains("\"error\":\"error.generic.notFound\"")
}

func (t *GroupTest) Test302UpdateInvalid() {
	updatedGroupName := "Second group name"
	updateGroup := &models.Group{
			Name: null.NewString(updatedGroupName, true),
		}
	jsonByte, _ := json.Marshal(updateGroup)
	jsonString := string(jsonByte)
	req := t.PutCustom(t.Url("/group"),
			commons.ApplicationJsonContentType, strings.NewReader(jsonString))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContains("\"error\":\"error.generic.invalidUpdateId\"")
}

func (t *GroupTest) Test400DeleteSucess() {
	url := fmt.Sprintf("/group/%d", lastCreatedGroup.Id.Int64)
	req := t.DeleteCustom(t.Url(url))
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContains(fmt.Sprintf("\"Name\":\"%s\"", lastCreatedGroup.Name.String))
}

func (t *GroupTest) Test401UpdateNotFound() {
	url := fmt.Sprintf("/group/%d", lastCreatedGroup.Id.Int64)
	req := t.DeleteCustom(t.Url(url))
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertStatus(http.StatusNotFound)
	t.AssertContains("\"error\":\"error.generic.notFound\"")
}
