package compliancemanagement

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ComplianceManagement struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ComplianceManagementService_GetRecentRuns", "parameters": [{"in": "query", "name": "clusterId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "standardId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "since", "required": false, "schema": {"format": "date-time", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetRecentComplianceRunsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceManagementService"]}, {"operationId": "ComplianceManagementService_GetRunStatuses", "parameters": [{"explode": true, "in": "query", "name": "runIds", "required": false, "schema": {"items": {"type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetComplianceRunStatusesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceManagementService"]}], "method": "get"}

        
func (a ComplianceManagement) GetRecentRuns(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("clusterId-string,standardId-string,since-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/complianceManagement/runs"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a ComplianceManagement) GetRunStatuses(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("runIds-array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/compliancemanagement/runstatuses"

    tools.GetResource(&a.Client, uriPath, args)

}                // debug: {"detail": [{"operationId": "ComplianceManagementService_TriggerRuns", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1TriggerComplianceRunsRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1TriggerComplianceRunsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceManagementService"]}], "method": "post"}

        
func (a ComplianceManagement) TriggerRuns(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/compliancemanagement/runs"

    tools.PostResource(&a.Client, uriPath, args)

}    