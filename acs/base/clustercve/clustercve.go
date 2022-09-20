package clustercve

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ClusterCVE struct {
	Client client.Client
}
                // debug: {"detail": [{"operationId": "ClusterCVEService_SuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1SuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SuppressCVE suppresses cluster cves.", "tags": ["ClusterCVEService"]}, {"operationId": "ClusterCVEService_UnsuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1UnsuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UnsuppressCVE unsuppresses cluster cves.", "tags": ["ClusterCVEService"]}], "method": "patch"}

        
// SuppressCVE suppresses cluster cves.

func (a ClusterCVE) SuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clustercves/suppress"

    tools.PatchResource(&a.Client, uriPath, args)

}
// UnsuppressCVE unsuppresses cluster cves.

func (a ClusterCVE) UnsuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/clustercves/unsuppress"

    tools.PatchResource(&a.Client, uriPath, args)

}        