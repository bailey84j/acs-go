package networkgraph

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type NetworkGraph struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "NetworkGraphService_GetNetworkGraph", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "since", "required": false, "schema": {"format": "date-time", "type": "string"}}, {"in": "query", "name": "includePorts", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "scope.query", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkGraph"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}, {"operationId": "NetworkGraphService_GetExternalNetworkEntities", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetExternalNetworkEntitiesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}, {"operationId": "NetworkGraphService_GetNetworkGraphConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkGraphConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}], "method": "get"}

        
func (a NetworkGraph) GetNetworkGraph(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,since_string,includePorts_boolean,scope_query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkgraph/cluster/" + clusterId + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkGraph) GetExternalNetworkEntities(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/networkgraph/cluster/" + clusterId + "/externalentities"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a NetworkGraph) GetNetworkGraphConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkgraph/config"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "NetworkGraphService_DeleteExternalNetworkEntity", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}], "method": "delete"}

        
func (a NetworkGraph) DeleteExternalNetworkEntity(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkgraph/externalentities/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "NetworkGraphService_PatchExternalNetworkEntity", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PatchNetworkEntityRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkEntity"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}], "method": "patch"}

        
func (a NetworkGraph) PatchExternalNetworkEntity(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkgraph/externalentities/" + id + ""

    tools.PatchResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "NetworkGraphService_CreateExternalNetworkEntity", "parameters": [{"in": "path", "name": "clusterId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CreateNetworkEntityRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkEntity"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}], "method": "post"}

        
func (a NetworkGraph) CreateExternalNetworkEntity(clusterId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkgraph/cluster/" + clusterId + "/externalentities"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "NetworkGraphService_PutNetworkGraphConfig", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PutNetworkGraphConfigRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkGraphConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkGraphService"]}], "method": "put"}

        
func (a NetworkGraph) PutNetworkGraphConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkgraph/config"

    tools.PutResource(&a.Client, uriPath, args)

}