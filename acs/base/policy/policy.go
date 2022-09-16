package policy

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Policy struct {
	Client client.Client
}
// ListPolicies returns the list of policies.
func (a Policy) ListPolicies(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// ListPoliciesResource()





}
func (a Policy) QueryDryRunJobStatus(jobId string,args map[string]interface{}) {

//  
// NOT Required []




// QueryDryRunJobStatusResource()





}
// GetPolicy returns the requested policy by ID.
func (a Policy) GetPolicy(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetResource()





}
// GetMitreVectorsForPolicy returns the requested policy by ID.
func (a Policy) GetPolicyMitreVectors(id string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'options.excludePolicy', 'description': 'If set to true, policy is excluded from the response.', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("options_excludePolicy_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// GetResource()





}
// GetPolicyCategories returns the policy categories.
func (a Policy) GetPolicyCategories(args map[string]interface{}) {

//  
// NOT Required 




// GetResource()





}
func (a Policy) CancelDryRunJob(jobId string,args map[string]interface{}) {

//  
// NOT Required []




// CancelDryRunJobResource()





}
// DeletePolicy removes a policy by ID.
func (a Policy) DeletePolicy(id string,args map[string]interface{}) {

//  
// NOT Required []




// DeleteResource()





}
// DeletePolicyCategory removes the given policy category.
func (a Policy) DeletePolicyCategory(category string,args map[string]interface{}) {

//  
// NOT Required []




// DeleteResource()





}
