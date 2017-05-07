package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/filters"
	"github.com/elirenato/golangseed/app/services"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UsersController struct {
	BaseController
}

var service = services.UsersService{}

func (c UsersController) Index() revel.Result {
	users, err := service.ListAll(Dbm)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	return c.RenderOK(users)
}

func (c UsersController) Authenticate() revel.Result {
	login, password := c.Params.Get("login"), c.Params.Get("password") 
	models.ValidateLogin(c.Validation, login)
	models.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("authentication.invalidCredentials")
	}
	var token string
	user, err := service.FindUserByLogin(Dbm, login)
	if err == nil  {
		err = bcrypt.CompareHashAndPassword([]byte(*user.Password_Hash), []byte(password))
	}
	if err == nil  {
		token, err = filters.CreateToken(*user.Id,[]string{})	
	}
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("authentication.invalidCredentials")
	}	
	result := make(map[string]interface{})
	result["authenticated"] = true
	result["token"] = token
	return c.RenderOK(result)
}

func (c UsersController) RegisterAccount() revel.Result {
	//TOD IN PROGRESS
	login := c.Params.Get("login")
	user, err := service.FindUserByLogin(Dbm, login)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}	
	result := make(map[string]interface{})
	result["created"] = true
	result["user"] = user.Id;
	return c.RenderOK(result)	
}