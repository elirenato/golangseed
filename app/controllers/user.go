package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/filters"
	"github.com/elirenato/golangseed/app/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"database/sql"
)

type UserController struct {
	BaseController
}

var userRepository = repositories.NewUserRepository()

func (c UserController) Index() revel.Result {
	users, err := userRepository.ListAll(Dbm)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	return c.RenderOK(users)
}

func (c UserController) Authenticate() revel.Result {
	email, password := c.Params.Get("email"), c.Params.Get("password") 
	models.ValidateEmail(c.Validation, email)
	models.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.authentication")
	}
	var token string
	user, err := userRepository.FindByEmail(Dbm, email)
	if err == nil  {
		err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password))
	}
	if err == nil  {
		token, err = filters.CreateToken(*user.Id,[]string{})	
	}
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("error.authentication")
	}	
	result := make(map[string]interface{})
	result["authenticated"] = true
	result["token"] = token
	return c.RenderOK(result)
}

func (c UserController) Register() revel.Result {
	firstName := c.Params.Get("firstName")
	email := c.Params.Get("email") 
	password := c.Params.Get("password")
	models.ValidateFirstName(c.Validation, firstName)
	models.ValidateEmail(c.Validation, email)
	models.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.registration")
	}
	user, err := userRepository.FindByEmail(Dbm, email)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	user = &models.User{
		// Id: nil,
		FirstName: commons.InitializeString(firstName),
		// LastName: nil,
		Email: commons.InitializeString(email),
		// ImageUrl: nil,
		Activated: commons.InitializeBool(true), //TODO create the activation flow by email
		Language: commons.InitializeString("en-US"), //TODO option to change language
		// ActivationKey: nil,
		// Resetkey: nil,
		// ResetDate: nil,
		CreatedDate: commons.InitializeTime(time.Now()),
		// LastModifiedDate: nil,
		PasswordHash: commons.InitializeString(string(hashedPassword)),
		// Password: nil,
	}
	err = userRepository.Persist(Dbm, user);
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
 	result := make(map[string]interface{})
	result["created"] = true
	return c.RenderOK(result)	
}