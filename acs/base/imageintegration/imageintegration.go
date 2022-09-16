package imageintegration

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type ImageIntegration struct {
	Client client.Client
}
// GetImageIntegrations returns all image integrations that match the request filters.
func (a ImageIntegration) GetImageIntegrations(args map[string]interface{}) {

//  
// NOT Required [{'name': 'name', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'cluster', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("name_string,cluster_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
// GetImageIntegration returns the image integration given its ID.
func (a ImageIntegration) GetImageIntegration(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
// DeleteImageIntegration removes a image integration given its ID.
func (a ImageIntegration) DeleteImageIntegration(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteResource()





}
