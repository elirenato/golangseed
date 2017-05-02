package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"time"
)

type User struct {
	Id 					*int64
	Login				*string
	First_Name			*string
	Last_name			*string
	Email				*string
	Image_url			*string
	Activated			*bool	`json:"-"`
	Lang_key			*string
	Activation_key		*string	`json:"-"`
	Reset_key			*string	`json:"-"`
	Created_by			*string	`json:"-"`
	Created_date		*time.Time	`json:"-"`
	Reset_date			*time.Time	`json:"-"`
	Last_modified_by	*string	`json:"-"`
	Last_modified_date	*time.Time	`json:"-"`
	Password			*string	`json:"-"`
	Password_Hash		*string	`json:"-"`
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Login)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Login,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, *user.Password).
		Key("user.Password")

	v.Check(user.First_Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidateLogin(v *revel.Validation, login string) *revel.ValidationResult {
	return v.Check(login,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{5},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
