package namespace

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Namespace struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "NamespaceService_GetNamespaces", "parameters": [{"in": "query", "name": "query.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetNamespacesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NamespaceService"]}, {"operationId": "NamespaceService_GetNamespace", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Namespace"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["NamespaceService"]}], "method": "get"}

        
func (a Namespace) GetNamespaces(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_query-string,query_pagination_limit-integer,query_pagination_offset-integer,query_pagination_sortOption_field-string,query_pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/namespaces"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Namespace) GetNamespace(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/namespaces/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}                