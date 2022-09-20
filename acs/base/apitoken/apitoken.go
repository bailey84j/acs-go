package apitoken

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type APIToken struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "APITokenService_GetAPITokens", "parameters": [{"in": "query", "name": "revoked", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAPITokensResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAPITokens returns all the API tokens.", "tags": ["APITokenService"]}, {"operationId": "APITokenService_GetAPIToken", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageTokenMetadata"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetAPIToken returns API token metadata for a given id.", "tags": ["APITokenService"]}], "method": "get"}

        
// GetAPITokens returns all the API tokens.

func (a APIToken) GetAPITokens(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("revoked-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/apitokens"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetAPIToken returns API token metadata for a given id.

func (a APIToken) GetAPIToken(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/apitokens/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}            // debug: {"detail": [{"operationId": "APITokenService_RevokeToken", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "RevokeToken removes the API token for a given id.", "tags": ["APITokenService"]}], "method": "patch"}

        
// RevokeToken removes the API token for a given id.

func (a APIToken) RevokeToken(id string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/apitokens/revoke/" + id + ""

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "APITokenService_GenerateToken", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GenerateTokenRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GenerateTokenResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GenerateToken generates API token for a given user and role.", "tags": ["APITokenService"]}], "method": "post"}

        
// GenerateToken generates API token for a given user and role.

func (a APIToken) GenerateToken(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/apitokens/generate"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    