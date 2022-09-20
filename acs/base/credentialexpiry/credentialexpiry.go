package credentialexpiry

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type CredentialExpiry struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "CredentialExpiryService_GetCertExpiry", "parameters": [{"in": "query", "name": "component", "required": false, "schema": {"default": "UNKNOWN", "enum": ["UNKNOWN", "CENTRAL", "SCANNER"], "type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetCertExpiryResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetCertExpiry returns information related to the expiry component mTLS certificate.", "tags": ["CredentialExpiryService"]}], "method": "get"}

        
// GetCertExpiry returns information related to the expiry component mTLS certificate.

func (a CredentialExpiry) GetCertExpiry(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("component-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/credentialexpiry"

    tools.GetResource(&a.Client, uriPath, args)

}                