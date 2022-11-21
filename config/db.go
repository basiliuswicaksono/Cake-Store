package config

import (
	"database/sql"
	"log"
)

type Store struct {
	DB *sql.DB
}

func ConnectDB() (Store, error) {
	db, err := sql.Open("mysql", "user1:mypassword@tcp(db-mariadb:3306)/store?parseTime=true")
	// defer db.Close()

	if err != nil {
		return Store{}, err
	}

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)

	return Store{
		DB: db,
	}, nil
}
