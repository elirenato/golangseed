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
	"strings"
	"fmt"
)

type UserController struct {
	BaseController
}

var userRepository = repositories.NewUserRepository()

func (c UserController) Read(email string) revel.Result {
	if commons.IsBlank(email) {
		log.Println("Param required not informed")
		return c.RenderInternalServerError()
	}
	user, err := userRepository.GetByEmail(Dbm, email)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	if (user == nil) {
		return c.RenderNotFound()
	}
	return c.RenderOK(user)
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
	user, err := userRepository.GetByEmail(Dbm, email)
	if err == nil  {
		if user == nil {
			log.Println(c.Validation.Errors)
			return c.RenderBadRequest("error.authentication")
		}
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
	models.ValidateFirstName(c.Validation, firstName).Message(fmt.Sprintf("First name not filled: %s", firstName))
	models.ValidateEmail(c.Validation, email).Message("Email is not filled or is invalid")
	models.ValidatePassword(c.Validation, password).Message("Password not filled")
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.registration")
	}
	var err error
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	user := &models.User{
		FirstName: commons.InitializeString(firstName),
		Email: commons.InitializeString(email),
		Activated: commons.InitializeBool(true), //TODO create the activation flow by email
		Language: commons.InitializeString("en-US"), //TODO option to change language
		CreatedDate: commons.InitializeTime(time.Now()),
		PasswordHash: commons.InitializeString(string(hashedPassword)),
	}
	err = userRepository.Persist(Dbm, user);
	if err != nil {
		// TODO improve the error handler for user already exists
		// pqError := err.(*pq.Error)
		// pqError.Constraint is empty, why?
		if strings.Contains(err.Error(), models.UserUniqueKeyName) {
			log.Println(fmt.Sprintf("An user already exists with the name %s", email))
			return c.RenderBadRequest("error.registration.alreadyExists")
		}
		return c.RenderInternalServerError()
	}
	err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password))
	var token string
	if err == nil  {
		token, err = filters.CreateToken(*user.Id,[]string{})	
	}
	if err != nil {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.registration")
	}
 	result := make(map[string]interface{})
	result["registered"] = true
	result["token"] = token
	return c.RenderOK(result)
}