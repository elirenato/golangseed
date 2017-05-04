package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"log"
)

type Users struct {
	GorpController
}

type ResultError struct {
	Error *string
	Authenticated *bool
	Token *string
}

func (c Users) Index() revel.Result {
	var users []models.User
	_, err := Dbm.Select(&users,"select * from users order by id")
	if err != nil  {
		panic(err)
	}	
	return c.RenderJSON(users)
}

func (c Users) Authenticate() revel.Result {
	login, password := c.Params.Get("login"), c.Params.Get("password") 
	models.ValidateLogin(c.Validation, login)
	models.ValidatePassword(c.Validation, password)
	errorMap := make(map[string]interface{})
	result := make(map[string]interface{})
	result["authenticated"] = false
	result["error"] = &errorMap
	if !c.Validation.HasErrors() {
		var user models.User
		err := Dbm.SelectOne(&user, "select * from users where login=$1", login)
		if err == nil  {
			err = bcrypt.CompareHashAndPassword([]byte(*user.Password_Hash), []byte(password))
			if err == nil {
				var token string
				token, err = middleware.CreateToken(*user.Id,[]string{})
				if err == nil {
					result["authenticated"] = true
					result["token"] = token
					delete(result, "error")
				} else {
					errorMap["type"] = "internal-server-error"
					log.Print(err)
				}
			} else {
				log.Print(err)
				errorMap["type"] = "invalid-credentials"
			}
		} else {
			errorMap["type"] = "invalid-credentials"
			log.Print(err)
		}
	} else {
		errorMap["type"] = "invalid-credentials"
	}
	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"
	return c.RenderJSON(result)
}