package sql

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationPath = "file://infrastructure/sql/migrations"

var migrateDriverFunc = map[string]func(instance *sql.DB) (database.Driver, error){
	"mysql": func(instance *sql.DB) (database.Driver, error) {
		return mysql.WithInstance(instance, &mysql.Config{})
	},
	//"postgres": func(instance *sql.DB) (database.Driver, error) {
	//	return postgres.WithInstance(instance, &postgres.Config{})
	//},
}

func (sql *databaseSQL) GetMigrator() (*migrate.Migrate, error) {
	driverFunc, exists := migrateDriverFunc[sql.engineName]
	if !exists {
		return nil, fmt.Errorf("'%s' is not available", sql.engineName)
	}

	sqlDB, err := sql.GetInstance()
	if err != nil {
		return nil, err
	}

	driver, err := driverFunc(sqlDB)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(migrationPath, sql.config.DbName, driver)
}
