package alert

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Alert struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "AlertService_ListAlerts", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "List returns the slim list version of the alerts.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertsCounts", "parameters": [{"in": "query", "name": "request.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "request.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "request.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "request.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "request.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "groupBy", "required": false, "schema": {"default": "UNSET", "enum": ["UNSET", "CATEGORY", "CLUSTER"], "type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertsCountsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertsCounts returns the number of alerts in the requested cluster or category.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertsGroup", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertsGroupResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertsGroup returns alerts grouped by policy.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlertTimeseries", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAlertTimeseriesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlertTimeseries returns the alerts sorted by time.", "tags": ["AlertService"]}, {"operationId": "AlertService_GetAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAlert"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAlert returns the alert given its id.", "tags": ["AlertService"]}, {"operationId": "AlertService_CountAlerts", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountAlerts counts how many alerts match the get request.", "tags": ["AlertService"]}], "method": "get"}

        
// List returns the slim list version of the alerts.

func (a Alert) ListAlerts(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetAlertsCounts returns the number of alerts in the requested cluster or category.

func (a Alert) GetAlertsCounts(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("request_query-string,request_pagination_limit-integer,request_pagination_offset-integer,request_pagination_sortOption_field-string,request_pagination_sortOption_reversed-boolean,groupBy-string",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/counts"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetAlertsGroup returns alerts grouped by policy.

func (a Alert) GetAlertsGroup(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/groups"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetAlertTimeseries returns the alerts sorted by time.

func (a Alert) GetAlertTimeseries(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts/summary/timeseries"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetAlert returns the alert given its id.

func (a Alert) GetAlert(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/alerts/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// CountAlerts counts how many alerts match the get request.

func (a Alert) CountAlerts(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alertscount"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "AlertService_DeleteAlerts", "parameters": [{"in": "query", "name": "query.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "confirm", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeleteAlertsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AlertService"]}], "method": "delete"}

        
func (a Alert) DeleteAlerts(args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query_query-string,query_pagination_limit-integer,query_pagination_offset-integer,query_pagination_sortOption_field-string,query_pagination_sortOption_reversed-boolean,confirm-boolean",args)
    if !ok {
		return  fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/alerts"

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}        // debug: {"detail": [{"operationId": "AlertService_ResolveAlerts", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ResolveAlertsRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ResolveAlertsByQuery marks alerts matching search query as resolved.", "tags": ["AlertService"]}, {"operationId": "AlertService_ResolveAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ResolveAlertRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ResolveAlert marks the given alert (by ID) as resolved.", "tags": ["AlertService"]}, {"operationId": "AlertService_SnoozeAlert", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SnoozeAlertRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SnoozeAlert is deprecated.", "tags": ["AlertService"]}], "method": "patch"}

        
// ResolveAlertsByQuery marks alerts matching search query as resolved.

func (a Alert) ResolveAlerts(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/alerts/resolve"

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// ResolveAlert marks the given alert (by ID) as resolved.

func (a Alert) ResolveAlert(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/alerts/" + id + "/resolve"

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// SnoozeAlert is deprecated.

func (a Alert) SnoozeAlert(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/alerts/" + id + "/snooze"

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        