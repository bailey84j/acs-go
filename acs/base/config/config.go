package config

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Config struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ConfigService_GetConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ConfigService"]}, {"operationId": "ConfigService_GetPrivateConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePrivateConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ConfigService"]}, {"operationId": "ConfigService_GetPublicConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePublicConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ConfigService"]}], "method": "get"}

        
func (a Config) GetConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/config"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Config) GetPrivateConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/config/private"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Config) GetPublicConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/config/public"

    tools.GetResource(&a.Client, uriPath, args)

}                    // debug: {"detail": [{"operationId": "ConfigService_PutConfig", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PutConfigRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageConfig"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ConfigService"]}], "method": "put"}

        
func (a Config) PutConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/config"

    tools.PutResource(&a.Client, uriPath, args)

}