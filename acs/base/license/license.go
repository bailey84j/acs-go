package license

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type License struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "LicenseService_GetActiveLicenseKey", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetActiveLicenseKeyResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["LicenseService"]}, {"operationId": "LicenseService_GetActiveLicenseExpiration", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetActiveLicenseExpirationResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["LicenseService"]}, {"operationId": "LicenseService_GetLicenses", "parameters": [{"in": "query", "name": "active", "required": false, "schema": {"type": "boolean"}}, {"explode": true, "in": "query", "name": "statuses", "required": false, "schema": {"items": {"enum": ["UNKNOWN", "VALID", "REVOKED", "NOT_YET_VALID", "EXPIRED", "OTHER"], "type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetLicensesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["LicenseService"]}], "method": "get"}

        
func (a License) GetActiveLicenseKey(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/licenses/activekey"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a License) GetActiveLicenseExpiration(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/licenses/expiration"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a License) GetLicenses(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("active-boolean,statuses-array",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/licenses/list"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                // debug: {"detail": [{"operationId": "LicenseService_AddLicense", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1AddLicenseRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1AddLicenseResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["LicenseService"]}], "method": "post"}

        
func (a License) AddLicense(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/licenses/add"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    