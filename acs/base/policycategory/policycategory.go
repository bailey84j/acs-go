package policycategory

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type PolicyCategory struct {
	Client client.Client
}
// GetPolicyCategories returns the list of policy categories
func (a PolicyCategory) GetPolicyCategories(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// GetPolicyCategoriesResource()





}
// GetPolicyCategory returns the requested policy category by ID.
func (a PolicyCategory) GetPolicyCategory(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetResource()





}
// DeletePolicyCategory removes the given policy category.
func (a PolicyCategory) DeletePolicyCategory(id string,args map[string]interface{}) {

//  
// NOT Required []




// DeleteResource()





}
