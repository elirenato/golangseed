package repositories

import (
	"log"
	"database/sql"
	"github.com/elirenato/gorp"
	"fmt"
	"errors"
	"github.com/elirenato/golangseed/app/models"
	"reflect"
	"strings"
	"bytes"
	"sort"
	"github.com/elirenato/golangseed/app/commons"
)

type BaseRepository struct {
	Dbm **gorp.DbMap
	TableName string
	ModelType reflect.Type
	JsonFieldsMap map[string]string
}

func (c *BaseRepository) SetDependencies(dbm **gorp.DbMap, tableName string, modelType reflect.Type) {
	c.Dbm = dbm
	c.TableName = tableName
	c.ModelType = modelType
	c.JsonFieldsMap = commons.MapDbFieldByJsonName(modelType)
}

func (c *BaseRepository) findOne(id int64, modelPointer interface{}) (bool, error) {
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

func (c *BaseRepository) Delete(modelPointer interface{}) (error) {
	count, err := (*c.Dbm).Delete(modelPointer)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New(fmt.Sprintf("Delete count does not match with the value returned %d", count))
	}
	return nil
}

func (c *BaseRepository) Persist(modelPointer interface{}) (error) {
	return (*c.Dbm).Insert(modelPointer)
}

func (c *BaseRepository) Update(modelPointer interface{}) (error) {
	updateCount, err := (*c.Dbm).Update(modelPointer)
	if err != nil {
		return err
	}	
	if updateCount != 1 {
		return errors.New(fmt.Sprintf("Update count does not match with the value returned %d", updateCount))
	}
	return nil
}

func (c *BaseRepository) CreateOrderByStatement(page models.Pageable) (string, error) {
	sortStatement := ""
	if len(page.Sort) > 0 {
		for _, sort := range page.Sort {
			fieldName := c.JsonFieldsMap[sort.Property]
			if fieldName == "" {
				err := fmt.Errorf("Sort field name '%s' is not valid", sort.Property)
				log.Println(err)
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
				log.Println(err)
				return "", err
			}
		}
	}
	return sortStatement, nil
}

func (c *BaseRepository) PageableList(pageble models.Pageable, params map[string]interface{}) (*models.Page, error) {
	query, paramValues, err := c.CreatePageableQuery(true, pageble, params)
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
	page.TotalElements, err = (*c.Dbm).SelectInt(query, paramValues...)
	if err != nil {
		return nil, err
	}
	query, paramValues, err = c.CreatePageableQuery(false, pageble, params)
	if err != nil {
		return nil, err
	}
	_, err = (*c.Dbm).Select(&items, query, paramValues...)
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

func (c *BaseRepository) parseJsonParams(jsonParams map[string]string) (map[string]interface{}, error) {
	var keys []string
	for k := range jsonParams {
		keys = append(keys, k)
	}
	result := map[string]interface{}{}
	for _, k := range keys {
		fieldName := c.JsonFieldsMap[k]
		if fieldName == "" {
			err := fmt.Errorf("Param field name '%s' is not valid", k)
			log.Println(err)
			return nil, err
		}
		result[fieldName] = jsonParams[k]
	}
	return result, nil
}

//Returns the query to count the records if the countQuery is true or returns the query to retrieve the models
func (c *BaseRepository) CreatePageableQuery(countQuery bool, pageble models.Pageable, filterParams map[string]interface{}) (string, []interface{}, error) {
	var buffer bytes.Buffer
	buffer.WriteString("select")
	if countQuery {
		buffer.WriteString(" count(*) ")
	} else {
		buffer.WriteString(" * " )
	}
	buffer.WriteString(fmt.Sprintf("from %s",  c.TableName))
	var paramValues []interface{}
	if len(filterParams) > 0 {
		var keys []string
		for k := range filterParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		first := true
		index := 1
		for _, k := range keys {
			keyOperatorInfo := strings.Split(k,",")
			if first {
				buffer.WriteString(" where ")
				first = false
			} else {
				buffer.WriteString(" and ")
			}
			buffer.WriteString(fmt.Sprintf("%s%s$%d", keyOperatorInfo[0], keyOperatorInfo[1], index ))
			v := filterParams[k]
			paramValues = append(paramValues, v)
			index++
		}
	}
	if !countQuery {
		orderBy, err := c.CreateOrderByStatement(pageble)
		if err != nil {
			return "", nil, err
		}
		if orderBy != "" {
			buffer.WriteString(" ")
			buffer.WriteString(orderBy)
		}
		buffer.WriteString(" ")
		buffer.WriteString(c.CreateLimitStatement(pageble))
	}
	return buffer.String(), paramValues, nil
}


func (c *BaseRepository) CreateLimitStatement(page models.Pageable) (string) {
	return fmt.Sprintf("offset %d limit %d", page.Page * page.Size, page.Size)
}