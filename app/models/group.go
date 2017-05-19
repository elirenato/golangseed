package models

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/elirenato/null"
)

type Group struct {
	Id 					null.Int  `db:"id, primarykey, autoincrement"`
	Name				null.String `db:"name"`
	ImageUrl			null.String `db:"image_url"`
	CreatedDate			null.Time	`json:"-" db:"created_date"`
	LastModifiedDate	null.Time	`json:"-" db:"last_modified_date"`
}

const (
	GroupUniqueName = "name_key"
)

func (u *Group) String() string {
	return fmt.Sprintf("[Name: %s]", u.Name)
}

func (u *Group) Validate(v *revel.Validation) {
	u.ValidateName(v)
}

func (u *Group) ValidateName(v *revel.Validation) *revel.ValidationResult {
	return v.Check(
		u.Name.String,
		revel.Required{},
	)
}

func (u *Group) SetLastInsertID(value int64) {
	u.Id = null.NewInt(value, true)
}