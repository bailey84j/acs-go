package process

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Process struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ProcessService_GetProcessesByDeployment", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetProcessesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetProcessesByDeployment returns the processes executed in the given deployment.", "tags": ["ProcessService"]}, {"operationId": "ProcessService_GetGroupedProcessByDeployment", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetGroupedProcessesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetGroupedProcessByDeployment returns all the processes executed grouped by deployment.", "tags": ["ProcessService"]}, {"operationId": "ProcessService_GetGroupedProcessByDeploymentAndContainer", "parameters": [{"in": "path", "name": "deploymentId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetGroupedProcessesWithContainerResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetGroupedProcessByDeploymentAndContainer returns all the processes executed grouped by deployment and container.", "tags": ["ProcessService"]}], "method": "get"}

        
// GetProcessesByDeployment returns the processes executed in the given deployment.

func (a Process) GetProcessesByDeployment(deploymentId string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/processes/deployment/" + deploymentId + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetGroupedProcessByDeployment returns all the processes executed grouped by deployment.

func (a Process) GetGroupedProcessByDeployment(deploymentId string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/processes/deployment/" + deploymentId + "/grouped"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetGroupedProcessByDeploymentAndContainer returns all the processes executed grouped by deployment and container.

func (a Process) GetGroupedProcessByDeploymentAndContainer(deploymentId string,args map[string]interface{}) (map[string]interface{}, error) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/processes/deployment/" + deploymentId + "/grouped/container"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                