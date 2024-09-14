package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Map[T any] map[string]T

func (self Map[T]) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self Map[T]) Value() (driver.Value, error) {
	value, err := json.Marshal(self)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[models:map]: %s", err.Error()))
	}

	return value, nil
}

func (self *Map[T]) Scan(value any) error {
	err := json.Unmarshal(value.([]byte), self)

	if err != nil {
		return errors.New(fmt.Sprintf("[models:map]: %s", err.Error()))
	}

	return nil
}
