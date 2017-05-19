package controllers

import (
	"fmt"
	"database/sql"
	"github.com/elirenato/gorp"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/models"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {	
	userName := revel.Config.StringDefault("db.user","username")
	password := revel.Config.StringDefault("db.password","password")
	database := revel.Config.StringDefault("db.database","database")
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable",
	 	userName, password, database)
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}
	err = db.Ping();
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//add table user to ORM. The name of the table is users to avoid conflict with "select * from User" of postgres
	t := Dbm.AddTableWithName(models.User{},"users").SetKeys(true, "Id")
	t.ColMap("Password").Transient = true

	t = Dbm.AddTableWithName(models.Group{},"groups").SetKeys(true, "Id")
	
	Dbm.TraceOn("[gorp]", revel.INFO)
}