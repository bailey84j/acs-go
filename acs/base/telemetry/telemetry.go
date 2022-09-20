package telemetry

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Telemetry struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "TelemetryService_GetTelemetryConfiguration", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageTelemetryConfiguration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["TelemetryService"]}], "method": "get"}

        
func (a Telemetry) GetTelemetryConfiguration(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/telemetry/configure"

    tools.GetResource(&a.Client, uriPath, args)

}                    // debug: {"detail": [{"operationId": "TelemetryService_ConfigureTelemetry", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ConfigureTelemetryRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageTelemetryConfiguration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["TelemetryService"]}], "method": "put"}

        
func (a Telemetry) ConfigureTelemetry(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/telemetry/configure"

    tools.PutResource(&a.Client, uriPath, args)

}