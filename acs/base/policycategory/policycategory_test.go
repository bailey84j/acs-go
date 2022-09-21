package policycategory_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type policycategoryTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestGetPolicyCategories(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policycategory/getpolicycategories.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policycategoryTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                			
            }

			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := params{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			acs := acs.NewACSAPI(url)
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			msi, err := acs.PolicyCategory.GetPolicyCategories(testParams.args)
        			if msi != nil {
				jsonbyte, err := json.Marshal(msi)
				if err != nil {
					t.Errorf("error parsing JSON string - %s", err)
				}

				jsonstring := string(jsonbyte)
				t.Log(jsonstring)

				if jsonstring != tt.Result {
					t.Errorf("Expected: " + tt.Result + "\nGot: " + jsonstring)
				}
				//fmt.Println(j)
			} else {
				//fmt.Println(err)
				if err.Error() != tt.Result {
					//t.Log(tt.Result)
					t.Errorf("test failed: %v", err.Error())
				}
			}
        
		})
	}
}
func TestGetPolicyCategory(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policycategory/getpolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policycategoryTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                id string `json:"id"`			
            }

			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := params{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			acs := acs.NewACSAPI(url)
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			msi, err := acs.PolicyCategory.GetPolicyCategory(testParams.id,testParams.args)
        			if msi != nil {
				jsonbyte, err := json.Marshal(msi)
				if err != nil {
					t.Errorf("error parsing JSON string - %s", err)
				}

				jsonstring := string(jsonbyte)
				t.Log(jsonstring)

				if jsonstring != tt.Result {
					t.Errorf("Expected: " + tt.Result + "\nGot: " + jsonstring)
				}
				//fmt.Println(j)
			} else {
				//fmt.Println(err)
				if err.Error() != tt.Result {
					//t.Log(tt.Result)
					t.Errorf("test failed: %v", err.Error())
				}
			}
        
		})
	}
}


func TestDeletePolicyCategory(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policycategory/deletepolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policycategoryTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                id string `json:"id"`			
            }

			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := params{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			acs := acs.NewACSAPI(url)
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			 err := acs.PolicyCategory.DeletePolicyCategory(testParams.id,testParams.args)
                	if err.Error() != tt.Result {
				//t.Log(tt.Result)
			    t.Errorf("test failed: %v", err.Error())
			}
        
		})
	}
}


func TestPostPolicyCategory(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policycategory/postpolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policycategoryTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                			
            }

			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := params{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			acs := acs.NewACSAPI(url)
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			msi, err := acs.PolicyCategory.PostPolicyCategory(testParams.args)
        			if msi != nil {
				jsonbyte, err := json.Marshal(msi)
				if err != nil {
					t.Errorf("error parsing JSON string - %s", err)
				}

				jsonstring := string(jsonbyte)
				t.Log(jsonstring)

				if jsonstring != tt.Result {
					t.Errorf("Expected: " + tt.Result + "\nGot: " + jsonstring)
				}
				//fmt.Println(j)
			} else {
				//fmt.Println(err)
				if err.Error() != tt.Result {
					//t.Log(tt.Result)
					t.Errorf("test failed: %v", err.Error())
				}
			}
        
		})
	}
}


func TestRenamePolicyCategory(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policycategory/renamepolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policycategoryTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                			
            }

			var args map[string]interface{}
			json.Unmarshal([]byte(tt.Args["args"].(string)), &args)

			testParams := params{
				fieldTypes: tt.Args["fieldTypes"].(string),
				args:       args,
			}

			acs := acs.NewACSAPI(url)
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			msi, err := acs.PolicyCategory.RenamePolicyCategory(testParams.args)
        			if msi != nil {
				jsonbyte, err := json.Marshal(msi)
				if err != nil {
					t.Errorf("error parsing JSON string - %s", err)
				}

				jsonstring := string(jsonbyte)
				t.Log(jsonstring)

				if jsonstring != tt.Result {
					t.Errorf("Expected: " + tt.Result + "\nGot: " + jsonstring)
				}
				//fmt.Println(j)
			} else {
				//fmt.Println(err)
				if err.Error() != tt.Result {
					//t.Log(tt.Result)
					t.Errorf("test failed: %v", err.Error())
				}
			}
        
		})
	}
}


