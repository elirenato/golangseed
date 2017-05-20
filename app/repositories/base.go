package repositories

import (
	"log"
	"database/sql"
	"github.com/elirenato/gorp"
	"fmt"
	"errors"
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