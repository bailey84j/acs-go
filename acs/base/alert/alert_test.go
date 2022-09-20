package alert_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bailey84j/acs-go/acs"
)

type alertTests struct {
	Name   string
	Args   map[string]interface{}
	Result string
}

func TestListAlerts(t *testing.T) {
	// variables declaration

	password := os.Getenv("acs_password")
	username := os.Getenv("acs_username")
	//url := os.Getenv("acs_url")
	fmt.Printf("ux: %s\n", username)
	fmt.Printf("px: %s\n", password)

	//LogLevel := (*test_LogLevel)
	tableFile := "../../test_tables/alert/listalerts.json"
	State, err := ioutil.ReadFile(tableFile)
	if err != nil {
		t.Errorf("could not read test table file: %d", err)
	}

	var getTests []alertTests
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

			acs := acs.NewACSAPI("https://central-rhacs-operator.apps.rosa2.x72a.p1.openshiftapps.com")
			err = acs.Authenticate(username, password)
			if err != nil {
				if err.Error() == tt.Result {
					t.Errorf("test failed: %v", err.Error())
				}
			}

			fmt.Printf("u+p: %+v", acs.Alert.Client.Client)
			acs.Alert.ListAlerts(testParams.args)

		})
	}
}
