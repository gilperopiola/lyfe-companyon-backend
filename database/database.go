package database

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gilperopiola/lyfe-companyon-backend/config"
	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseActions interface {
	Setup()
	Purge()
	Migrate()

	LoadTestingData()
	BeautifyError(error) string
}

type MyDatabase struct {
	*sql.DB
}

func (db *MyDatabase) Setup(cfg config.MyConfig) {
	var err error
	db.DB, err = sql.Open(
		cfg.DATABASE.TYPE, cfg.DATABASE.USERNAME+":"+cfg.DATABASE.PASSWORD+"@tcp("+cfg.DATABASE.HOSTNAME+":"+
			cfg.DATABASE.PORT+")/"+cfg.DATABASE.SCHEMA+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	err = db.DB.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	if cfg.DATABASE.CREATE_SCHEMA {
		db.CreateSchema()
	}

	if cfg.DATABASE.PURGE {
		db.Purge()
	}

	if cfg.DATABASE.CREATE_ADMIN {
		db.CreateAdmin()
	}
}

func (db *MyDatabase) CreateSchema() {
	if _, err := db.DB.Exec(createUsersTableQuery); err != nil {
		log.Println(err.Error())
	}

	if _, err := db.DB.Exec(createTagsTableQuery); err != nil {
		log.Println(err.Error())
	}

	if _, err := db.DB.Exec(createTasksTableQuery); err != nil {
		log.Println(err.Error())
	}

	if _, err := db.DB.Exec(createTasksTagsTableQuery); err != nil {
		log.Println(err.Error())
	}
}

func (db *MyDatabase) Purge() {
	db.DB.Exec("DELETE FROM users")
	db.DB.Exec("DELETE FROM tags")
	db.DB.Exec("DELETE FROM tasks")
	db.DB.Exec("DELETE FROM tasks_tags")
}

func (db *MyDatabase) CreateAdmin() {
	email := "ferra.main@gmail.com"
	password := utils.Hash(email, "password")
	firstName := "Franco"
	lastName := "Ferraguti"

	_, err := db.DB.Exec(`INSERT INTO users (email, password, firstName, lastName) VALUES (?, ?, ?, ?)`, email, password, firstName, lastName)
	if err != nil {
		log.Println(err.Error())
	}
}

func (db *MyDatabase) BeautifyError(err error) string {
	s := err.Error()

	if strings.Contains(s, "Duplicate entry") {
		duplicateField := strings.Split(s, "'")[3]
		return duplicateField + " already in use"
	}

	return s
}
