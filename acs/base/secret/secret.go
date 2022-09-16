package secret

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Secret struct {
	Client client.Client
}
// ListSecrets returns the list of secrets.
func (a Secret) ListSecrets(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // ListResource()





}
// GetSecret returns a secret given its ID.
func (a Secret) GetSecret(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
// CountSecrets returns the number of secrets.
func (a Secret) CountSecrets(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // CountResource()





}
