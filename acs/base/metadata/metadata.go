package metadata

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Metadata struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "MetadataService_GetMetadata", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Metadata"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["MetadataService"]}, {"operationId": "MetadataService_TLSChallenge", "parameters": [{"description": "generated challenge token by the service asking for TLS certs.", "in": "query", "name": "challengeToken", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1TLSChallengeResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["MetadataService"]}], "method": "get"}

        
func (a Metadata) GetMetadata(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/metadata"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Metadata) TLSChallenge(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("challengeToken_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/tls-challenge"

    tools.GetResource(&a.Client, uriPath, args)

}                