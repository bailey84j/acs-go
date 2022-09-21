package signatureintegration_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type signatureintegrationTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestListSignatureIntegrations(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/signatureintegration/listsignatureintegrations.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []signatureintegrationTests
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

			msi, err := acs.SignatureIntegration.ListSignatureIntegrations(testParams.args)
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
func TestGetSignatureIntegration(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/signatureintegration/getsignatureintegration.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []signatureintegrationTests
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

			msi, err := acs.SignatureIntegration.GetSignatureIntegration(testParams.id,testParams.args)
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


func TestDeleteSignatureIntegration(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/signatureintegration/deletesignatureintegration.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []signatureintegrationTests
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

			 err := acs.SignatureIntegration.DeleteSignatureIntegration(testParams.id,testParams.args)
                	if err.Error() != tt.Result {
				//t.Log(tt.Result)
			    t.Errorf("test failed: %v", err.Error())
			}
        
		})
	}
}


func TestPostSignatureIntegration(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/signatureintegration/postsignatureintegration.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []signatureintegrationTests
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

			msi, err := acs.SignatureIntegration.PostSignatureIntegration(testParams.args)
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


func TestPutSignatureIntegration(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/signatureintegration/putsignatureintegration.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []signatureintegrationTests
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

			msi, err := acs.SignatureIntegration.PutSignatureIntegration(testParams.id,testParams.args)
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


