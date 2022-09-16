package networkgraph

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type NetworkGraph struct {
	Client client.Client
}
func (a NetworkGraph) GetNetworkGraph(clusterId string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'since', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'format': 'date-time'}}, {'name': 'includePorts', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'scope.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("query_string,since_string,includePorts_boolean,scope_query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
func (a NetworkGraph) GetExternalNetworkEntities(clusterId string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetExternalNetworkEntitiesResource()





}
func (a NetworkGraph) GetNetworkGraphConfig(args map[string]interface{}) {

//  
// NOT Required 




    // GetResource()





}
func (a NetworkGraph) DeleteExternalNetworkEntity(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteExternalNetworkEntityResource()





}
