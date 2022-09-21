package summary

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Summary struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "SummaryService_GetSummaryCounts", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SummaryCountsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SummaryService"]}], "method": "get"}

        
func (a Summary) GetSummaryCounts(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/summary/counts"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                