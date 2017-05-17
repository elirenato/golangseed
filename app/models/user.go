package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"time"
)

const (
	UserUniqueKeyName = "user_email_key"
)

type User struct {
	Id 					*int64  `db:"id, primarykey, autoincrement"`
	FirstName			*string `db:"first_name"`
	LastName			*string `db:"last_name"`
	Email				*string` db:"email"`
	ImageUrl			*string `db:"image_url"`
	Activated			*bool	`json:"-" db:"activated"`
	Language			*string `db:"lang_key"`
	ActivationKey		*string	`json:"-" db:"activation_key"`
	Resetkey			*string	`json:"-" db:"reset_key"`
	CreatedDate			*time.Time	`json:"-" db:"created_date"`
	ResetDate			*time.Time	`json:"-" db:"reset_date"`
	LastModifiedDate	*time.Time	`json:"-" db:"last_modified_date"`
	PasswordHash		*string	`json:"-" db:"password_hash"`
	Password			*string	`json:"-" db:"-"`
}

func (u *User) String() string {
	return fmt.Sprintf("[Email: %s]", u.Email)
}

var emailPattern = regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")

func ValidateFirstName(v *revel.Validation, firstName string) *revel.ValidationResult {
	return v.Check(firstName,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{3},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

func ValidateEmail(v *revel.Validation, email string) *revel.ValidationResult {	
	return v.Check(email,
		revel.Required{},
		revel.Match{emailPattern},
	)
}
