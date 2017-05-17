package models

import (
	"fmt"
	"github.com/revel/revel"
	"time"
)

type Group struct {
	Id 					*int64  `db:"id, primarykey, autoincrement"`
	Name				*string `db:"name"`
	ImageUrl			*string `db:"image_url"`
	CreatedDate			*time.Time	`json:"-" db:"created_date"`
	LastModifiedDate	*time.Time	`json:"-" db:"last_modified_date"`
}

func (u *Group) String() string {
	return fmt.Sprintf("[Name: %s]", u.Name)
}

func (u *Group) Validate(v *revel.Validation) {
	u.ValidateName(v, u.Name)
}

func (u *Group) ValidateName(v *revel.Validation, namePointer *string) *revel.ValidationResult {	
	var name string
	if namePointer == nil {
		name = ""
	} else {
		name = *namePointer
	}
	return v.Check(name,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{5},
	)
}