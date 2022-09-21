package rbac_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type rbacTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestListRoleBindings(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/listrolebindings.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.ListRoleBindings(testParams.args)
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
func TestGetRoleBinding(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/getrolebinding.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.GetRoleBinding(testParams.id,testParams.args)
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
func TestListRoles(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/listroles.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.ListRoles(testParams.args)
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
func TestGetRole(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/getrole.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.GetRole(testParams.id,testParams.args)
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
func TestGetSubject(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/getsubject.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.GetSubject(testParams.id,testParams.args)
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
func TestListSubjects(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/rbac/listsubjects.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []rbacTests
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

			msi, err := acs.Rbac.ListSubjects(testParams.args)
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


