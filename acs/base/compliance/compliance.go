package compliance

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Compliance struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ComplianceService_GetAggregatedResults", "parameters": [{"explode": true, "in": "query", "name": "groupBy", "required": false, "schema": {"items": {"enum": ["UNKNOWN", "STANDARD", "CLUSTER", "CATEGORY", "CONTROL", "NAMESPACE", "NODE", "DEPLOYMENT", "CHECK"], "type": "string"}, "type": "array"}}, {"in": "query", "name": "unit", "required": false, "schema": {"default": "UNKNOWN", "enum": ["UNKNOWN", "STANDARD", "CLUSTER", "CATEGORY", "CONTROL", "NAMESPACE", "NODE", "DEPLOYMENT", "CHECK"], "type": "string"}}, {"in": "query", "name": "where.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "where.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "where.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "where.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "where.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageComplianceAggregationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceService"]}, {"operationId": "ComplianceService_GetRunResults", "parameters": [{"in": "query", "name": "clusterId", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "standardId", "required": false, "schema": {"type": "string"}}, {"description": "Specifies the run ID for which to return results. If empty, the most recent run is returned.\nCAVEAT: Setting this field circumvents the results cache on the server-side, which may lead to significantly\n        increased memory pressure and decreased performance.", "in": "query", "name": "runId", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetComplianceRunResultsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceService"]}, {"operationId": "ComplianceService_GetStandards", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetComplianceStandardsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceService"]}, {"operationId": "ComplianceService_GetStandard", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetComplianceStandardResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ComplianceService"]}], "method": "get"}

        
func (a Compliance) GetAggregatedResults(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("groupBy-array,unit-string,where_query-string,where_pagination_limit-integer,where_pagination_offset-integer,where_pagination_sortOption_field-string,where_pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/compliance/aggregatedresults"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Compliance) GetRunResults(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("clusterId-string,standardId-string,runId-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/compliance/runresults"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Compliance) GetStandards(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/compliance/standards"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Compliance) GetStandard(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/compliance/standards/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}                