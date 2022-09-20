package summary

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Summary struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "SummaryService_GetSummaryCounts", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SummaryCountsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SummaryService"]}], "method": "get"}

        
func (a Summary) GetSummaryCounts(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/summary/counts"

    tools.GetResource(&a.Client, uriPath, args)

}                