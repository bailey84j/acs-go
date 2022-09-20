package auth

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Auth struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "AuthService_GetAuthStatus", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1AuthStatus"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["AuthService"]}], "method": "get"}

        
func (a Auth) GetAuthStatus(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/auth/status"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                