package model

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/log"
)

func init() {
	err := global.DB.AutoMigrate(
		&User{},
		&Post{},
		&Comment{},
		&Product{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
