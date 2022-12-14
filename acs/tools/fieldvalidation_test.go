package tools

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"
)

type toolTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestCheckFieldsValid(t *testing.T) {
	// variables declaration

	//LogLevel := (*test_LogLevel)
	tableFile := "./test_tables/fieldvalidation.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []toolTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	//
	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type checkFieldsParams struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
			}
			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := checkFieldsParams{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			//fmt.Printf("testParams.fieldTypes: %s", testParams.fieldTypes)
			//fmt.Printf("testParams.args: %v", testParams.args)

			ok := CheckFieldsValid(testParams.fieldTypes, testParams.args)

			//fmt.Printf("Result: %t", ok)
			if strconv.FormatBool(ok) == tt.Result {
				t.Log("Success")
			} else {
				if strconv.FormatBool(ok) != tt.Result {
					t.Log("Fail")
				}
			}
		})
	}
}
