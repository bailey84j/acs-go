package detection

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Detection struct {
	Client *client.Client
}
                    // debug: {"detail": [{"operationId": "DetectionService_DetectBuildTime", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1BuildDetectionRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1BuildDetectionResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DetectBuildTime checks if any images violate build time policies.", "tags": ["DetectionService"]}, {"operationId": "DetectionService_DetectDeployTime", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeployDetectionRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeployDetectionResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DetectDeployTime checks if any deployments violate deploy time policies.", "tags": ["DetectionService"]}, {"operationId": "DetectionService_DetectDeployTimeFromYAML", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeployYAMLDetectionRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeployDetectionResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DetectDeployTimeFromYAML checks if the given deployment yaml violates any deploy time policies.", "tags": ["DetectionService"]}], "method": "post"}

        
// DetectBuildTime checks if any images violate build time policies.

func (a Detection) DetectBuildTime(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/detect/build"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// DetectDeployTime checks if any deployments violate deploy time policies.

func (a Detection) DetectDeployTime(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/detect/deploy"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// DetectDeployTimeFromYAML checks if the given deployment yaml violates any deploy time policies.

func (a Detection) DetectDeployTimeFromYAML(args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/detect/deploy/yaml"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    