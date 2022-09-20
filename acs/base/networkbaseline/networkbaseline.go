package networkbaseline

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type NetworkBaseline struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "NetworkBaselineService_GetNetworkBaseline", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNetworkBaseline"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkBaselineService"]}], "method": "get"}

        
func (a NetworkBaseline) GetNetworkBaseline(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkbaseline/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "NetworkBaselineService_ModifyBaselineStatusForPeers", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ModifyBaselineStatusForPeersRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkBaselineService"]}, {"operationId": "NetworkBaselineService_LockNetworkBaseline", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/v1ResourceByID"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkBaselineService"]}, {"operationId": "NetworkBaselineService_UnlockNetworkBaseline", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/v1ResourceByID"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkBaselineService"]}], "method": "patch"}

        
func (a NetworkBaseline) ModifyBaselineStatusForPeers(deploymentId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkbaseline/" + deploymentId + "/peers"

    tools.PatchResource(&a.Client, uriPath, args)

}
func (a NetworkBaseline) LockNetworkBaseline(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkbaseline/" + id + "/lock"

    tools.PatchResource(&a.Client, uriPath, args)

}
func (a NetworkBaseline) UnlockNetworkBaseline(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkbaseline/" + id + "/unlock"

    tools.PatchResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "NetworkBaselineService_GetNetworkBaselineStatusForFlows", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkBaselineStatusRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1NetworkBaselineStatusResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NetworkBaselineService"]}], "method": "post"}

        
func (a NetworkBaseline) GetNetworkBaselineStatusForFlows(deploymentId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/networkbaseline/" + deploymentId + "/status"

    tools.PostResource(&a.Client, uriPath, args)

}    