package repositories

import (
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/gorp"
	"reflect"
)

type GroupRepository struct {
	BaseRepository
}

func NewGroupRepository(dbm **gorp.DbMap) GroupRepository {
	instance := GroupRepository{}
	instance.SetDependencies(dbm, "groups", reflect.TypeOf(models.Group{}))
	return instance
}

func (c GroupRepository) FindOne(id int64) (*models.Group, error) {
	var modelPtr models.Group
	result, err := c.findOne(id, &modelPtr)
	if !result {
		return nil, err
	}
	return &modelPtr, nil
}

func (c GroupRepository) List(pageble models.Pageable, ownerId int64) (*models.Page, error) {
	params := map[string]interface{}{}
	params["owner_id,="] = ownerId
	return c.PageableList(pageble, params)
}
