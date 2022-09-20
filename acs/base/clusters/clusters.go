package clusters

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Clusters struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "ClustersService_GetClusterDefaultValues", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ClusterDefaultsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}, {"operationId": "ClustersService_GetClusters", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ClustersList"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}, {"operationId": "ClustersService_GetKernelSupportAvailable", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1KernelSupportAvailableResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetKernelSupportAvailable is deprecated in favor of GetClusterDefaultValues.", "tags": ["ClustersService"]}, {"operationId": "ClustersService_GetCluster", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ClusterResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}], "method": "get"}

        
func (a Clusters) GetClusterDefaultValues(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/cluster-defaults"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Clusters) GetClusters(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query-string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/clusters"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetKernelSupportAvailable is deprecated in favor of GetClusterDefaultValues.

func (a Clusters) GetKernelSupportAvailable(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clusters-env/kernel-support-available"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Clusters) GetCluster(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clusters/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "ClustersService_DeleteCluster", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}], "method": "delete"}

        
func (a Clusters) DeleteCluster(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clusters/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}            // debug: {"detail": [{"operationId": "ClustersService_PostCluster", "requestBody": {"$ref": "#/components/requestBodies/storageCluster"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ClusterResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}], "method": "post"}

        
func (a Clusters) PostCluster(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clusters"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "ClustersService_PutCluster", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storageCluster"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ClusterResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ClustersService"]}], "method": "put"}

        
func (a Clusters) PutCluster(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clusters/" + id + ""

    tools.PutResource(&a.Client, uriPath, args)

}