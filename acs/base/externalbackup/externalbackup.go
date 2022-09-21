package externalbackup

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ExternalBackup struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ExternalBackupService_GetExternalBackups", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetExternalBackupsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetExternalBackups returns all external backup configurations.", "tags": ["ExternalBackupService"]}, {"operationId": "ExternalBackupService_GetExternalBackup", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageExternalBackup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetExternalBackup returns the external backup configuration given its ID.", "tags": ["ExternalBackupService"]}], "method": "get"}

        
// GetExternalBackups returns all external backup configurations.

func (a ExternalBackup) GetExternalBackups(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetExternalBackup returns the external backup configuration given its ID.

func (a ExternalBackup) GetExternalBackup(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ExternalBackupService_DeleteExternalBackup", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteExternalBackup removes an external backup configuration given its ID.", "tags": ["ExternalBackupService"]}], "method": "delete"}

        
// DeleteExternalBackup removes an external backup configuration given its ID.

func (a ExternalBackup) DeleteExternalBackup(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}        // debug: {"detail": [{"operationId": "ExternalBackupService_UpdateExternalBackup", "parameters": [{"in": "path", "name": "externalBackup.id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/v1UpdateExternalBackupRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageExternalBackup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UpdateExternalBackup modifies a given external backup, with optional stored credential reconciliation.", "tags": ["ExternalBackupService"]}], "method": "patch"}

        
// UpdateExternalBackup modifies a given external backup, with optional stored credential reconciliation.

func (a ExternalBackup) UpdateExternalBackup(externalBackup_id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/" + externalBackup_id + ""

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ExternalBackupService_PostExternalBackup", "requestBody": {"$ref": "#/components/requestBodies/storageExternalBackup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageExternalBackup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostExternalBackup creates an external backup configuration.", "tags": ["ExternalBackupService"]}, {"operationId": "ExternalBackupService_TestExternalBackup", "requestBody": {"$ref": "#/components/requestBodies/storageExternalBackup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestExternalBackup tests an external backup configuration.", "tags": ["ExternalBackupService"]}, {"operationId": "ExternalBackupService_TestUpdatedExternalBackup", "requestBody": {"$ref": "#/components/requestBodies/v1UpdateExternalBackupRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestUpdatedExternalBackup checks if the given external backup is correctly configured, with optional stored credential reconciliation.", "tags": ["ExternalBackupService"]}, {"operationId": "ExternalBackupService_TriggerExternalBackup", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TriggerExternalBackup initiates an external backup for the given configuration.", "tags": ["ExternalBackupService"]}], "method": "post"}

        
// PostExternalBackup creates an external backup configuration.

func (a ExternalBackup) PostExternalBackup(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestExternalBackup tests an external backup configuration.

func (a ExternalBackup) TestExternalBackup(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/test"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestUpdatedExternalBackup checks if the given external backup is correctly configured, with optional stored credential reconciliation.

func (a ExternalBackup) TestUpdatedExternalBackup(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/test/updated"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TriggerExternalBackup initiates an external backup for the given configuration.

func (a ExternalBackup) TriggerExternalBackup(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/" + id + ""

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ExternalBackupService_PutExternalBackup", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageExternalBackup"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageExternalBackup"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PutExternalBackup modifies a given external backup, without using stored credential reconciliation.", "tags": ["ExternalBackupService"]}], "method": "put"}

        
// PutExternalBackup modifies a given external backup, without using stored credential reconciliation.

func (a ExternalBackup) PutExternalBackup(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/externalbackups/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}