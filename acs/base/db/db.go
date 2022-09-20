package db

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type DB struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "DBService_GetExportCapabilities", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetDBExportCapabilitiesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}, {"operationId": "DBService_GetActiveRestoreProcess", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetActiveDBRestoreProcessResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "get"}

        
func (a DB) GetExportCapabilities(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/db/exportcaps"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a DB) GetActiveRestoreProcess(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/db/restore"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "DBService_CancelRestoreProcess", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "delete"}

        
func (a DB) CancelRestoreProcess(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/db/restore/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "DBService_InterruptRestoreProcess", "parameters": [{"in": "path", "name": "processId", "required": true, "schema": {"type": "string"}}, {"in": "path", "name": "attemptId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InterruptDBRestoreProcessResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["DBService"]}], "method": "post"}

        
func (a DB) InterruptRestoreProcess(processId string,attemptId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/db/interruptrestore/" + processId + "/" + attemptId + ""

    tools.PostResource(&a.Client, uriPath, args)

}    