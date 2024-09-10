package schemas

import (
	"errors"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Schemas map[string]*jsonschema.Schema

func (self Schemas) Validate(name string, data any) error {
	schema, exists := self[name]

	if !exists {
		return errors.New(fmt.Sprintf(`schema "%s" not found`, name))
	}

	return schema.Validate(data)
}
