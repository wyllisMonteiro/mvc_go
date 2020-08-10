package models

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

// Make connexion with database
// Use this function before making database request
// Example of use
// db, err := ConnectToBDD()
// if err != nil {
//	return "what you want"
// }
// defer db.Close()
func ConnectToBDD() (*gorm.DB, error) {
	userDb, err := goDotEnvVariable("USERDB")
	if err != nil {
		userDb = "root"
	}

	passDb, err := goDotEnvVariable("PASSDB")
	if err != nil {
		passDb = "root"
	}

	nameDb, err := goDotEnvVariable("NAMEDB")
	if err != nil {
		nameDb = "wiki"
	}

	db, err := gorm.Open("mysql", userDb+":"+passDb+"@(db)/"+nameDb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
