package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var USER_DB = "root"
var PASS_DB = "root"
var NAME_DB = "wiki"

/**
 *
 * Make connexion with database
 * Use this function before making database request
 *
 * Example of use
 *
 * db, err := ConnectToBDD()
 * if err != nil {
 *	return "what you want"
 * }
 *
 * defer db.Close()
 *
 */
func ConnectToBDD() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", USER_DB+":"+PASS_DB+"@/"+NAME_DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
