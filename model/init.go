package model

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/log"
)

func init() {
	err := global.DB.AutoMigrate(
		&User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}