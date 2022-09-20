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

	//flag.StringVar(&password, "p", "password", "Specify pass. Default is password")
	//flag.Parse()

	//
	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type checkFieldsParams struct {
				fieldTypes string
				args       map[string]interface{}
			}

			ttArgsByte, err := json.Marshal(tt.Args)
			if err != nil {
				t.Errorf("error occurred while Marshling testParams: %d", err)
			}
			var testParams checkFieldsParams
			err = json.Unmarshal(ttArgsByte, testParams)
			if err != nil {
				t.Errorf("error occurred while unmarshling testParams: %d", err)
			}
			ok := CheckFieldsValid(testParams.fieldTypes, testParams.args)
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
