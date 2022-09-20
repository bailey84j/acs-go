package integrationhealth

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type IntegrationHealth struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "IntegrationHealthService_GetBackupPlugins", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetImageIntegrations", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetNotifiers", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetIntegrationHealthResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}, {"operationId": "IntegrationHealthService_GetVulnDefinitionsInfo", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1VulnDefinitionsInfo"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["IntegrationHealthService"]}], "method": "get"}

        
func (a IntegrationHealth) GetBackupPlugins(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/externalbackups"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a IntegrationHealth) GetImageIntegrations(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/imageintegrations"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a IntegrationHealth) GetNotifiers(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/notifiers"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a IntegrationHealth) GetVulnDefinitionsInfo(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/integrationhealth/vulndefinitions"

    tools.GetResource(&a.Client, uriPath, args)

}                