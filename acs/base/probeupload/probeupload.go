package probeupload

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ProbeUpload struct {
	Client client.Client
}
                    // debug: {"detail": [{"operationId": "ProbeUploadService_GetExistingProbes", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetExistingProbesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ProbeUploadService"]}], "method": "post"}

        
func (a ProbeUpload) GetExistingProbes(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/probeupload/getexisting"

    tools.PostResource(&a.Client, uriPath, args)

}    