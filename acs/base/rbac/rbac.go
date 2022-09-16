package rbac

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Rbac struct {
	Client client.Client
}
func (a Rbac) ListRoleBindings(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// ListRoleBindingsResource()





}
func (a Rbac) GetRoleBinding(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetRoleBindingResource()





}
func (a Rbac) ListRoles(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// ListRolesResource()





}
func (a Rbac) GetRole(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetRoleResource()





}
// Subjects served from this API are Groups and Users only.
Id in this case is the Name field, since for users and groups, that is unique, and subjects do not have IDs.
func (a Rbac) GetSubject(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetSubjectResource()





}
func (a Rbac) ListSubjects(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// ListSubjectsResource()





}
