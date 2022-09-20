package policy

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Policy struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "PolicyService_ListPolicies", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListPoliciesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListPolicies returns the list of policies.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_QueryDryRunJobStatus", "parameters": [{"in": "path", "name": "jobId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DryRunJobStatusResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["PolicyService"]}, {"operationId": "PolicyService_GetPolicy", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePolicy"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetPolicy returns the requested policy by ID.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_GetPolicyMitreVectors", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}, {"description": "If set to true, policy is excluded from the response.", "in": "query", "name": "options.excludePolicy", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetPolicyMitreVectorsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetMitreVectorsForPolicy returns the requested policy by ID.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_GetPolicyCategories", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PolicyCategoriesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetPolicyCategories returns the policy categories.", "tags": ["PolicyService"]}], "method": "get"}

        
// ListPolicies returns the list of policies.

func (a Policy) ListPolicies(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/policies"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Policy) QueryDryRunJobStatus(jobId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/dryrunjob/" + jobId + ""

    tools.GetResource(&a.Client, uriPath, args)

}
// GetPolicy returns the requested policy by ID.

func (a Policy) GetPolicy(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/" + id + ""

    tools.GetResource(&a.Client, uriPath, args)

}
// GetMitreVectorsForPolicy returns the requested policy by ID.

func (a Policy) GetPolicyMitreVectors(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("options_excludePolicy-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/policies/" + id + "/mitrevectors"

    tools.GetResource(&a.Client, uriPath, args)

}
// GetPolicyCategories returns the policy categories.

func (a Policy) GetPolicyCategories(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policyCategories"

    tools.GetResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "PolicyService_CancelDryRunJob", "parameters": [{"in": "path", "name": "jobId", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["PolicyService"]}, {"operationId": "PolicyService_DeletePolicy", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeletePolicy removes a policy by ID.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_DeletePolicyCategory", "parameters": [{"in": "path", "name": "category", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeletePolicyCategory removes the given policy category.", "tags": ["PolicyService"]}], "method": "delete"}

        
func (a Policy) CancelDryRunJob(jobId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/dryrunjob/" + jobId + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}
// DeletePolicy removes a policy by ID.

func (a Policy) DeletePolicy(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/" + id + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}
// DeletePolicyCategory removes the given policy category.

func (a Policy) DeletePolicyCategory(category string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policyCategories/" + category + ""

    tools.DeleteResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "PolicyService_PatchPolicy", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PatchPolicyRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PatchPolicy edits an existing policy.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_EnableDisablePolicyNotification", "parameters": [{"in": "path", "name": "policyId", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1EnableDisablePolicyNotificationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "EnableDisablePolicyNotification enables or disables notifications for a policy by ID.", "tags": ["PolicyService"]}], "method": "patch"}

        
// PatchPolicy edits an existing policy.

func (a Policy) PatchPolicy(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/" + id + ""

    tools.PatchResource(&a.Client, uriPath, args)

}
// EnableDisablePolicyNotification enables or disables notifications for a policy by ID.

func (a Policy) EnableDisablePolicyNotification(policyId string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/" + policyId + "/notifiers"

    tools.PatchResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "PolicyService_PostPolicy", "parameters": [{"in": "query", "name": "enableStrictValidation", "required": false, "schema": {"type": "boolean"}}], "requestBody": {"$ref": "#/components/requestBodies/storagePolicy"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storagePolicy"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PostPolicy creates a new policy.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_DryRunPolicy", "requestBody": {"$ref": "#/components/requestBodies/storagePolicy"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DryRunResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DryRunPolicy evaluates the given policy and returns any alerts without creating the policy.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_SubmitDryRunPolicyJob", "requestBody": {"$ref": "#/components/requestBodies/storagePolicy"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1JobId"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["PolicyService"]}, {"operationId": "PolicyService_ExportPolicies", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ExportPoliciesRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageExportPoliciesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ExportPolicies takes a list of policy IDs and returns either the entire list of policies or an error message", "tags": ["PolicyService"]}, {"operationId": "PolicyService_PolicyFromSearch", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PolicyFromSearchRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1PolicyFromSearchResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["PolicyService"]}, {"operationId": "PolicyService_ImportPolicies", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ImportPoliciesRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ImportPoliciesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ImportPolicies accepts a list of Policies and returns a list of the policies which could not be imported", "tags": ["PolicyService"]}, {"operationId": "PolicyService_ReassessPolicies", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ReassessPolicies reevaluates all the policies.", "tags": ["PolicyService"]}], "method": "post"}

        
// PostPolicy creates a new policy.

func (a Policy) PostPolicy(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("enableStrictValidation-boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/policies"

    tools.PostResource(&a.Client, uriPath, args)

}
// DryRunPolicy evaluates the given policy and returns any alerts without creating the policy.

func (a Policy) DryRunPolicy(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/dryrun"

    tools.PostResource(&a.Client, uriPath, args)

}
func (a Policy) SubmitDryRunPolicyJob(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/dryrunjob"

    tools.PostResource(&a.Client, uriPath, args)

}
// ExportPolicies takes a list of policy IDs and returns either the entire list of policies or an error message

func (a Policy) ExportPolicies(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/export"

    tools.PostResource(&a.Client, uriPath, args)

}
func (a Policy) PolicyFromSearch(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/from-search"

    tools.PostResource(&a.Client, uriPath, args)

}
// ImportPolicies accepts a list of Policies and returns a list of the policies which could not be imported

func (a Policy) ImportPolicies(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/import"

    tools.PostResource(&a.Client, uriPath, args)

}
// ReassessPolicies reevaluates all the policies.

func (a Policy) ReassessPolicies(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/reassess"

    tools.PostResource(&a.Client, uriPath, args)

}        // debug: {"detail": [{"operationId": "PolicyService_PutPolicy", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "requestBody": {"$ref": "#/components/requestBodies/storagePolicy"}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "PutPolicy modifies an existing policy.", "tags": ["PolicyService"]}, {"operationId": "PolicyService_RenamePolicyCategory", "parameters": [{"in": "path", "name": "oldCategory", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1RenamePolicyCategoryRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "RenamePolicyCategory renames the given policy category.", "tags": ["PolicyService"]}], "method": "put"}

        
// PutPolicy modifies an existing policy.

func (a Policy) PutPolicy(id string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policies/" + id + ""

    tools.PutResource(&a.Client, uriPath, args)

}
// RenamePolicyCategory renames the given policy category.

func (a Policy) RenamePolicyCategory(oldCategory string,args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")

uriPath := "/v1/policyCategories/" + oldCategory + ""

    tools.PutResource(&a.Client, uriPath, args)

}