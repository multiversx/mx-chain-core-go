package block

import (
	bytes "bytes"
	"encoding/json"
	"fmt"
	reflect "reflect"
	strings "strings"

	"github.com/multiversx/mx-chain-core-go/data"
)

func PrettifyHeaderHandler(hh data.HeaderHandler) string {
	jsonBytes, _ := json.Marshal(hh)

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonBytes, "", "  "); err != nil {
		return string(jsonBytes)
	}

	return prettyJSON.String()
}

func prettifyValue(val reflect.Value, typ reflect.Type) interface{} {
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
        if !val.IsValid() {
            return nil
        }
        typ = val.Type()
    }

    if val.Kind() == reflect.Struct {
        out := make(map[string]interface{})
        for i := 0; i < val.NumField(); i++ {
            field := val.Field(i)
            fieldType := typ.Field(i)
            name := fieldType.Tag.Get("json")
            if name == "" {
                name = fieldType.Name
            } else {
                name = strings.Split(name, ",")[0]
            }

            if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Uint8 {
                out[name] = fmt.Sprintf("%x", field.Bytes()) //string(field.Bytes())
            } else {
                out[name] = prettifyValue(field, field.Type())
            }
        }
        return out
    }

    return val.Interface()
}

func PrettifyStruct(x interface{}) (string, error) {
    val := reflect.ValueOf(x)
    result := prettifyValue(val, val.Type())

    jsonBytes, err := json.Marshal(result)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

func PrettifyHeader(hdr interface{}) {
	out := make(map[string]interface{})
	t := reflect.TypeOf(hdr)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	v := reflect.ValueOf(hdr)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		out[field.Name] = value.Interface()

	}
}