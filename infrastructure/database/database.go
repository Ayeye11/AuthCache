package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ayeye11/se-thr/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseSQL struct {
	config config.ConfigSQL
	db     *gorm.DB
}

var initDriverFunc = map[string]func(c config.ConfigSQL) gorm.Dialector{
	"mysql":    func(c config.ConfigSQL) gorm.Dialector { return mysql.Open(c.DSN_mysql()) },
	"postgres": func(c config.ConfigSQL) gorm.Dialector { return postgres.Open(c.DSN_postgres()) },
}

func InitSQL(engineName string, config config.ConfigSQL, attempts int) (*databaseSQL, error) {
	driverFunc, exists := initDriverFunc[engineName]
	if !exists {
		return nil, fmt.Errorf("'%s' is not available", engineName)
	}

	trys := attempts
	if attempts < 1 || attempts > 10 {
		trys = 1
	}

	for i := 1; i <= trys; i++ {
		db, err := gorm.Open(driverFunc(config), &gorm.Config{})
		if err == nil {
			return &databaseSQL{config, db}, nil
		}

		if trys == 1 {
			break
		}

		fmt.Printf("failed to connect to the database, attempt %d/%d...\n", i, trys)
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to the database")
}

// Methods:
func (sql *databaseSQL) GetInstance() (*sql.DB, error) {
	return sql.db.DB()
}

func (sql *databaseSQL) Close() error {
	ins, err := sql.db.DB()
	if err != nil {
		return err
	}

	return ins.Close()
}
