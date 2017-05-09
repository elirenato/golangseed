package tests

import (
	"github.com/revel/revel/testing"
	"encoding/json"
	"strings"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/controllers"
	"database/sql"
	"github.com/go-gorp/gorp"
)

type UserTest struct {
	testing.TestSuite
}

var Txn *gorp.Transaction

func (t *UserTest) Before() {
	println("Set up")
	var err error
	Txn, err = controllers.Dbm.Begin()
	if err != nil {
		panic(err)
	}
}

func (t *UserTest) TestRegister() {
	body := make(map[string]interface{})
	body["email"] = "admin@localhost"
	body["firstName"] = "Eli"
	body["password"] = "123456"
	request,_ := json.Marshal(body)
	t.Post("/register", commons.ApplicationJsonContentType, strings.NewReader(string(request)))
	t.AssertOk()
	t.AssertContentType(commons.ApplicationJsonContentType)
}

func (t *UserTest) After() {
	println("Tear down")
	if err := Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
}
