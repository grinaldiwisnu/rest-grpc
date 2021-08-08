package config

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5433")
	options := &pg.Options{
		User:     "user",
		Password: "password",
		Addr:     address,
		Database: "tokopedia-bms",
		PoolSize: 50,
	}
	con = pg.Connect(options)
	if con == nil {
		fmt.Println("cannot connect to postgres")
	}

	return
}
