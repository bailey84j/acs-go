package imageintegration

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ImageIntegration struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ImageIntegrationService_GetImageIntegrations", "parameters": [{"in": "query", "name": "name", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "cluster", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetImageIntegrationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetImageIntegrations returns all image integrations that match the request filters.", "tags": ["ImageIntegrationService"]}, {"operationId": "ImageIntegrationService_GetImageIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageImageIntegration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetImageIntegration returns the image integration given its ID.", "tags": ["ImageIntegrationService"]}], "method": "get"}

        
// GetImageIntegrations returns all image integrations that match the request filters.

func (a ImageIntegration) GetImageIntegrations(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("name-string,cluster-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/imageintegrations"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetImageIntegration returns the image integration given its ID.

func (a ImageIntegration) GetImageIntegration(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ImageIntegrationService_DeleteImageIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteImageIntegration removes a image integration given its ID.", "tags": ["ImageIntegrationService"]}], "method": "delete"}

        
// DeleteImageIntegration removes a image integration given its ID.

func (a ImageIntegration) DeleteImageIntegration(id string,args map[string]interface{}) ( error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}        // debug: {"detail": [{"operationId": "ImageIntegrationService_UpdateImageIntegration", "parameters": [{"in": "path", "name": "config.id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/v1UpdateImageIntegrationRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UpdateImageIntegration modifies a given image integration, with optional stored credential reconciliation.", "tags": ["ImageIntegrationService"]}], "method": "patch"}

        
// UpdateImageIntegration modifies a given image integration, with optional stored credential reconciliation.

func (a ImageIntegration) UpdateImageIntegration(config_id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/" + config_id + ""

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ImageIntegrationService_PostImageIntegration", "requestBody": {"$ref": "#/components/requestBodies/storageImageIntegration"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageImageIntegration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostImageIntegration creates a image integration.", "tags": ["ImageIntegrationService"]}, {"operationId": "ImageIntegrationService_TestImageIntegration", "requestBody": {"$ref": "#/components/requestBodies/storageImageIntegration"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestImageIntegration checks if the given image integration is correctly configured, without using stored credential reconciliation.", "tags": ["ImageIntegrationService"]}, {"operationId": "ImageIntegrationService_TestUpdatedImageIntegration", "requestBody": {"$ref": "#/components/requestBodies/v1UpdateImageIntegrationRequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "TestUpdatedImageIntegration checks if the given image integration is correctly configured, with optional stored credential reconciliation.", "tags": ["ImageIntegrationService"]}], "method": "post"}

        
// PostImageIntegration creates a image integration.

func (a ImageIntegration) PostImageIntegration(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestImageIntegration checks if the given image integration is correctly configured, without using stored credential reconciliation.

func (a ImageIntegration) TestImageIntegration(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/test"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// TestUpdatedImageIntegration checks if the given image integration is correctly configured, with optional stored credential reconciliation.

func (a ImageIntegration) TestUpdatedImageIntegration(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/test/updated"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ImageIntegrationService_PutImageIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageImageIntegration"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PutImageIntegration modifies a given image integration, without using stored credential reconciliation.", "tags": ["ImageIntegrationService"]}], "method": "put"}

        
// PutImageIntegration modifies a given image integration, without using stored credential reconciliation.

func (a ImageIntegration) PutImageIntegration(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imageintegrations/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}