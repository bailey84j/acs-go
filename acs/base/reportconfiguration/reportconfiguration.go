package reportconfiguration

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ReportConfiguration struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ReportConfigurationService_CountReportConfigurations", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountReportConfigurationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountReportConfigurations returns the number of report configurations.", "tags": ["ReportConfigurationService"]}, {"operationId": "ReportConfigurationService_GetReportConfigurations", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetReportConfigurationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ReportConfigurationService"]}, {"operationId": "ReportConfigurationService_GetReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetReportConfigurationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ReportConfigurationService"]}], "method": "get"}

        
// CountReportConfigurations returns the number of report configurations.

func (a ReportConfiguration) CountReportConfigurations(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/report-configurations-count"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a ReportConfiguration) GetReportConfigurations(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/report/configurations"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a ReportConfiguration) GetReportConfiguration(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/report/configurations/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ReportConfigurationService_DeleteReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteReportConfiguration removes a report configuration given its id", "tags": ["ReportConfigurationService"]}], "method": "delete"}

        
// DeleteReportConfiguration removes a report configuration given its id

func (a ReportConfiguration) DeleteReportConfiguration(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/report/configurations/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"operationId": "ReportConfigurationService_PostReportConfiguration", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PostReportConfigurationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PostReportConfigurationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostReportConfiguration creates a report configuration", "tags": ["ReportConfigurationService"]}], "method": "post"}

        
// PostReportConfiguration creates a report configuration

func (a ReportConfiguration) PostReportConfiguration(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/report/configurations"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ReportConfigurationService_UpdateReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateReportConfigurationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UpdateReportConfiguration updates a report configuration", "tags": ["ReportConfigurationService"]}], "method": "put"}

        
// UpdateReportConfiguration updates a report configuration

func (a ReportConfiguration) UpdateReportConfiguration(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/report/configurations/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}