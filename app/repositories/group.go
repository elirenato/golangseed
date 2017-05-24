package repositories

import (
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/gorp"
	"fmt"
	"bytes"
	"reflect"
	"database/sql"
	"log"
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

//Returns the query to count the records if the countQuery is true or returns the query to retrieve the models
func (c BaseRepository) createSelectQuery(countQuery bool, pageble models.Pageable) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("select")
	if countQuery {
		buffer.WriteString(" count(*) ")
	} else {
		buffer.WriteString(" * " )
	}
	buffer.WriteString(fmt.Sprintf("from %s",  c.TableName))
	if !countQuery {
		orderBy, err := CreateOrderByStatement(pageble, reflect.TypeOf(models.Group{}))
		if err != nil {
			return "", err
		}
		if orderBy != "" {
			buffer.WriteString(" ")
			buffer.WriteString(orderBy)
		}
		buffer.WriteString(" ")
		buffer.WriteString(createLimitStatement(pageble))
	}
	return buffer.String(), nil
}

func (c BaseRepository) ListAll(pageble models.Pageable) (*models.Page, error) {
	query, err := c.createSelectQuery(true, pageble)
	if err != nil {
		return nil, err
	}
	var items []models.Group
	var page models.Page = models.Page{
		TotalElements: 0,
		TotalPages: 0,
		PageNumber: pageble.Page,
		PageSize: pageble.Size,
		NumberOfElements: 0,
		Items: &items,
	}
	page.TotalElements, err = (*c.Dbm).SelectInt(query)
	if err != nil {
		return nil, err
	}
	query, err = c.createSelectQuery(false, pageble)
	if err != nil {
		return nil, err
	}
	_, err = (*c.Dbm).Select(&items, query)
	if err != nil  {
		if sql.ErrNoRows == err {
			return &page, nil
		}
		log.Println(err)
		return nil, err
	}
	page.NumberOfElements = int64(len(items))
	if page.PageSize > 0 {
		page.TotalPages = page.TotalElements / page.PageSize
	} else {
		page.TotalPages = 0
	}
	return &page, nil
}
