package networkpolicy

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type NetworkPolicy struct {
	Client client.Client
}
func (a NetworkPolicy) GetNetworkPolicies(args map[string]interface{}) {

//  
// NOT Required [{'name': 'clusterId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'deploymentQuery', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'namespace', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("clusterId_string,deploymentQuery_string,namespace_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetNetworkPoliciesResource()





}
func (a NetworkPolicy) GetAllowedPeersFromCurrentPolicyForDeployment(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetAllowedPeersFromCurrentPolicyForDeploymentResource()





}
func (a NetworkPolicy) GetDiffFlowsBetweenPolicyAndBaselineForDeployment(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetDiffFlowsBetweenPolicyAndBaselineForDeploymentResource()





}
func (a NetworkPolicy) GetNetworkGraph(clusterId string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'includePorts', 'description': 'If set to true, include port-level information in the network policy graph.', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'scope.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("query_string,includePorts_boolean,scope_query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetNetworkGraphResource()





}
func (a NetworkPolicy) GenerateNetworkPolicies(clusterId string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'deleteExisting', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'enum': ['UNKNOWN', 'NONE', 'GENERATED_ONLY', 'ALL'], 'default': 'UNKNOWN'}}, {'name': 'networkDataSince', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'format': 'date-time'}}, {'name': 'includePorts', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("query_string,deleteExisting_string,networkDataSince_string,includePorts_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GenerateNetworkPoliciesResource()





}
func (a NetworkPolicy) GetNetworkGraphEpoch(args map[string]interface{}) {

//  
// NOT Required [{'name': 'clusterId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("clusterId_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetNetworkGraphEpochResource()





}
func (a NetworkPolicy) GetUndoModificationForDeployment(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetUndoModificationForDeploymentResource()





}
func (a NetworkPolicy) GetUndoModification(clusterId string,args map[string]interface{}) {

//  
// NOT Required []




    // GetUndoModificationResource()





}
func (a NetworkPolicy) GetDiffFlowsFromUndoModificationForDeployment(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetDiffFlowsFromUndoModificationForDeploymentResource()





}
func (a NetworkPolicy) GetNetworkPolicy(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
