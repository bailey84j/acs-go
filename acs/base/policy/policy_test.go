package policy_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type policyTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestListPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/listpolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.ListPolicies(testParams.args)
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
func TestQueryDryRunJobStatus(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/querydryrunjobstatus.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                jobId string `json:"jobId"`			
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

			msi, err := acs.Policy.QueryDryRunJobStatus(testParams.jobId,testParams.args)
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
func TestGetPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/getpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.GetPolicy(testParams.id,testParams.args)
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
func TestGetPolicyMitreVectors(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/getpolicymitrevectors.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.GetPolicyMitreVectors(testParams.id,testParams.args)
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
func TestGetPolicyCategories(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/getpolicycategories.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.GetPolicyCategories(testParams.args)
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


func TestCancelDryRunJob(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/canceldryrunjob.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                jobId string `json:"jobId"`			
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

			 err := acs.Policy.CancelDryRunJob(testParams.jobId,testParams.args)
                	if err.Error() != tt.Result {
				//t.Log(tt.Result)
			    t.Errorf("test failed: %v", err.Error())
			}
        
		})
	}
}
func TestDeletePolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/deletepolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			 err := acs.Policy.DeletePolicy(testParams.id,testParams.args)
                	if err.Error() != tt.Result {
				//t.Log(tt.Result)
			    t.Errorf("test failed: %v", err.Error())
			}
        
		})
	}
}
func TestDeletePolicyCategory(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/deletepolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                category string `json:"category"`			
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

			 err := acs.Policy.DeletePolicyCategory(testParams.category,testParams.args)
                	if err.Error() != tt.Result {
				//t.Log(tt.Result)
			    t.Errorf("test failed: %v", err.Error())
			}
        
		})
	}
}


func TestPatchPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/patchpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.PatchPolicy(testParams.id,testParams.args)
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
func TestEnableDisablePolicyNotification(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/enabledisablepolicynotification.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                policyId string `json:"policyId"`			
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

			msi, err := acs.Policy.EnableDisablePolicyNotification(testParams.policyId,testParams.args)
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


func TestPostPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/postpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.PostPolicy(testParams.args)
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
func TestDryRunPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/dryrunpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.DryRunPolicy(testParams.args)
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
func TestSubmitDryRunPolicyJob(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/submitdryrunpolicyjob.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.SubmitDryRunPolicyJob(testParams.args)
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
func TestExportPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/exportpolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.ExportPolicies(testParams.args)
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
func TestPolicyFromSearch(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/policyfromsearch.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.PolicyFromSearch(testParams.args)
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
func TestImportPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/importpolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.ImportPolicies(testParams.args)
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
func TestReassessPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/reassesspolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.ReassessPolicies(testParams.args)
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


func TestPutPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/policy/putpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
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

			msi, err := acs.Policy.PutPolicy(testParams.id,testParams.args)
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

	tableFile := "../../test_tables/policy/renamepolicycategory.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []policyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                oldCategory string `json:"oldCategory"`			
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

			msi, err := acs.Policy.RenamePolicyCategory(testParams.oldCategory,testParams.args)
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


