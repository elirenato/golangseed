package tests

// TODO Right now the methods are persistent, I mean, the information that is create inside
// each method is not rolling back. Check if we are going to rolling back the methods at tearDown
// or run the methods in the correct order?

import (
	"github.com/revel/revel/testing"
	"github.com/elirenato/golangseed/app/filters"
	"github.com/elirenato/golangseed/app/repositories"
	"github.com/elirenato/golangseed/app/commons"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

type BaseTest struct {
	testing.TestSuite
}

const (
	defaultUserEmail = "admin@localhost"
)

var (
	authenticationToken = "-"
	userRepository = repositories.NewUserRepository(&commons.Dbm)
)

//Helper method to create new user. It is used by usertest and also another tests to create different accounts
func (t *BaseTest) CreateNewUser(userEmail, userFirstName, userPassword string) {
	body := make(map[string]interface{})
	body["email"] = userEmail
	body["firstName"] = userFirstName
	body["password"] = userPassword
	request,_ := json.Marshal(body)
	t.Post("/register", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
	t.AssertContains("\"registered\":true")
}


func (t *BaseTest) CreateAuthenticationToken(userEmail string) (string, error){
	user, err := userRepository.GetByEmail(userEmail)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	if user == nil {
		err = fmt.Errorf("User %s not found", userEmail)
		log.Fatal(err)
		return "", err
	}
	var token string
	token, err = filters.CreateToken(user.Id.Int64,[]string{})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return token, nil
}

func (t *BaseTest) Before() {
	if authenticationToken == "-" {
		var err error
		authenticationToken, err = t.CreateAuthenticationToken(defaultUserEmail)
		if err != nil {
			panic(err)
		}
	}
}

func (t *BaseTest) setAuthorization(header *http.Header) {
	t.setCustomAuthorization(header, authenticationToken)
}

func (t *BaseTest) setCustomAuthorization(header *http.Header, token string) {
	header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (t *BaseTest) Url(endpoint string) string {
	return commons.GenerateUri(endpoint, nil)
}
