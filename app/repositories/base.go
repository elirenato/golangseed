package repositories

import (
	"log"
	"database/sql"
	"github.com/elirenato/gorp"
	"fmt"
	"errors"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/models"
	"reflect"
	"strings"
)

type BaseRepository struct {
	Dbm **gorp.DbMap
	TableName string
}

func (c BaseRepository) findOne(id int64, modelPointer interface{}) (bool, error) {
	err := (*c.Dbm).SelectOne(modelPointer, fmt.Sprintf("select * from %s where id=%d", c.TableName, id))
	if err != nil  {
		if sql.ErrNoRows == err {
			return false, nil
		}
		log.Println(err)
		return false, err
	}	
	return true, nil
}

func (c BaseRepository) Delete(modelPointer interface{}) (error) {
	count, err := (*c.Dbm).Delete(modelPointer)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New(fmt.Sprintf("Delete count does not match with the value returned %d", count))
	}
	return nil
}

func (c BaseRepository) Persist(modelPointer interface{}) (error) {
	return (*c.Dbm).Insert(modelPointer)
}

func (c BaseRepository) Update(modelPointer interface{}) (error) {
	updateCount, err := (*c.Dbm).Update(modelPointer)
	if err != nil {
		return err
	}	
	if updateCount != 1 {
		return errors.New(fmt.Sprintf("Update count does not match with the value returned %d", updateCount))
	}
	return nil
}

func CreateOrderByStatement(page models.Pageable, modelType reflect.Type) (string, error) {
	jsonFieldsMap := commons.MapDbFieldByJsonName(modelType)
	sortStatement := ""
	if len(page.Sort) > 0 {
		for _, sort := range page.Sort {
			fieldName := jsonFieldsMap[sort.Property]
			if fieldName == "" {
				err := fmt.Errorf("Sort field name '%s' is not valid", sort.Property)
				fmt.Println(err)
				return "", err
			}
			if sortStatement == "" {
				sortStatement = sortStatement + "order by "
			} else {
				sortStatement = sortStatement + ", "
			}
			if strings.EqualFold(sort.Direction, "desc") {
				sortStatement = sortStatement + fieldName + " desc"
			} else if strings.EqualFold(sort.Direction, "asc") || sort.Direction == "" {
				sortStatement = sortStatement + fieldName
			} else {
				err := fmt.Errorf("Sort direction of the field '%s' is not valid. It should be 'asc' or 'desc'", sort.Property)
				fmt.Println(err)
				return "", err
			}
		}
	}
	return sortStatement, nil
}

func createLimitStatement(page models.Pageable) (string) {
	return fmt.Sprintf("offset %d limit %d", page.Page * page.Size, page.Size)
}