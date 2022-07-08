package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// real project use the different host
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "my-todoapp"
)

func DB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migrate() {
	db := DB()
	db.AutoMigrate(&UserDB{})
}
