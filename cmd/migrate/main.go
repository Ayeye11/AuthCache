package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ayeye11/AuthCache/config"
	"github.com/Ayeye11/AuthCache/infrastructure/sql"
	"github.com/golang-migrate/migrate/v4"
)

func logf(err error) {
	if err != nil {
		log.Fatalf("fatal: %v\n", err)
	}
}

func main() {
	cfg := config.LoadConfig().SQL

	sql, err := sql.InitSQL("mysql", cfg)
	logf(err)

	m, err := sql.GetMigrator()
	logf(err)

	arg := os.Args[1:]

	if err := runMigration(m, arg...); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatalf("fatal: %v\n", err)
		}

		log.Println("No migrations to apply")
		return
	}

	log.Println("migrations applied successfully")
}

func runMigration(m *migrate.Migrate, arg ...string) error {
	if len(arg) > 0 && arg[0] != "down" && arg[0] != "up" {
		return fmt.Errorf("invalid command: use 'up' or 'down'")
	}

	if len(arg) > 0 && arg[0] == "down" {
		return m.Down()
	}

	return m.Up()
}
