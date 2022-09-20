package featureflag

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type FeatureFlag struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "FeatureFlagService_GetFeatureFlags", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetFeatureFlagsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["FeatureFlagService"]}], "method": "get"}

        
func (a FeatureFlag) GetFeatureFlags(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/featureflags"

    tools.GetResource(&a.Client, uriPath, args)

}                