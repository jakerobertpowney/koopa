package models

import (
	"fmt"
)

type Setting struct {
	Value string
}

func (s *Setting) Save(key string) error {

	if err := boltSave("settings", key, s); err != nil {
		return fmt.Errorf("failed to save settings: %v", err)
	}

	return nil

}

func LoadSetting(key string) (*Setting, error) {

	setting := &Setting{}

	if err := boltGet("settings", key, setting); err != nil {
		return nil, fmt.Errorf("failed to get setting: %v", err.Error())
	}

	return setting, nil

}
