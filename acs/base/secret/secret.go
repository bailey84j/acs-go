package secret

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Secret struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "SecretService_ListSecrets", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListSecretsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListSecrets returns the list of secrets.", "tags": ["SecretService"]}, {"operationId": "SecretService_GetSecret", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageSecret"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetSecret returns a secret given its ID.", "tags": ["SecretService"]}, {"operationId": "SecretService_CountSecrets", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountSecretsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "CountSecrets returns the number of secrets.", "tags": ["SecretService"]}], "method": "get"}

        
// ListSecrets returns the list of secrets.

func (a Secret) ListSecrets(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/secrets"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetSecret returns a secret given its ID.

func (a Secret) GetSecret(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/secrets/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
// CountSecrets returns the number of secrets.

func (a Secret) CountSecrets(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/secretscount"

    tools.GetResource(&a.Client, uriPath, args)

}                