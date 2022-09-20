package ping

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Ping struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "PingService_Ping", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PongMessage"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["PingService"]}], "method": "get"}

        
func (a Ping) Ping(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/ping"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                