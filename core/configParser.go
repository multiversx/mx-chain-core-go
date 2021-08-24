package core

import (
	"fmt"
	"math"
	"reflect"
)

const defaultUint32Value = math.MaxUint32 - 5
const defaultStringValue = "default-value"

// LoadTomlFileWithDefaultChecks will load a toml file to the destination, applying checks so all the fields are required
func LoadTomlFileWithDefaultChecks(dest interface{}, relativePath string) error {
	populateStructWithDefaultValues(dest)

	err := LoadTomlFile(dest, relativePath)
	if err != nil {
		return err
	}

	return searchDefaultValues(dest)
}

func searchDefaultValues(object interface{}) error {
	var err error
	objectType, objectValue := getReflectValueAndType(object)

	for i := 0; i < objectType.NumField(); i++ {
		v := objectValue.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			err = searchDefaultValues(v)
			if err != nil {
				return err
			}
		case reflect.Slice:
			// TODO: add support for slices
		default:
			err = checkDefaultForType(v, objectType, i)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkDefaultForType(object reflect.Value, bigObjectType reflect.Type, index int) error {
	switch object.Interface().(type) {
	case uint32:
		if object.Interface().(uint32) == defaultUint32Value {
			return fmt.Errorf("config value for field %s not set", bigObjectType.Field(index).Name)
		}
	case string:
		if object.Interface().(string) == defaultStringValue {
			return fmt.Errorf("config value for field %s not set", bigObjectType.Field(index).Name)
		}
	}

	return nil
}

func populateStructWithDefaultValues(object interface{}) {
	objectType, objectValue := getReflectValueAndType(object)

	num := objectType.NumField()
	for i := 0; i < num; i++ {
		f := objectValue.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			populateStructWithDefaultValues(f)
			continue
		case reflect.Slice:
			// TODO: add support for slices
		default:
			switch f.Interface().(type) {
			case uint32:
				f.SetUint(defaultUint32Value)
			case string:
				f.SetString(defaultStringValue) //Set a value to this field
			}
		}
	}
}

func getReflectValueAndType(object interface{}) (reflect.Type, reflect.Value) {
	var objectValue reflect.Value
	var objectType reflect.Type

	switch object.(type) {
	case reflect.Value:
		objectValue = object.(reflect.Value)
		objectType = objectValue.Type()
	default:
		objectValue = reflect.ValueOf(object)
		objectType = reflect.TypeOf(object)
	}

	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
	}
	if objectValue.Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}

	return objectType, objectValue
}
