package networkpolicy_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type networkpolicyTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestGetNetworkPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getnetworkpolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetNetworkPolicies(testParams.args)
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
func TestGetAllowedPeersFromCurrentPolicyForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getallowedpeersfromcurrentpolicyfordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetAllowedPeersFromCurrentPolicyForDeployment(testParams.id,testParams.args)
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
func TestGetDiffFlowsBetweenPolicyAndBaselineForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getdiffflowsbetweenpolicyandbaselinefordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetDiffFlowsBetweenPolicyAndBaselineForDeployment(testParams.id,testParams.args)
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
func TestGetNetworkGraph(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getnetworkgraph.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.GetNetworkGraph(testParams.clusterId,testParams.args)
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
func TestGenerateNetworkPolicies(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/generatenetworkpolicies.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.GenerateNetworkPolicies(testParams.clusterId,testParams.args)
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
func TestGetNetworkGraphEpoch(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getnetworkgraphepoch.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetNetworkGraphEpoch(testParams.args)
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
func TestGetUndoModificationForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getundomodificationfordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetUndoModificationForDeployment(testParams.id,testParams.args)
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
func TestGetUndoModification(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getundomodification.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.GetUndoModification(testParams.clusterId,testParams.args)
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
func TestGetDiffFlowsFromUndoModificationForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getdiffflowsfromundomodificationfordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetDiffFlowsFromUndoModificationForDeployment(testParams.id,testParams.args)
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
func TestGetNetworkPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getnetworkpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
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

			msi, err := acs.NetworkPolicy.GetNetworkPolicy(testParams.id,testParams.args)
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


func TestApplyNetworkPolicyYamlForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/applynetworkpolicyyamlfordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                deploymentId string `json:"deploymentId"`			
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

			msi, err := acs.NetworkPolicy.ApplyNetworkPolicyYamlForDeployment(testParams.deploymentId,testParams.args)
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
func TestApplyNetworkPolicy(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/applynetworkpolicy.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.ApplyNetworkPolicy(testParams.clusterId,testParams.args)
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
func TestGetBaselineGeneratedNetworkPolicyForDeployment(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/getbaselinegeneratednetworkpolicyfordeployment.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                deploymentId string `json:"deploymentId"`			
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

			msi, err := acs.NetworkPolicy.GetBaselineGeneratedNetworkPolicyForDeployment(testParams.deploymentId,testParams.args)
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
func TestSimulateNetworkGraph(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/simulatenetworkgraph.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.SimulateNetworkGraph(testParams.clusterId,testParams.args)
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
func TestSendNetworkPolicyYAML(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	url := os.Getenv("acs_url")

	tableFile := "../../test_tables/networkpolicy/sendnetworkpolicyyaml.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []networkpolicyTests
	err = json.Unmarshal(State, &getTests)
	if err != nil {
		t.Errorf("error occurred while unmarshling Tests: %d", err)
	}

	for _, tt := range getTests {
		t.Run(tt.Name, func(t *testing.T) {

			type params struct {
				fieldTypes string                 `json:"fieldTypes"`
				args       map[string]interface{} `json:"args"`
                clusterId string `json:"clusterId"`			
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

			msi, err := acs.NetworkPolicy.SendNetworkPolicyYAML(testParams.clusterId,testParams.args)
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


