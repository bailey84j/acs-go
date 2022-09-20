package sensorupgrade

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type SensorUpgrade struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "SensorUpgradeService_GetSensorUpgradeConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetSensorUpgradeConfigResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SensorUpgradeService"]}], "method": "get"}

        
func (a SensorUpgrade) GetSensorUpgradeConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/sensorupgrades/config"

    tools.GetResource(&a.Client, uriPath, args)

}                // debug: {"detail": [{"operationId": "SensorUpgradeService_TriggerSensorUpgrade", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SensorUpgradeService"]}, {"operationId": "SensorUpgradeService_UpdateSensorUpgradeConfig", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1UpdateSensorUpgradeConfigRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SensorUpgradeService"]}, {"operationId": "SensorUpgradeService_TriggerSensorCertRotation", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SensorUpgradeService"]}], "method": "post"}

        
func (a SensorUpgrade) TriggerSensorUpgrade(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/sensorupgrades/cluster/" + id + ""

    tools.PostResource(&a.Client, uriPath, args)

}
func (a SensorUpgrade) UpdateSensorUpgradeConfig(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/sensorupgrades/config"

    tools.PostResource(&a.Client, uriPath, args)

}
func (a SensorUpgrade) TriggerSensorCertRotation(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/sensorupgrades/rotateclustercerts/" + id + ""

    tools.PostResource(&a.Client, uriPath, args)

}    