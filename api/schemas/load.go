package schemas

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func Load() (Schemas, error) {
	schemas := Schemas{}
	errs := []error{}
	compiler := jsonschema.NewCompiler()
	err := filepath.Walk("./api/schemas", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}

		name := path
		name = name[7 : len(name)-5]
		schema, err := compiler.Compile(path)

		if err != nil {
			errs = append(errs, err)
			return nil
		}

		schemas[name] = schema
		return nil
	})

	if err != nil {
		return nil, err
	}

	return schemas, errors.Join(errs...)
}
