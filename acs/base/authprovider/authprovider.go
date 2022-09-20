package authprovider

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type AuthProvider struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "AuthProviderService_GetAuthProviders", "parameters": [{"in": "query", "name": "name", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "type", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetAuthProvidersResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}, {"operationId": "AuthProviderService_GetAuthProvider", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAuthProvider"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}, {"operationId": "AuthProviderService_ListAvailableProviderTypes", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1AvailableProviderTypesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}, {"operationId": "AuthProviderService_GetLoginAuthProviders", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetLoginAuthProvidersResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}], "method": "get"}

        
func (a AuthProvider) GetAuthProviders(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("name_string,type_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/authProviders"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a AuthProvider) GetAuthProvider(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a AuthProvider) ListAvailableProviderTypes(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/availableAuthProviders"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a AuthProvider) GetLoginAuthProviders(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/login/authproviders"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AuthProviderService_DeleteAuthProvider", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}], "method": "delete"}

        
func (a AuthProvider) DeleteAuthProvider(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AuthProviderService_UpdateAuthProvider", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateAuthProviderRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAuthProvider"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}], "method": "patch"}

        
func (a AuthProvider) UpdateAuthProvider(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders/" + id + ""

    tools.PatchResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AuthProviderService_PostAuthProvider", "requestBody": {"$ref": "#/components/requestBodies/storageAuthProvider"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAuthProvider"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}, {"operationId": "AuthProviderService_ExchangeToken", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ExchangeTokenRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ExchangeTokenResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}], "method": "post"}

        
func (a AuthProvider) PostAuthProvider(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders"

    tools.PostResource(&a.Client, uriPath, args)

}
func (a AuthProvider) ExchangeToken(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders/exchangeToken"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "AuthProviderService_PutAuthProvider", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageAuthProvider"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageAuthProvider"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthProviderService"]}], "method": "put"}

        
func (a AuthProvider) PutAuthProvider(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/authProviders/" + id + ""

    tools.PutResource(&a.Client, uriPath, args)

}