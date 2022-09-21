package db

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type DB struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "DBService_GetExportCapabilities", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetDBExportCapabilitiesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}, {"operationId": "DBService_GetActiveRestoreProcess", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetActiveDBRestoreProcessResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "get"}

        
func (a DB) GetExportCapabilities(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/db/exportcaps"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a DB) GetActiveRestoreProcess(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/db/restore"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "DBService_CancelRestoreProcess", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "delete"}

        
func (a DB) CancelRestoreProcess(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/db/restore/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"operationId": "DBService_InterruptRestoreProcess", "parameters": [{"in": "path", "name": "processId", "required": true, "schema": {"type": "string"}}, {"in": "path", "name": "attemptId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InterruptDBRestoreProcessResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "post"}

        
func (a DB) InterruptRestoreProcess(processId string,attemptId string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/db/interruptrestore/" + processId + "/" + attemptId + ""

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    