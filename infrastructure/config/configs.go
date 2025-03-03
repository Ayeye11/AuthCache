package config

import "fmt"

// Main Config
type ConfigAPI struct {
	APP ConfigApp
	SQL ConfigSQL
}

func LoadConfig() ConfigAPI {
	return ConfigAPI{

		ConfigApp{
			Host:     appHost,
			Port:     appPort,
			TokenKey: appTokenKey,
		},

		ConfigSQL{
			Host:     sqlHost,
			User:     sqlUser,
			Password: sqlPassword,
			DbName:   sqlDbName,
			Port:     sqlPort,
		},
	}
}

// SQL Config
type ConfigSQL struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func (c *ConfigSQL) DSN_mysql() string {
	return fmt.Sprintf("%s:%s@tcp(%s%s)/%s?allowMultiStatements=true", c.User, c.Password, c.Host, c.Port, c.DbName)
}

func (c *ConfigSQL) DSN_postgres() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.DbName, c.Port)
}

// APP Config
type ConfigApp struct {
	Host     string
	Port     string
	TokenKey string
}
