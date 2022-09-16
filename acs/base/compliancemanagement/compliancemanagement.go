package compliancemanagement

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type ComplianceManagement struct {
	Client client.Client
}
func (a ComplianceManagement) GetRecentRuns(args map[string]interface{}) {

//  
// NOT Required [{'name': 'clusterId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'standardId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'since', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'format': 'date-time'}}]

    ok := tools.CheckFieldsValid("clusterId_string,standardId_string,since_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetRecentRunsResource()





}
func (a ComplianceManagement) GetRunStatuses(args map[string]interface{}) {

//  
// NOT Required [{'name': 'runIds', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string'}}}]

    ok := tools.CheckFieldsValid("runIds_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetRunStatusesResource()





}
