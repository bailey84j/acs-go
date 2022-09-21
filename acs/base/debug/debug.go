package debug

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Debug struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "DebugService_StreamAuthzTraces", "responses": {"200": {"content": {"application/json": {"schema": {"properties": {"error": {"$ref": "#/components/schemas/runtimeStreamError"}, "result": {"$ref": "#/components/schemas/v1AuthorizationTraceResponse"}}, "title": "Stream result of v1AuthorizationTraceResponse", "type": "object"}}}, "description": "A successful response.(streaming responses)"}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Stream authorization traces for all incoming requests.", "tags": ["DebugService"]}, {"operationId": "DebugService_GetLogLevel", "parameters": [{"explode": true, "in": "query", "name": "modules", "required": false, "schema": {"items": {"type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1LogLevelResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Get the current logging level for StackRox services.", "tags": ["DebugService"]}], "method": "get"}

        
// Stream authorization traces for all incoming requests.

func (a Debug) StreamAuthzTraces(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/debug/authz/trace"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// Get the current logging level for StackRox services.

func (a Debug) GetLogLevel(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("modules-array",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/debug/loglevel"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}            // debug: {"detail": [{"operationId": "DebugService_SetLogLevel", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1LogLevelRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"properties": {}}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Set logging level for StackRox services.", "tags": ["DebugService"]}], "method": "patch"}

        
// Set logging level for StackRox services.

func (a Debug) SetLogLevel(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/debug/loglevel"

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        