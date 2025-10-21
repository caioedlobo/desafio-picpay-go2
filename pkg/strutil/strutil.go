package strutil

import (
	"fmt"
	"reflect"
	"strings"
)

type Envelope map[string]any

func ErrorEnvelope(obj any) Envelope {
	return Envelope{"error": obj}
}

func JSONStringify(json any) string {
	t := reflect.TypeOf(json)
	v := reflect.ValueOf(json)

	// If it is a pointer, get the element
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	var result []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		result = append(result, fmt.Sprintf(`"%s": "%v"`,
			field.Tag.Get("json"),
			value.Interface()))
	}
	return "{" + strings.Join(result, "\n") + "}"
}
