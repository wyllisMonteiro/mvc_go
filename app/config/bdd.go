package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var USER_DB 		= "root"
var PASS_DB 		= "root"
var NAME_DB 		= "wiki"

func ConnectToBDD() (*gorm.DB, error){
	db, err := gorm.Open("mysql", USER_DB + ":" + PASS_DB + "@/" + NAME_DB + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}