package services

import (
	"github.com/elirenato/golangseed/app/models"
	"log"
	"github.com/go-gorp/gorp"
)

type UsersService struct {
	BaseService
}

func (c UsersService) ListAll(dbm *gorp.DbMap) ([]models.User, error) {
	var users []models.User
	_, err := dbm.Select(&users,"select * from users order by id")
	if err != nil  {
		log.Println(err)
		return nil, err
	}	
	return users, nil
}

func (c UsersService) FindUserByLogin(dbm *gorp.DbMap, login string) (*models.User, error) {
	var user models.User
	err := dbm.SelectOne(&user, "select * from users where login=$1", login)
	if err != nil  {
		log.Println(err)
		return nil, err
	}	
	return &user, nil	
}