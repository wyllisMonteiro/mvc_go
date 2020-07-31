package models

import (
	"fmt"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type config struct {
	UserDB string `env:"USERDB" envDefault:"root"`
	PassDB string `env:"PASSDB" envDefault:"root"`
	NameDB string `env:"NAMEDB" envDefault:"wiki"`
}

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
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	db, err := gorm.Open("mysql", cfg.UserDB+":"+cfg.PassDB+"@/"+cfg.NameDB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
