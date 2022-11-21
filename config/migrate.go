package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (s *Store) MigrateDB() error {

	if err := s.DB.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
		return err
	}

	// Run migrations
	driver, err := mysql.WithInstance(s.DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
		return err
	}

	log.Println("Database migrated")

	return nil
}
