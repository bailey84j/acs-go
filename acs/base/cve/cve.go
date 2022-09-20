package cve

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type CVE struct {
	Client client.Client
}
                // debug: {"detail": [{"operationId": "CVEService_SuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1SuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SuppressCVE suppresses cves.", "tags": ["CVEService"]}, {"operationId": "CVEService_UnsuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1UnsuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UnsuppressCVE unsuppresses cves.", "tags": ["CVEService"]}], "method": "patch"}

        
// SuppressCVE suppresses cves.

func (a CVE) SuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/cves/suppress"

    tools.PatchResource(&a.Client, uriPath, args)

}
// UnsuppressCVE unsuppresses cves.

func (a CVE) UnsuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/cves/unsuppress"

    tools.PatchResource(&a.Client, uriPath, args)

}        