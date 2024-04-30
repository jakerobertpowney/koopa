package models

import "fmt"

type Connection struct {
	Label    string `json:"label"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func (c *Connection) Save() error {

	statement := fmt.Sprintf("INSERT INTO connections (label, username, hostname, port) VALUES ('%s', '%s', '%s', %s)", c.Label, c.Username, c.Hostname, c.Port)
	if err := save(statement); err != nil {
		return fmt.Errorf("error saving connection: %s", err)
	}

	return nil

}

func LoadConnection(key string) (*Connection, error) {

	connection := &Connection{}

	if err := get("connections", key, connection); err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err.Error())
	}

	return connection, nil

}
