package process

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Process struct {
	Client client.Client
}
// GetProcessesByDeployment returns the processes executed in the given deployment.
func (a Process) GetProcessesByDeployment(deploymentId string,args map[string]interface{}) {

//  
// NOT Required []




// GetResource()





}
// GetGroupedProcessByDeployment returns all the processes executed grouped by deployment.
func (a Process) GetGroupedProcessByDeployment(deploymentId string,args map[string]interface{}) {

//  
// NOT Required []




// GetGroupedResource()





}
// GetGroupedProcessByDeploymentAndContainer returns all the processes executed grouped by deployment and container.
func (a Process) GetGroupedProcessByDeploymentAndContainer(deploymentId string,args map[string]interface{}) {

//  
// NOT Required []




// GetGroupedResource()





}
