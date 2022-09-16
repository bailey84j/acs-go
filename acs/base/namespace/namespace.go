package namespace

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Namespace struct {
	Client client.Client
}
func (a Namespace) GetNamespaces(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("query_query_string,query_pagination_limit_integer,query_pagination_offset_integer,query_pagination_sortOption_field_string,query_pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
func (a Namespace) GetNamespace(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
