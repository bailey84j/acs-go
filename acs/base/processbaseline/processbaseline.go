package processbaseline

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ProcessBaseline struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ProcessBaselineService_GetProcessBaseline", "parameters": [{"description": "The idea is for the keys to be flexible.\nOnly certain combinations of these will be supported.", "in": "query", "name": "key.deploymentId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key.containerName", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key.clusterId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key.namespace", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageProcessBaseline"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "`GetProcessBaselineById` returns the single\nprocess baseline referenced by the given ID.", "tags": ["ProcessBaselineService"]}], "method": "get"}

        
// `GetProcessBaselineById` returns the single process baseline referenced by the given ID.

func (a ProcessBaseline) GetProcessBaseline(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("key_deploymentId_string,key_containerName_string,key_clusterId_string,key_namespace_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/processbaselines/key"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "ProcessBaselineService_DeleteProcessBaselines", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "confirm", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeleteProcessBaselinesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "`DeleteProcessBaselines` deletes baselines.", "tags": ["ProcessBaselineService"]}], "method": "delete"}

        
// `DeleteProcessBaselines` deletes baselines.

func (a ProcessBaseline) DeleteProcessBaselines(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,confirm_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/processbaselines"

    tools.DeleteResource(&a.Client, uriPath, args)

}                // debug: {"detail": [{"operationId": "ProcessBaselineService_UpdateProcessBaselines", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateProcessBaselinesRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateProcessBaselinesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "`AddToProcessBaselines` adds a list of process\nnames to each of a list of process baselines.", "tags": ["ProcessBaselineService"]}, {"operationId": "ProcessBaselineService_LockProcessBaselines", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1LockProcessBaselinesRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateProcessBaselinesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "`LockProcessBaselines` accepts a list of baseline IDs, locks\nthose baselines, and returns the updated baseline objects.", "tags": ["ProcessBaselineService"]}], "method": "put"}

        
// `AddToProcessBaselines` adds a list of process names to each of a list of process baselines.

func (a ProcessBaseline) UpdateProcessBaselines(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/processbaselines"

    tools.PutResource(&a.Client, uriPath, args)

}
// `LockProcessBaselines` accepts a list of baseline IDs, locks those baselines, and returns the updated baseline objects.

func (a ProcessBaseline) LockProcessBaselines(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/processbaselines/lock"

    tools.PutResource(&a.Client, uriPath, args)

}