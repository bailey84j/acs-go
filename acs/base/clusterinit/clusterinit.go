package clusterinit

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ClusterInit struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ClusterInitService_GetCAConfig", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetCAConfigResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClusterInitService"]}, {"operationId": "ClusterInitService_GetInitBundles", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InitBundleMetasResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClusterInitService"]}], "method": "get"}

        
func (a ClusterInit) GetCAConfig(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/cluster-init/ca-config"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
func (a ClusterInit) GetInitBundles(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/cluster-init/init-bundles"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}            // debug: {"detail": [{"operationId": "ClusterInitService_RevokeInitBundle", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InitBundleRevokeRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InitBundleRevokeResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "RevokeInitBundle deletes cluster init bundle. If this operation impacts any cluster\nthen its ID should be included in request.\nIf confirm_impacted_clusters_ids does not match with current impacted clusters\nthen request will fail with error that includes all impacted clusters.", "tags": ["ClusterInitService"]}], "method": "patch"}

        
// RevokeInitBundle deletes cluster init bundle. If this operation impacts any cluster then its ID should be included in request. If confirm_impacted_clusters_ids does not match with current impacted clusters then request will fail with error that includes all impacted clusters.

func (a ClusterInit) RevokeInitBundle(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/cluster-init/init-bundles/revoke"

    msi, err := tools.PatchResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ClusterInitService_GenerateInitBundle", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InitBundleGenRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1InitBundleGenResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClusterInitService"]}], "method": "post"}

        
func (a ClusterInit) GenerateInitBundle(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/cluster-init/init-bundles"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    