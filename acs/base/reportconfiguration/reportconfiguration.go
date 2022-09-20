package reportconfiguration

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ReportConfiguration struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ReportConfigurationService_CountReportConfigurations", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountReportConfigurationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountReportConfigurations returns the number of report configurations.", "tags": ["ReportConfigurationService"]}, {"operationId": "ReportConfigurationService_GetReportConfigurations", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetReportConfigurationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ReportConfigurationService"]}, {"operationId": "ReportConfigurationService_GetReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetReportConfigurationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ReportConfigurationService"]}], "method": "get"}

        
// CountReportConfigurations returns the number of report configurations.

func (a ReportConfiguration) CountReportConfigurations(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/report-configurations-count"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a ReportConfiguration) GetReportConfigurations(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/report/configurations"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a ReportConfiguration) GetReportConfiguration(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/report/configurations/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "ReportConfigurationService_DeleteReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteReportConfiguration removes a report configuration given its id", "tags": ["ReportConfigurationService"]}], "method": "delete"}

        
// DeleteReportConfiguration removes a report configuration given its id

func (a ReportConfiguration) DeleteReportConfiguration(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/report/configurations/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "ReportConfigurationService_PostReportConfiguration", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PostReportConfigurationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PostReportConfigurationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostReportConfiguration creates a report configuration", "tags": ["ReportConfigurationService"]}], "method": "post"}

        
// PostReportConfiguration creates a report configuration

func (a ReportConfiguration) PostReportConfiguration(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/report/configurations"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "ReportConfigurationService_UpdateReportConfiguration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateReportConfigurationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UpdateReportConfiguration updates a report configuration", "tags": ["ReportConfigurationService"]}], "method": "put"}

        
// UpdateReportConfiguration updates a report configuration

func (a ReportConfiguration) UpdateReportConfiguration(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/report/configurations/" + id + ""

    tools.PutResource(&a.Client, uriPath, args)

}