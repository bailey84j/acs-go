package integrationhealth

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type IntegrationHealth struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "IntegrationHealthService_GetBackupPlugins", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetImageIntegrations", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetNotifiers", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetVulnDefinitionsInfo", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1VulnDefinitionsInfo"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}], "method": "get"}

        
func (a IntegrationHealth) GetBackupPlugins(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/externalbackups"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a IntegrationHealth) GetImageIntegrations(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/imageintegrations"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a IntegrationHealth) GetNotifiers(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/notifiers"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a IntegrationHealth) GetVulnDefinitionsInfo(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/vulndefinitions"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                