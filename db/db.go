package db

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/hakumizuki/gin-template/config"
)

var db *pg.DB

func Init() {
	c := config.GetConfig()

	db = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf(":%v", c.GetString("db.postgres_addr")),
		User:     c.GetString("db.postgres_user"),
		Password: c.GetString("db.postgres_password"),
		Database: c.GetString("db.postgres_database"),
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
}

func GetDB() *pg.DB {
	return db
}
