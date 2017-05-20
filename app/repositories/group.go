package repositories

import (
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/gorp"
)

type GroupRepository struct {
	BaseRepository
}

func NewGroupRepository(dbm **gorp.DbMap) GroupRepository {
	instance := GroupRepository{}
	instance.TableName = "groups"
	instance.Dbm = dbm;
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