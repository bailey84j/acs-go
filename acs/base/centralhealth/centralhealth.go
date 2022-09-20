package centralhealth

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type CentralHealth struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "CentralHealthService_GetUpgradeStatus", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetUpgradeStatusResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["CentralHealthService"]}], "method": "get"}

        
func (a CentralHealth) GetUpgradeStatus(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/centralhealth/upgradestatus"

    tools.GetResource(&a.Client, uriPath, args)

}                