package group

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Group struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "GroupService_GetGroup", "parameters": [{"description": "Unique identifier for group properties and respectively the group.", "in": "query", "name": "id", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageGroup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}, {"operationId": "GroupService_GetGroups", "parameters": [{"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "id", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetGroupsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "get"}

        
func (a Group) GetGroup(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("id_string,authProviderId_string,key_string,value_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/group"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Group) GetGroups(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("authProviderId_string,key_string,value_string,id_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/groups"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "GroupService_DeleteGroup", "parameters": [{"description": "Unique identifier for group properties and respectively the group.", "in": "query", "name": "id", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "delete"}

        
func (a Group) DeleteGroup(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("id_string,authProviderId_string,key_string,value_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/groups"

    tools.DeleteResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "GroupService_CreateGroup", "requestBody": {"$ref": "#/components/requestBodies/storageGroup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}, {"operationId": "GroupService_BatchUpdate", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GroupBatchUpdateRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "post"}

        
func (a Group) CreateGroup(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groups"

    tools.PostResource(&a.Client, uriPath, args)

}
func (a Group) BatchUpdate(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groupsbatch"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "GroupService_UpdateGroup", "requestBody": {"$ref": "#/components/requestBodies/storageGroup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "put"}

        
func (a Group) UpdateGroup(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groups"

    tools.PutResource(&a.Client, uriPath, args)

}