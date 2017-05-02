package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Users struct {
	GorpController
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
	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusBadRequest
		c.Response.ContentType = "application/json"
		return c.RenderJSON(c.Validation.Errors)
	}
	var user models.User
	err := Dbm.SelectOne(&user, "select * from users where login=$1", login)
	if err != nil  {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password_Hash), []byte(password))
	if err == nil {
		var token string
		token, err = middleware.CreateToken(*user.Id,[]string{})
		return c.RenderJSON(token)
	}
	c.Response.Status = http.StatusBadRequest
	c.Response.ContentType = "application/json"
	return c.RenderJSON("invalid username of password")
}