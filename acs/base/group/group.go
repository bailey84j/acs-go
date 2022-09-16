package group

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Group struct {
	Client client.Client
}
func (a Group) GetGroup(args map[string]interface{}) {

//  
// NOT Required [{'name': 'id', 'description': 'Unique identifier for group properties and respectively the group.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'authProviderId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'value', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("id_string,authProviderId_string,key_string,value_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
func (a Group) GetGroups(args map[string]interface{}) {

//  
// NOT Required [{'name': 'authProviderId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'value', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'id', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("authProviderId_string,key_string,value_string,id_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
func (a Group) DeleteGroup(args map[string]interface{}) {

//  
// NOT Required [{'name': 'id', 'description': 'Unique identifier for group properties and respectively the group.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'authProviderId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'key', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'value', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("id_string,authProviderId_string,key_string,value_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // DeleteResource()





}
