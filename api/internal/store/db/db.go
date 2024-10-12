package db

import (
	// "database/sql"
	// "fmt"
	// "goride/internal/store"
	// "os"
	// "fmt"
	"log"

	_ "github.com/shaxbee/go-spatialite"
	// "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {

	// connStr := "user=myuser dbname=mydb password=mypassword host=postgis-db port=5432 sslmode=disable"
	connStr := "user=myuser dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	// fmt.Print(tx)
	// if err != nil {
	// 	log.Fatalf("Failed to create PostGIS extension: %v", err)
	// }

	db.Exec(`
		CREATE TABLE IF NOT EXISTS routes (
			id SERIAL PRIMARY KEY,
			geometry GEOMETRY(LineString, 4326)
		);
	`)

	// err = db.AutoMigrate(&store.Route{}, &store.User{})

	// if err != nil {
	// 	log.Fatalf("Failed to create table: %v", err)
	// }
	return db
}
