package role

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Role struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "RoleService_GetMyPermissions", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetPermissionsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_ListPermissionSets", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListPermissionSetsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_GetPermissionSet", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePermissionSet"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_GetResources", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetResourcesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_GetRoles", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetRolesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_GetRole", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageRole"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_ListSimpleAccessScopes", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListSimpleAccessScopesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_GetSimpleAccessScope", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageSimpleAccessScope"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}], "method": "get"}

        
func (a Role) GetMyPermissions(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/mypermissions"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) ListPermissionSets(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/permissionsets"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) GetPermissionSet(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/permissionsets/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) GetResources(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/resources"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) GetRoles(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/roles"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) GetRole(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/roles/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) ListSimpleAccessScopes(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/simpleaccessscopes"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) GetSimpleAccessScope(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/simpleaccessscopes/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "RoleService_DeletePermissionSet", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_DeleteRole", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_DeleteSimpleAccessScope", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}], "method": "delete"}

        
func (a Role) DeletePermissionSet(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/permissionsets/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}
func (a Role) DeleteRole(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/roles/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}
func (a Role) DeleteSimpleAccessScope(id string,args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/simpleaccessscopes/" + id + ""

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"description": "Returns effective access scope based on the rules in the request. Does\nnot persist anything; not idempotent due to possible changes to clusters\nand namespaces. POST is chosen due to potentially large payload.\n\nThere are advantages in both keeping the response slim and detailed. If\nonly IDs of selected clusters and namespaces are included, response\nlatency and processing time are lower but the caller shall overlay the\nresponse with its view of the world which is susceptible to consistency\nissues. Listing all clusters and namespaces with related metadata is\nconvenient for the caller but bloat the message with secondary data.\n\nWe let the caller decide what level of detail they would like to have:\n\n  * Minimal, when only roots of included subtrees are listed by their\n    IDs. Clusters can be either INCLUDED (its namespaces are included but\n    are not listed) or PARTIAL (at least one namespace is explicitly\n    included). Namespaces can only be INCLUDED.\n\n  * Standard [default], when all known clusters and namespaces are listed\n    with their IDs and names. Clusters can be INCLUDED (all its\n    namespaces are explicitly listed as INCLUDED), PARTIAL (all its\n    namespaces are explicitly listed, some as INCLUDED and some as\n    EXCLUDED), and EXCLUDED (all its namespaces are explicitly listed as\n    EXCLUDED). Namespaces can be either INCLUDED or EXCLUDED.\n\n  * High, when every cluster and namespace is augmented with metadata.", "operationId": "RoleService_ComputeEffectiveAccessScope", "parameters": [{"in": "query", "name": "detail", "required": false, "schema": {"default": "STANDARD", "enum": ["STANDARD", "MINIMAL", "HIGH"], "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/ComputeEffectiveAccessScopeRequestPayload"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageEffectiveAccessScope"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ComputeEffectiveAccessScope", "tags": ["RoleService"]}, {"description": "PermissionSet.id is disallowed in request and set in response.", "operationId": "RoleService_PostPermissionSet", "requestBody": {"$ref": "#/components/requestBodies/storagePermissionSet"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePermissionSet"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostPermissionSet", "tags": ["RoleService"]}, {"operationId": "RoleService_CreateRole", "parameters": [{"in": "path", "name": "name", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageRole"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"description": "SimpleAccessScope.id is disallowed in request and set in response.", "operationId": "RoleService_PostSimpleAccessScope", "requestBody": {"$ref": "#/components/requestBodies/storageSimpleAccessScope"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageSimpleAccessScope"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostSimpleAccessScope", "tags": ["RoleService"]}], "method": "post"}

        
// ComputeEffectiveAccessScope

func (a Role) ComputeEffectiveAccessScope(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("detail-string",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/computeeffectiveaccessscope"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// PostPermissionSet

func (a Role) PostPermissionSet(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/permissionsets"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) CreateRole(name string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/roles/" + name + ""

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// PostSimpleAccessScope

func (a Role) PostSimpleAccessScope(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/simpleaccessscopes"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "RoleService_PutPermissionSet", "parameters": [{"description": "id is generated and cannot be changed.", "in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storagePermissionSet"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_UpdateRole", "parameters": [{"description": "`name` and `description` are provided by the user and can be changed.", "in": "path", "name": "name", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageRole"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}, {"operationId": "RoleService_PutSimpleAccessScope", "parameters": [{"description": "`id` is generated and cannot be changed.", "in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageSimpleAccessScope"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["RoleService"]}], "method": "put"}

        
func (a Role) PutPermissionSet(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/permissionsets/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) UpdateRole(name string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/roles/" + name + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a Role) PutSimpleAccessScope(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/simpleaccessscopes/" + id + ""

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}