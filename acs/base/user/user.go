package user

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type User struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "UserService_GetUsers", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetUsersResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["UserService"]}, {"operationId": "UserService_GetUser", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageUser"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["UserService"]}, {"operationId": "UserService_GetUsersAttributes", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetUsersAttributesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["UserService"]}], "method": "get"}

        
func (a User) GetUsers(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/users"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a User) GetUser(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/users/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
func (a User) GetUsersAttributes(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/usersattributes"

    tools.GetResource(&a.Client, uriPath, args)

}                