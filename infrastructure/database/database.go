package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Ayeye11/se-thr/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseSQL struct {
	engineName string
	config     config.ConfigSQL
	db         *gorm.DB
}

var initDriverFunc = map[string]func(c config.ConfigSQL) gorm.Dialector{
	"mysql": func(c config.ConfigSQL) gorm.Dialector { return mysql.Open(c.DSN_mysql()) },
	//"postgres": func(c config.ConfigSQL) gorm.Dialector { return postgres.Open(c.DSN_postgres()) },
}

func InitSQL(engineName string, config config.ConfigSQL, attempts ...int) (*databaseSQL, error) {
	driverFunc, exists := initDriverFunc[engineName]
	if !exists {
		return nil, fmt.Errorf("'%s' is not available", engineName)
	}

	trys := 1
	if len(attempts) > 0 && attempts[0] >= 1 && attempts[0] <= 10 {
		trys = attempts[0]
	}

	conn := func(ctx context.Context) (*gorm.DB, error) {
		dbCh := make(chan *gorm.DB, 1)
		errCh := make(chan error, 1)

		go func() {
			db, err := gorm.Open(driverFunc(config), &gorm.Config{})
			if err != nil {
				errCh <- err
				return
			}

			dbCh <- db
		}()

		select {
		case err := <-errCh:
			return nil, err

		case db := <-dbCh:
			return db, nil

		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to the database: timeout")
		}
	}

	for i := 1; i <= trys; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		db, err := conn(ctx)
		if err == nil {
			return &databaseSQL{engineName, config, db}, nil
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

func (sql *databaseSQL) GetDB() *gorm.DB {
	return sql.db
}
