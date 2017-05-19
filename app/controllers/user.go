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
	"github.com/elirenato/null"
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
		log.Println(fmt.Sprintf("User not found with the email %s", email))
		return c.RenderNotFound()
	}
	return c.RenderOK(user)
}

func (c UserController) Authenticate() revel.Result {
	params := map[string]string{}
	err := commons.DecodeJsonBody(c.Request.Body, &params)
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("error.generic.invalidBody")
	}
	email, password := params["email"], params["password"]
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
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
	}
	if err == nil  {
		token, err = filters.CreateToken(user.Id.Int64,[]string{})	
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
	params := map[string]string{}
	err := commons.DecodeJsonBody(c.Request.Body, &params)
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("error.generic.invalidBody")
	}	
	firstName := params["firstName"]
	email := params["email"]
	password := params["password"]
	models.ValidateFirstName(c.Validation, firstName)
	models.ValidateEmail(c.Validation, email)
	models.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.registration.validation")
	}
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	user := &models.User{
		FirstName: null.NewString(firstName, true),
		Email: null.NewString(email, true),
		Activated: null.NewBool(true, true), //TODO create the activation flow by email
		Language: null.NewString("en-US", true), //TODO option to change language
		CreatedDate: null.NewTime(time.Now(), true),
		PasswordHash: null.NewString(string(hashedPassword), true),
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
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
	var token string
	if err == nil  {
		token, err = filters.CreateToken(user.Id.Int64,[]string{})	
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