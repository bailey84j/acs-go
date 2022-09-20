package deployment

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Deployment struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "DeploymentService_ListDeployments", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListDeploymentsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListDeployments returns the list of deployments.", "tags": ["DeploymentService"]}, {"operationId": "DeploymentService_GetLabels", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeploymentLabelsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetLabels returns the labels used by deployments.", "tags": ["DeploymentService"]}, {"operationId": "DeploymentService_GetDeployment", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageDeployment"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetDeployment returns a deployment given its ID.", "tags": ["DeploymentService"]}, {"operationId": "DeploymentService_CountDeployments", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountDeploymentsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountDeployments returns the number of deployments.", "tags": ["DeploymentService"]}, {"operationId": "DeploymentService_ListDeploymentsWithProcessInfo", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListDeploymentsWithProcessInfoResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListDeploymentsWithProcessInfo returns the list of deployments with process information.", "tags": ["DeploymentService"]}, {"operationId": "DeploymentService_GetDeploymentWithRisk", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetDeploymentWithRiskResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetDeploymentWithRisk returns a deployment and its risk given its ID.", "tags": ["DeploymentService"]}], "method": "get"}

        
// ListDeployments returns the list of deployments.

func (a Deployment) ListDeployments(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/deployments"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetLabels returns the labels used by deployments.

func (a Deployment) GetLabels(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/deployments/metadata/labels"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetDeployment returns a deployment given its ID.

func (a Deployment) GetDeployment(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/deployments/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
// CountDeployments returns the number of deployments.

func (a Deployment) CountDeployments(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/deploymentscount"

    tools.GetResource(&a.Client, uriPath, args)

}
// ListDeploymentsWithProcessInfo returns the list of deployments with process information.

func (a Deployment) ListDeploymentsWithProcessInfo(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/deploymentswithprocessinfo"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetDeploymentWithRisk returns a deployment and its risk given its ID.

func (a Deployment) GetDeploymentWithRisk(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/deploymentswithrisk/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}                