package imagecve

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ImageCVE struct {
	Client client.Client
}
                // debug: {"detail": [{"operationId": "ImageCVEService_SuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1SuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "SuppressCVE suppresses image cves.", "tags": ["ImageCVEService"]}, {"operationId": "ImageCVEService_UnsuppressCVEs", "requestBody": {"$ref": "#/components/requestBodies/v1UnsuppressCVERequest"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UnsuppressCVE unsuppresses image cves.", "tags": ["ImageCVEService"]}], "method": "patch"}

        
// SuppressCVE suppresses image cves.

func (a ImageCVE) SuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imagecves/suppress"

    tools.PatchResource(&a.Client, uriPath, args)

}
// UnsuppressCVE unsuppresses image cves.

func (a ImageCVE) UnsuppressCVEs(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/imagecves/unsuppress"

    tools.PatchResource(&a.Client, uriPath, args)

}        