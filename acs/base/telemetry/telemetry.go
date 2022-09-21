package telemetry

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Telemetry struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "TelemetryService_GetTelemetryConfiguration", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageTelemetryConfiguration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["TelemetryService"]}], "method": "get"}

        
func (a Telemetry) GetTelemetryConfiguration(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/telemetry/configure"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                    // debug: {"detail": [{"operationId": "TelemetryService_ConfigureTelemetry", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ConfigureTelemetryRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageTelemetryConfiguration"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["TelemetryService"]}], "method": "put"}

        
func (a Telemetry) ConfigureTelemetry(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/telemetry/configure"

    msi, err := tools.PutResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}