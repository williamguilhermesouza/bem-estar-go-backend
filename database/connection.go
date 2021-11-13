package database

import (
	"log"

	//"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectToDb() *gorm.DB {
	// ********* LOCAL SQLITE DB ***********
	// db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("error opening db")
	// }

	// ********* REMOTE POSTGRES DB ***********
	dsn := "host=ec2-54-147-207-184.compute-1.amazonaws.com user=ishyxjgsidpkdg password=be8dc6619bc3ab0a9f451581a3f9b13adf40885ea5718182c7c2a3e71e53caf7 dbname=dbapl0o3ribrna port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		log.Fatal("error opening db")
	}

	return db
}
