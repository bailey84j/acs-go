package probeupload

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type ProbeUpload struct {
	Client *client.Client
}
                    // debug: {"detail": [{"operationId": "ProbeUploadService_GetExistingProbes", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetExistingProbesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["ProbeUploadService"]}], "method": "post"}

        
func (a ProbeUpload) GetExistingProbes(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/probeupload/getexisting"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    