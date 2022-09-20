package group

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Group struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "GroupService_GetGroup", "parameters": [{"description": "Unique identifier for group properties and respectively the group.", "in": "query", "name": "id", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageGroup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}, {"operationId": "GroupService_GetGroups", "parameters": [{"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "id", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetGroupsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "get"}

        
func (a Group) GetGroup(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("id-string,authProviderId-string,key-string,value-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/group"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Group) GetGroups(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("authProviderId-string,key-string,value-string,id-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/groups"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "GroupService_DeleteGroup", "parameters": [{"description": "Unique identifier for group properties and respectively the group.", "in": "query", "name": "id", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "authProviderId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "key", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "value", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "delete"}

        
func (a Group) DeleteGroup(args map[string]interface{}) ( error) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("id-string,authProviderId-string,key-string,value-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/groups"

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"operationId": "GroupService_CreateGroup", "requestBody": {"$ref": "#/components/requestBodies/storageGroup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}, {"operationId": "GroupService_BatchUpdate", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GroupBatchUpdateRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "post"}

        
func (a Group) CreateGroup(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groups"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Group) BatchUpdate(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groupsbatch"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "GroupService_UpdateGroup", "requestBody": {"$ref": "#/components/requestBodies/storageGroup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["GroupService"]}], "method": "put"}

        
func (a Group) UpdateGroup(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/groups"

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}