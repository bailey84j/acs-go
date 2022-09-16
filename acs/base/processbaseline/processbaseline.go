package processbaseline

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type ProcessBaseline struct {
	Client client.Client
}
// `GetProcessBaselineById` returns the single
process baseline referenced by the given ID.
func (a ProcessBaseline) GetProcessBaseline(args map[string]interface{}) {

//  
// NOT Required [{'name': 'key.deploymentId', 'description': 'The idea is for the keys to be flexible.\nOnly certain combinations of these will be supported.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key.containerName', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key.clusterId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key.namespace', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

ok := tools.CheckFieldsValid("key_deploymentId_string,key_containerName_string,key_clusterId_string,key_namespace_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// GetResource()





}
// `DeleteProcessBaselines` deletes baselines.
func (a ProcessBaseline) DeleteProcessBaselines(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'confirm', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,confirm_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// DeleteResource()





}
