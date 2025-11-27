package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	*sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to MySQL database") // ✅ 日志更新
	return &Database{db}, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}
