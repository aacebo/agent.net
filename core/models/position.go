package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (self Position) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self Position) Value() (driver.Value, error) {
	value, err := json.Marshal(self)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[models:position]: %s", err.Error()))
	}

	return value, nil
}

func (self *Position) Scan(value any) error {
	err := json.Unmarshal(value.([]byte), self)

	if err != nil {
		return errors.New(fmt.Sprintf("[models:position]: %s", err.Error()))
	}

	return nil
}
