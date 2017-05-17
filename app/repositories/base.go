package repositories

import (
	"log"
	"database/sql"
	"github.com/go-gorp/gorp"
	"fmt"
	"errors"
)

type BaseRepository struct {
	Dbm **gorp.DbMap
	TableName string
}

func (c BaseRepository) findOne(id int64, entityPtr interface{}) (bool, error) {
	err := (*c.Dbm).SelectOne(entityPtr, fmt.Sprintf("select * from %s where id=%d", c.TableName, id))
	if err != nil  {
		if sql.ErrNoRows == err {
			return false, nil
		}
		log.Println(err)
		return false, err
	}	
	return true, nil
}

func (c BaseRepository) Update(modelPointer interface{}) (error) {
	updateCount, err := (*c.Dbm).Update(modelPointer)
	if updateCount != 1 {
		return errors.New("Update count does not match with the value returned")
	}
	return err
}

func (c BaseRepository) Delete(modelPointer interface{}) (error) {
	updateCount, err := (*c.Dbm).Delete(modelPointer)
	if updateCount != 1 {
		return errors.New("Delete count does not match with the value returned")
	}
	return err
}
