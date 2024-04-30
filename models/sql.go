package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	db *sqlx.DB
)

func sqlDB() (*sqlx.DB, error) {

	activeDatabase, err := LoadSetting("ACTIVE_DATABASE")
	if err != nil {
		log.Fatalf("error loading setting: %v", err)
	}

	databaseConfig, err := LoadDatabase(activeDatabase.Value)
	if err != nil {
		log.Fatalf("error loading database: %v", err)
	}

	switch databaseConfig.Type {
	case "local":
		dbName := fmt.Sprintf("./%s.db", databaseConfig.Database)
		db, err = sqlx.Connect("sqlite", dbName)
	case "MySQL (Remote)":
		dbName := fmt.Sprintf("%s@tcp(%s)/%s", databaseConfig.User, databaseConfig.Hostname, databaseConfig.Database)
		db, err = sqlx.Connect("mysql", dbName)
	}

	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return db, nil

}

func save(statement string) error {

	db, err := sqlDB()
	if err != nil {
		return fmt.Errorf("error initilasing database driver: %v", err)
	}

	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(statement)
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error saving database record: %v", err)
	}

	return nil

}

func get(table string, key string, i interface{}) error {

	db, err := sqlDB()
	if err != nil {
		return fmt.Errorf("error initilasing database driver: %v", err)
	}

	defer db.Close()

	statement := fmt.Sprintf("SELECT * FROM %s WHERE label = '%s'", table, key)
	if err := db.Get(i, statement); err != nil {
		return fmt.Errorf("error selecting data from %s: %v", table, err)
	}

	return nil

}
