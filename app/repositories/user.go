package repositories

import (
	"github.com/elirenato/golangseed/app/models"
	"log"
	"github.com/go-gorp/gorp"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository() UserRepository {
	return UserRepository{} 
}

func (c UserRepository) ListAll(dbm *gorp.DbMap) ([]models.User, error) {
	var users []models.User
	_, err := dbm.Select(users, "select * from users order by id")
	if err != nil  {
		log.Println(err)
		return nil, err
	}	
	return users, err
}

func (c UserRepository) FindByEmail(dbm *gorp.DbMap, email string) (*models.User, error) {
	var user models.User
	err := dbm.SelectOne(&user, "select * from users where email=$1", email)
	if err != nil  {
		log.Println(err)
		return nil, err
	}	
	return &user, nil	
}

func (c UserRepository) Persist(dbm *gorp.DbMap, values interface{}) (error) {
	return dbm.Insert(values)
}