package repositories

import (
	"github.com/elirenato/golangseed/app/models"
	"log"
	"github.com/elirenato/gorp"
	"database/sql"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository() UserRepository {
	return UserRepository{} 
}

func (c UserRepository) GetByEmail(dbm *gorp.DbMap, email string) (*models.User, error) {
	var user models.User
	err := dbm.SelectOne(&user, "select * from users where email=$1", email)
	if err != nil  {
		if sql.ErrNoRows == err {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}	
	return &user, nil	
}

func (c UserRepository) Persist(dbm *gorp.DbMap, values interface{}) (error) {
	return dbm.Insert(values)
}