package nodecve

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type NodeCVE struct {
	Client client.Client
}
                // debug: {"detail": [{"operationId": "NodeCVEService_SuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1SuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SuppressCVE suppresses node cves.", "tags": ["NodeCVEService"]}, {"operationId": "NodeCVEService_UnsuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1UnsuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UnsuppressCVE unsuppresses node cves.", "tags": ["NodeCVEService"]}], "method": "patch"}

        
// SuppressCVE suppresses node cves.

func (a NodeCVE) SuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/nodecves/suppress"

    tools.PatchResource(&a.Client, uriPath, args)

}
// UnsuppressCVE unsuppresses node cves.

func (a NodeCVE) UnsuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/nodecves/unsuppress"

    tools.PatchResource(&a.Client, uriPath, args)

}        