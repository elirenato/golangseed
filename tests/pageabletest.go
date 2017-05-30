package tests

import (
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/null"
	"github.com/elirenato/golangseed/app/repositories"
	"github.com/elirenato/golangseed/app/models"
	"reflect"
)

type PageableTest struct {
	BaseTest
}

type PageableTestStruct struct {
	Id 			null.Int  `db:"id, primarykey, autoincrement"`
	Field1			null.String `json:"firstField, ommitempty" db:"first_name"`
	Field2			null.String `json:"-"`
	Field3 			null.String
}

func (t *PageableTest) Test001MapFieldByJsonField() {
	params := commons.MapDbFieldByJsonName(reflect.TypeOf(PageableTestStruct{}))
	t.Assert(params["firstField"]=="first_name")
	t.Assert(params["Field2"]=="")
	t.Assert(params["Field3"]=="Field3")
	t.Assert(params["Id"]=="id")
}

func (t *PageableTest) Test002CreateOrderByStatmentSuccess() {
	sort1 := models.SortOption {
				Property: "firstField",
				Direction: models.DirectionDesc,
			}
	sort2 := models.SortOption {
				Property: "Field3",
				Direction: models.DirectionAsc,
			}
	page := models.Pageable{
			Sort: []models.SortOption{sort1, sort2 },
		}
	dummyInstance := repositories.BaseRepository{}
	dummyInstance.SetDependencies(nil, "", reflect.TypeOf(PageableTestStruct{}))
	orderByStatement, err := dummyInstance.CreateOrderByStatement(page)
	t.Assert(err == nil)
	t.AssertEqual("order by first_name desc, Field3", orderByStatement)
}

func (t *PageableTest) Test003CreateOrderByStatmentErrorNotExistentField() {
	sort1 := models.SortOption {
		Property: "notExistentField",
		Direction: models.DirectionDesc,
	}
	sort2 := models.SortOption {
		Property: "Field3",
		Direction: models.DirectionAsc,
	}
	page := models.Pageable{
		Sort: []models.SortOption{sort1, sort2 },
	}
	dummyInstance := repositories.BaseRepository{}
	dummyInstance.SetDependencies(nil, "", reflect.TypeOf(PageableTestStruct{}))
	orderByStatement, err := dummyInstance.CreateOrderByStatement(page)
	t.Assert(err != nil)
	t.AssertEqual("", orderByStatement )
	t.AssertEqual(err.Error(), "Sort field name 'notExistentField' is not valid")
}

func (t *PageableTest) Test004CreateOrderByStatmentErrorInvalidOrderField() {
	sort1 := models.SortOption {
		Property: "firstField",
		Direction: "des",
	}
	sort2 := models.SortOption {
		Property: "Field3",
		Direction: models.DirectionAsc,
	}
	page := models.Pageable{
		Sort: []models.SortOption{sort1, sort2 },
	}
	dummyInstance := repositories.BaseRepository{}
	dummyInstance.SetDependencies(nil, "", reflect.TypeOf(PageableTestStruct{}))
	orderByStatement, err := dummyInstance.CreateOrderByStatement(page)
	t.Assert(err != nil)
	t.AssertEqual("", orderByStatement )
	t.AssertEqual(err.Error(), "Sort direction of the field 'firstField' is not valid. It should be 'asc' or 'desc'")
}
