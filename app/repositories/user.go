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

func NewUserRepository(dbm **gorp.DbMap) UserRepository {
	instance := UserRepository{}
	instance.TableName = "users"
	instance.Dbm = dbm;
	return instance
}


func (c UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := (*c.Dbm).SelectOne(&user, "select * from users where email=$1", email)
	if err != nil  {
		if sql.ErrNoRows == err {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}	
	return &user, nil	
}