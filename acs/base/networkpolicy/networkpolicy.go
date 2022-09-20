package networkpolicy

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type NetworkPolicy struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "NetworkPolicyService_GetNetworkPolicies", "parameters": [{"in": "query", "name": "clusterId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "deploymentQuery", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "namespace", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkPoliciesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetAllowedPeersFromCurrentPolicyForDeployment", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAllowedPeersFromCurrentPolicyForDeploymentResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetDiffFlowsBetweenPolicyAndBaselineForDeployment", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetDiffFlowsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetNetworkGraph", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"description": "If set to true, include port-level information in the network policy graph.", "in": "query", "name": "includePorts", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "scope.query", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkGraph"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GenerateNetworkPolicies", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "deleteExisting", "required": false, "schema": {"default": "UNKNOWN", "enum": ["UNKNOWN", "NONE", "GENERATED_ONLY", "ALL"], "type": "string"}}, {"in": "query", "name": "networkDataSince", "required": false, "schema": {"format": "date-time", "type": "string"}}, {"in": "query", "name": "includePorts", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GenerateNetworkPoliciesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetNetworkGraphEpoch", "parameters": [{"in": "query", "name": "clusterId", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkGraphEpoch"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetUndoModificationForDeployment", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetUndoModificationForDeploymentResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetUndoModification", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetUndoModificationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetDiffFlowsFromUndoModificationForDeployment", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetDiffFlowsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetNetworkPolicy", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkPolicy"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}], "method": "get"}

        
func (a NetworkPolicy) GetNetworkPolicies(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("clusterId_string,deploymentQuery_string,namespace_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetAllowedPeersFromCurrentPolicyForDeployment(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/allowedpeers/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetDiffFlowsBetweenPolicyAndBaselineForDeployment(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/baselinecomparison/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetNetworkGraph(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,includePorts_boolean,scope_query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies/cluster/" + clusterId + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GenerateNetworkPolicies(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,deleteExisting_string,networkDataSince_string,includePorts_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies/generate/" + clusterId + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetNetworkGraphEpoch(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("clusterId_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies/graph/epoch"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetUndoModificationForDeployment(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/undo/deployment/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetUndoModification(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/undo/" + clusterId + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetDiffFlowsFromUndoModificationForDeployment(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/undobaselinecomparison/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetNetworkPolicy(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}                // debug: {"detail": [{"operationId": "NetworkPolicyService_ApplyNetworkPolicyYamlForDeployment", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ApplyNetworkPolicyYamlForDeploymentRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_ApplyNetworkPolicy", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageNetworkPolicyModification"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_GetBaselineGeneratedNetworkPolicyForDeployment", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetBaselineGeneratedPolicyForDeploymentRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetBaselineGeneratedPolicyForDeploymentResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_SimulateNetworkGraph", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"description": "If set to true, include port-level information in the network policy graph.", "in": "query", "name": "includePorts", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "includeNodeDiff", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "scope.query", "required": false, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageNetworkPolicyModification"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SimulateNetworkGraphResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}, {"operationId": "NetworkPolicyService_SendNetworkPolicyYAML", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"explode": true, "in": "query", "name": "notifierIds", "required": false, "schema": {"items": {"type": "string"}, "type": "array"}}], "requestBody": {"$ref": "#/components/requestBodies/storageNetworkPolicyModification"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkPolicyService"]}], "method": "post"}

        
func (a NetworkPolicy) ApplyNetworkPolicyYamlForDeployment(deploymentId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/apply/deployment/" + deploymentId + ""

    tools.PostResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) ApplyNetworkPolicy(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/apply/" + clusterId + ""

    tools.PostResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) GetBaselineGeneratedNetworkPolicyForDeployment(deploymentId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkpolicies/generate/baseline/" + deploymentId + ""

    tools.PostResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) SimulateNetworkGraph(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,includePorts_boolean,includeNodeDiff_boolean,scope_query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies/simulate/" + clusterId + ""

    tools.PostResource(&a.Client, uriPath, args)

}
func (a NetworkPolicy) SendNetworkPolicyYAML(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("notifierIds_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkpolicies/simulate/" + clusterId + "/notify"

    tools.PostResource(&a.Client, uriPath, args)

}    