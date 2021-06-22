package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type PostgresConnector struct {
	Username string
	Password string
	DbName string
	DbHost string
}

func (p *PostgresConnector) GetConnection() (db *gorm.DB, err error)  {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		p.DbHost, p.Username, p.DbName, p.Password)
	return gorm.Open("postgres", dbURI)
}