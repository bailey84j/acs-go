package alert

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Alert struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "AlertService_ListAlerts", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "List returns the slim list version of the alerts.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertsCounts", "parameters": [{"in": "query", "name": "request.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "request.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "request.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "request.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "request.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "groupBy", "required": false, "schema": {"default": "UNSET", "enum": ["UNSET", "CATEGORY", "CLUSTER"], "type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertsCountsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertsCounts returns the number of alerts in the requested cluster or category.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertsGroup", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertsGroupResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertsGroup returns alerts grouped by policy.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertTimeseries", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertTimeseriesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertTimeseries returns the alerts sorted by time.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAlert"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlert returns the alert given its id.", "tags": ["AlertService"]}, {"operationId": "AlertService_CountAlerts", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountAlerts counts how many alerts match the get request.", "tags": ["AlertService"]}], "method": "get"}

        
// List returns the slim list version of the alerts.

func (a Alert) ListAlerts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetAlertsCounts returns the number of alerts in the requested cluster or category.

func (a Alert) GetAlertsCounts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("request_query_string,request_pagination_limit_integer,request_pagination_offset_integer,request_pagination_sortOption_field_string,request_pagination_sortOption_reversed_boolean,groupBy_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/counts"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetAlertsGroup returns alerts grouped by policy.

func (a Alert) GetAlertsGroup(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/groups"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetAlertTimeseries returns the alerts sorted by time.

func (a Alert) GetAlertTimeseries(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/timeseries"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetAlert returns the alert given its id.

func (a Alert) GetAlert(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/alerts/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
// CountAlerts counts how many alerts match the get request.

func (a Alert) CountAlerts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alertscount"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AlertService_DeleteAlerts", "parameters": [{"in": "query", "name": "query.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "confirm", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeleteAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AlertService"]}], "method": "delete"}

        
func (a Alert) DeleteAlerts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_query_string,query_pagination_limit_integer,query_pagination_offset_integer,query_pagination_sortOption_field_string,query_pagination_sortOption_reversed_boolean,confirm_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts"

    tools.DeleteResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AlertService_ResolveAlerts", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ResolveAlertsRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ResolveAlertsByQuery marks alerts matching search query as resolved.", "tags": ["AlertService"]}, {"operationId": "AlertService_ResolveAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ResolveAlertRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ResolveAlert marks the given alert (by ID) as resolved.", "tags": ["AlertService"]}, {"operationId": "AlertService_SnoozeAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SnoozeAlertRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SnoozeAlert is deprecated.", "tags": ["AlertService"]}], "method": "patch"}

        
// ResolveAlertsByQuery marks alerts matching search query as resolved.

func (a Alert) ResolveAlerts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/alerts/resolve"

    tools.PatchResource(&a.Client, uriPath, args)

}
// ResolveAlert marks the given alert (by ID) as resolved.

func (a Alert) ResolveAlert(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/alerts/" + id + "/resolve"

    tools.PatchResource(&a.Client, uriPath, args)

}
// SnoozeAlert is deprecated.

func (a Alert) SnoozeAlert(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/alerts/" + id + "/snooze"

    tools.PatchResource(&a.Client, uriPath, args)

}        