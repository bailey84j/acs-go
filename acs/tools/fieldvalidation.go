package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func CheckFieldsValid(fieldTypes string, args map[string]interface{}) bool {

	type FieldsAndTypes struct {
		Field string
		Type  string
	}
	var keyvalid bool
	var valuevalid bool
	var checkFields []FieldsAndTypes
	//fmt.Printf("fieldTypes: %s", fieldTypes)

	fieldsAndTypes := strings.Split(fieldTypes, ",")
	for _, fieldAndType := range fieldsAndTypes {
		splitFieldAndType := strings.Split(fieldAndType, "-")

		checkFields = append(checkFields, FieldsAndTypes{
			Field: splitFieldAndType[0],
			Type:  splitFieldAndType[1],
		})
	}

	//fmt.Printf("splitFieldAndType: %+v\n", checkFields)

	var checkargs map[string]interface{}

	argByte, err := json.Marshal(args)
	if err != nil {
		fmt.Errorf("could not read test table file: %d", err)
	}

	err = json.Unmarshal(argByte, &checkargs)
	if err != nil {
		fmt.Errorf("could not read test table file: %d", err)
	}

	for key, value := range checkargs {
		//check key valid
		keyvalid = false
		valuevalid = false
		for _, checkField := range checkFields {
			//fmt.Printf("key: %s\tcheck: %s\n", key, checkField.Field)
			if key == checkField.Field {
				keyvalid = true

				//fmt.Printf("key: %s\tvalue: %i\ttype: %s \texpectedType: %s\n", key, value, reflect.ValueOf(value).Kind().String(), checkField.Type)

				if reflect.ValueOf(value).Kind().String() == checkField.Type {

					valuevalid = true
				}
			}

		}

		//check value valid

	}

	if keyvalid && valuevalid {
		return true
	} else {
		return false
	}
}
