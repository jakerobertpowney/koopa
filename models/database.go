package models

import (
	"fmt"
)

type Database struct {
	Type     string
	Database string
	User     string
	Password string
	Hostname string
	Port     string
}

func (d *Database) Save(key string) error {

	if err := boltSave("databases", key, d); err != nil {
		return fmt.Errorf("failed to save database: %v", err.Error())
	}

	return nil

}

func LoadDatabase(key string) (*Database, error) {

	database := &Database{}

	if err := boltGet("databases", key, database); err != nil {
		return nil, fmt.Errorf("failed to get database: %v", err.Error())
	}

	return database, nil

}

func LoadAllDatabases() ([]Database, error) {

	databases := make([]Database, 0)

	if err := boltAll("databases", &databases); err != nil {
		return nil, fmt.Errorf("failed to get databases: %v", err.Error())
	}

	return databases, nil
}
