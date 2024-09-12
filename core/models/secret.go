package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/aacebo/agent.net/core/utils"
)

var secret = utils.GetEnv(
	"SECRET",
	"/tdpk50HHcQFvSFvF08MFg==",
)

type Secret string

func (self Secret) String() string {
	return string(self)
}

func (self Secret) Equals(s string) bool {
	return string(self) == s
}

func (self Secret) Value() (driver.Value, error) {
	value, err := utils.AESEncrypt([]byte(self), secret)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[models:secret]: %s", err.Error()))
	}

	return string(value), nil
}

func (self *Secret) Scan(value any) error {
	decrypted, err := utils.AESDecrypt([]byte(value.(string)), secret)

	if err != nil {
		return errors.New(fmt.Sprintf("[models:secret]: %s", err.Error()))
	}

	*self = Secret(decrypted)
	return nil
}
