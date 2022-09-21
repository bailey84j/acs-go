package notifier

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Notifier struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "NotifierService_GetNotifiers", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetNotifiersResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetNotifiers returns all notifier configurations.", "tags": ["NotifierService"]}, {"operationId": "NotifierService_GetNotifier", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNotifier"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetNotifier returns the notifier configuration given its ID.", "tags": ["NotifierService"]}], "method": "get"}

        
// GetNotifiers returns all notifier configurations.

func (a Notifier) GetNotifiers(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetNotifier returns the notifier configuration given its ID.

func (a Notifier) GetNotifier(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "NotifierService_DeleteNotifier", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "force", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteNotifier removes a notifier configuration given its ID.", "tags": ["NotifierService"]}], "method": "delete"}

        
// DeleteNotifier removes a notifier configuration given its ID.

func (a Notifier) DeleteNotifier(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("force-boolean",args)
    if !ok {
		return  fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/notifiers/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}        // debug: {"detail": [{"operationId": "NotifierService_UpdateNotifier", "parameters": [{"in": "path", "name": "notifier.id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/v1UpdateNotifierRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UpdateNotifier modifies a given notifier, with optional stored credential reconciliation.", "tags": ["NotifierService"]}], "method": "patch"}

        
// UpdateNotifier modifies a given notifier, with optional stored credential reconciliation.

func (a Notifier) UpdateNotifier(notifier_id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers/" + notifier_id + ""

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "NotifierService_PostNotifier", "requestBody": {"$ref": "#/components/requestBodies/storageNotifier"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageNotifier"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostNotifier creates a notifier configuration.", "tags": ["NotifierService"]}, {"operationId": "NotifierService_TestNotifier", "requestBody": {"$ref": "#/components/requestBodies/storageNotifier"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestNotifier checks if a notifier is correctly configured.", "tags": ["NotifierService"]}, {"operationId": "NotifierService_TestUpdatedNotifier", "requestBody": {"$ref": "#/components/requestBodies/v1UpdateNotifierRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestUpdatedNotifier checks if the given notifier is correctly configured, with optional stored credential reconciliation.", "tags": ["NotifierService"]}], "method": "post"}

        
// PostNotifier creates a notifier configuration.

func (a Notifier) PostNotifier(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestNotifier checks if a notifier is correctly configured.

func (a Notifier) TestNotifier(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers/test"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestUpdatedNotifier checks if the given notifier is correctly configured, with optional stored credential reconciliation.

func (a Notifier) TestUpdatedNotifier(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers/test/updated"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "NotifierService_PutNotifier", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageNotifier"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PutNotifier modifies a given notifier, without using stored credential reconciliation.", "tags": ["NotifierService"]}], "method": "put"}

        
// PutNotifier modifies a given notifier, without using stored credential reconciliation.

func (a Notifier) PutNotifier(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/notifiers/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}