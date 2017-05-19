package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"github.com/elirenato/null"
)

const (
	UserUniqueKeyName = "user_email_key"
)

type User struct {
	Id 					null.Int  `db:"id, primarykey, autoincrement"`
	FirstName			null.String `db:"first_name"`
	LastName			null.String `db:"last_name"`
	Email				null.String` db:"email"`
	ImageUrl			null.String `db:"image_url"`
	Activated			null.Bool	`json:"-" db:"activated"`
	Language			null.String `db:"lang_key"`
	ActivationKey		null.String	`json:"-" db:"activation_key"`
	Resetkey			null.String	`json:"-" db:"reset_key"`
	CreatedDate			null.Time	`json:"-" db:"created_date"`
	ResetDate			null.Time	`json:"-" db:"reset_date"`
	LastModifiedDate	null.Time	`json:"-" db:"last_modified_date"`
	PasswordHash		null.String	`json:"-" db:"password_hash"`
	Password			null.String	`json:"-" db:"-"`
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
