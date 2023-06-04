package env

import (
	"encoding"
	"fmt"
	"log"
	"reflect"
	"sort"

	"github.com/Dagosu/BookingApp/gohelpers/cases"
)

type envMap map[string]string

const separator = "_"

var skipFields = map[string]bool{
	"CreatedAt": true,
	"UpdatedAt": true,
}

func (e *envMap) add(key, value string) {
	(*e)[cases.ToDottedUpperSnakeCase(key)] = value
}

// inspired from https://gist.github.com/hvoecking/10772475
func (e *envMap) interpret(original reflect.Value, prefix string) error {
	// skip private fields
	if !original.CanInterface() {
		return nil
	}

	if original.IsZero() {
		return nil
	}

	// check custom marshalling first
	m, ok := original.Interface().(encoding.TextMarshaler)
	if ok {
		value, err := m.MarshalText()
		if err != nil {
			return err
		}

		e.add(prefix, string(value))

		return nil
	}

	// default env values
	kind := original.Kind()
	switch kind {
	case reflect.Ptr:
		valueInside := original.Elem()
		if !valueInside.IsValid() {
			return nil
		}

		return e.interpret(valueInside, prefix)

	case reflect.Interface:
		valueInside := original.Elem()

		return e.interpret(valueInside, prefix)

	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			fieldName := original.Type().Field(i).Name

			if skipFields[fieldName] {
				continue
			}

			err := e.interpret(original.Field(i), prefix+separator+fieldName)
			if err != nil {
				return err
			}
		}

	case reflect.Map:
		for _, key := range original.MapKeys() {
			err := e.interpret(original.MapIndex(key), prefix+separator+key.String())
			if err != nil {
				return err
			}
		}

	case reflect.String:
		e.add(prefix, original.String())

	case reflect.Bool:
		e.add(prefix, fmt.Sprintf("%t", original.Bool()))

	case reflect.Int, reflect.Int32, reflect.Int64:
		e.add(prefix, fmt.Sprintf("%d", original.Int()))

	default:
		log.Printf("error: unknown env struct reflect original value of kind '%+v', of type '%T': %#v", kind, original, original)
	}

	return nil
}

func ToSlice(m envMap) []string {
	ret := []string{}

	for k, v := range m {
		ret = append(ret, k+"="+v)
	}

	sort.Strings(ret)

	return ret
}

func FromStructRecursive(s interface{}, prefix string) (envMap, error) {
	ret := envMap{}

	original := reflect.ValueOf(s)

	err := ret.interpret(original, prefix)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
