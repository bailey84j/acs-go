package authprovider

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type AuthProvider struct {
	Client client.Client
}
func (a AuthProvider) GetAuthProviders(args map[string]interface{}) {

//  
// NOT Required [{'name': 'name', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'type', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("name_string,type_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
func (a AuthProvider) GetAuthProvider(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
func (a AuthProvider) ListAvailableProviderTypes(args map[string]interface{}) {

//  
// NOT Required 




    // ListAvailableProviderTypesResource()





}
func (a AuthProvider) GetLoginAuthProviders(args map[string]interface{}) {

//  
// NOT Required 




    // GetLoginResource()





}
func (a AuthProvider) DeleteAuthProvider(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteResource()





}
