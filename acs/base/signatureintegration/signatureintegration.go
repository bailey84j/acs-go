package signatureintegration

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type SignatureIntegration struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "SignatureIntegrationService_ListSignatureIntegrations", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListSignatureIntegrationsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SignatureIntegrationService"]}, {"operationId": "SignatureIntegrationService_GetSignatureIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageSignatureIntegration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SignatureIntegrationService"]}], "method": "get"}

        
func (a SignatureIntegration) ListSignatureIntegrations(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/signatureintegrations"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a SignatureIntegration) GetSignatureIntegration(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/signatureintegrations/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "SignatureIntegrationService_DeleteSignatureIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SignatureIntegrationService"]}], "method": "delete"}

        
func (a SignatureIntegration) DeleteSignatureIntegration(id string,args map[string]interface{}) ( error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/signatureintegrations/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"operationId": "SignatureIntegrationService_PostSignatureIntegration", "requestBody": {"$ref": "#/components/requestBodies/storageSignatureIntegration"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageSignatureIntegration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "Integration id should not be set.\nReturns signature integration with id filled.", "tags": ["SignatureIntegrationService"]}], "method": "post"}

        
// Integration id should not be set. Returns signature integration with id filled.

func (a SignatureIntegration) PostSignatureIntegration(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/signatureintegrations"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "SignatureIntegrationService_PutSignatureIntegration", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageSignatureIntegration"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SignatureIntegrationService"]}], "method": "put"}

        
func (a SignatureIntegration) PutSignatureIntegration(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/signatureintegrations/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}