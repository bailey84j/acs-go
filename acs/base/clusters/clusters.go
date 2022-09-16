package clusters

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Clusters struct {
	Client client.Client
}
func (a Clusters) GetClusterDefaultValues(args map[string]interface{}) {

//  
// NOT Required 




    // GetClusterDefaultValuesResource()





}
func (a Clusters) GetClusters(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
// GetKernelSupportAvailable is deprecated in favor of GetClusterDefaultValues.
func (a Clusters) GetKernelSupportAvailable(args map[string]interface{}) {

//  
// NOT Required 




    // GetKernelSupportAvailableResource()





}
func (a Clusters) GetCluster(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetClusterResource()





}
func (a Clusters) DeleteCluster(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteClusterResource()





}
