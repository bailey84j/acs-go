package apitoken

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type APIToken struct {
	Client client.Client
}
// GetAPITokens returns all the API tokens.
func (a APIToken) GetAPITokens(args map[string]interface{}) {

//  
// NOT Required [{'name': 'revoked', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("revoked_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
// GetAPIToken returns API token metadata for a given id.
func (a APIToken) GetAPIToken(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
