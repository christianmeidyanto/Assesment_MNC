package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbCOnnection interface {
	Conn() *sql.DB
}

type dbConnection struct {
	db  *sql.DB
	cfg *Config
}

func (d *dbConnection) initDb() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.cfg.DbConfig.User, d.cfg.DbConfig.Password, d.cfg.DbConfig.Host, d.cfg.DbConfig.Port, d.cfg.DbConfig.Name)
	db, err := sql.Open(d.cfg.DbConfig.Driver, dataSourceName)

	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *dbConnection) Conn() *sql.DB {
	return d.db
}

func NewDbConnection(configParam *Config) (DbCOnnection, error) {
	conn := &dbConnection{
		cfg: configParam,
	}

	err := conn.initDb()

	if err != nil {
		return nil, err
	}
	return conn, nil
}
