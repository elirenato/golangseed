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
	var entity models.Group
	result, err := c.findOne(id, &entity)
	if !result {
		return nil, err
	} 
	return &entity, nil
}

func (c BaseRepository) Persist(dbm *gorp.DbMap, modelPointer interface{}) (error) {
	return dbm.Insert(modelPointer)
}