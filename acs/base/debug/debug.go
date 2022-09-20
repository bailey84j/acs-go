package debug

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Debug struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "DebugService_StreamAuthzTraces", "responses": {"200": {"content": {"application/json": {"schema": {"properties": {"error": {"$ref": "#/components/schemas/runtimeStreamError"}, "result": {"$ref": "#/components/schemas/v1AuthorizationTraceResponse"}}, "title": "Stream result of v1AuthorizationTraceResponse", "type": "object"}}}, "description": "A successful response.(streaming responses)"}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Stream authorization traces for all incoming requests.", "tags": ["DebugService"]}, {"operationId": "DebugService_GetLogLevel", "parameters": [{"explode": true, "in": "query", "name": "modules", "required": false, "schema": {"items": {"type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1LogLevelResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Get the current logging level for StackRox services.", "tags": ["DebugService"]}], "method": "get"}

        
// Stream authorization traces for all incoming requests.

func (a Debug) StreamAuthzTraces(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/debug/authz/trace"

    tools.GetResource(&a.Client, uriPath, args)

}
// Get the current logging level for StackRox services.

func (a Debug) GetLogLevel(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("modules-array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/debug/loglevel"

    tools.GetResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "DebugService_SetLogLevel", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1LogLevelRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"properties": {}}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Set logging level for StackRox services.", "tags": ["DebugService"]}], "method": "patch"}

        
// Set logging level for StackRox services.

func (a Debug) SetLogLevel(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/debug/loglevel"

    tools.PatchResource(&a.Client, uriPath, args)

}        