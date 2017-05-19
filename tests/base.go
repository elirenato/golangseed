package tests

import (
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"github.com/elirenato/golangseed/app/controllers"
	"github.com/elirenato/golangseed/app/filters"
	"github.com/elirenato/golangseed/app/repositories"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/commons"
	"log"
	"fmt"
	"net/http"
)

type BaseTest struct {
	testing.TestSuite
}

const (
	defaultUserEmail = "admin@localhost"
)

var (
	authenticationToken = "-"
	userRepository = repositories.NewUserRepository()
)
func (t *BaseTest) Before() {
	var err error	
	var user *models.User
	if authenticationToken == "-" {
		user, err = userRepository.GetByEmail(controllers.Dbm, defaultUserEmail)
		if err != nil {
			log.Fatal(err)
			return
		}
		if user == nil {
			log.Fatal(fmt.Sprintf("User %s not found", defaultUserEmail))
			return
		}
		authenticationToken, err = filters.CreateToken(user.Id.Int64,[]string{})	
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func (t *BaseTest) setAuthorization(header *http.Header) {
	header.Set("Authorization", fmt.Sprintf("Bearer %s", authenticationToken))
}

func (t *BaseTest) Url(endpoint string) string {
	var protocol string
	if revel.Config.BoolDefault("http.ssl", false) {
		protocol = "https"
	} else {
		protocol = "http"
	}
	url := revel.Config.StringDefault("http.addr", "")
	if commons.IsBlank(url) {
		url = "localhost"
	}
	port := revel.Config.StringDefault("http.port", "9000")
	return fmt.Sprintf("%s://%s:%s%s", protocol, url, port, endpoint)
}

// func (t *BaseTest) After() {
// 	println("Tear down")
// }

