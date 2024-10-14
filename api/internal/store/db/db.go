package db

import (
	"log"

	_ "github.com/shaxbee/go-spatialite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func InitDatabase() *gorm.DB {

	// connStr := "user=myuser dbname=mydb password=mypassword host=postgis-db port=5432 sslmode=disable"
	// connStr := "user=myuser dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable host=localhost TimeZone=Europe/Stockholm",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	// fmt.Print(tx)
	// if err != nil {
	// 	log.Fatalf("Failed to create PostGIS extension: %v", err)
	// }

	// db.AutoMigrate(&store.Route{})

	// db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS routes (
	// 		id SERIAL PRIMARY KEY,
	// 		geometry GEOMETRY(LineString, 4326)
	// 	);
	// `)

	// err = db.AutoMigrate(&store.Route{}, &store.User{})

	// if err != nil {
	// 	log.Fatalf("Failed to create table: %v", err)
	// }
	return db
}

// func AddRoute(db *sql.DB, route types.Route) error {
// 	// Prepare the SQL insert statement
// 	stmt, err := db.Prepare("INSERT INTO routes(geometry) VALUES (ST_GeomFromGeoJSON($1))")
// 	if err != nil {
// 		return fmt.Errorf("failed to prepare insert statement: %v", err)
// 	}
// 	defer stmt.Close()
//
// 	// Insert each route's geometry into the database
// 	geoJSON, err := json.Marshal(route.Geometry)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal geometry to GeoJSON: %v", err)
// 	}
//
// 	_, err = stmt.Exec(string(geoJSON))
// 	if err != nil {
// 		return fmt.Errorf("failed to insert route geometry: %v", err)
// 	}
//
// 	return nil
// }
//
