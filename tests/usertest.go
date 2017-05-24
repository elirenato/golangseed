package tests

import (
	"encoding/json"
	"strings"
	"github.com/elirenato/golangseed/app/commons"
	"net/http"
	"fmt"
)

type UserTest struct {
	BaseTest
}

const (
	TestUserEmail = "testuser@localhost.com"
	TestUserFirstName = "UserFirstName"
	InvalidTestUserEmail = "noexistentuser@localhost"
)

func (t *UserTest) Test001Register() {
	body := make(map[string]interface{})
	body["email"] = TestUserEmail
	body["firstName"] = TestUserFirstName
	body["password"] = "123456"
	request,_ := json.Marshal(body)
	t.Post("/register", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"registered\":true")
}

func (t *UserTest) Test002RegisterEmptyField() {
	body := make(map[string]interface{})
	body["email"] = TestUserEmail
	// body["firstName"] = "Eli"
	body["password"] = "123456"
	request,_ := json.Marshal(body)
	t.Post("/register", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"error\":\"error.registration.validation\"")	
}

func (t *UserTest) Test003RegisterUserAlreadyExists() {
	body := make(map[string]interface{})
	body["email"] = TestUserEmail
	body["firstName"] = "Eli 2"
	body["password"] = "789611154"
	request,_ := json.Marshal(body)
	t.Post("/register", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"error\":\"error.registration.alreadyExists\"")	
}

func (t *UserTest) Test100Authentication() {
	body := make(map[string]interface{})
	body["email"] = TestUserEmail
	body["password"] = "123456"
	request,_ := json.Marshal(body)
	t.Post("/authenticate", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"authenticated\":true")
	t.AssertContains("\"token\":")
}

func (t *UserTest) Test101AuthenticationInvalidCredentials() {
	body := make(map[string]interface{})
	body["email"] = InvalidTestUserEmail
	body["password"] = "123456"
	request,_ := json.Marshal(body)
	t.Post("/authenticate", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"error\":\"error.authentication\"")
}

func (t *UserTest) Test201GetUserSuccess() {
	url := fmt.Sprintf("/users/%s", TestUserEmail)
	req := t.GetCustom(t.Url(url))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertOk()
	t.AssertContains(fmt.Sprintf("\"Email\":\"%s\"", TestUserEmail))
	t.AssertContains(fmt.Sprintf("\"FirstName\":\"%s\"", TestUserFirstName))
}

func (t *UserTest) Test202GetNoExistentUser() {
	url := fmt.Sprintf("/users?email=%s", InvalidTestUserEmail)
	req := t.GetCustom(t.Url(url))		
	t.setAuthorization(&req.Header)
	req.Send()
	t.AssertNotFound()
}