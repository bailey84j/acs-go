package search

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Search struct {
	Client client.Client
}
        // debug: {"detail": [{"operationId": "SearchService_Search", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"explode": true, "in": "query", "name": "categories", "required": false, "schema": {"items": {"enum": ["SEARCH_UNSET", "ALERTS", "IMAGES", "IMAGE_COMPONENTS", "IMAGE_VULN_EDGE", "IMAGE_COMPONENT_EDGE", "POLICIES", "DEPLOYMENTS", "ACTIVE_COMPONENT", "PODS", "SECRETS", "PROCESS_INDICATORS", "COMPLIANCE", "CLUSTERS", "NAMESPACES", "NODES", "NODE_COMPONENTS", "NODE_VULN_EDGE", "NODE_COMPONENT_EDGE", "NODE_COMPONENT_CVE_EDGE", "COMPLIANCE_STANDARD", "COMPLIANCE_CONTROL_GROUP", "COMPLIANCE_CONTROL", "SERVICE_ACCOUNTS", "ROLES", "ROLEBINDINGS", "REPORT_CONFIGURATIONS", "PROCESS_BASELINES", "SUBJECTS", "RISKS", "VULNERABILITIES", "CLUSTER_VULNERABILITIES", "IMAGE_VULNERABILITIES", "NODE_VULNERABILITIES", "COMPONENT_VULN_EDGE", "CLUSTER_VULN_EDGE", "NETWORK_ENTITY", "VULN_REQUEST", "NETWORK_BASELINE", "NETWORK_POLICIES", "PROCESS_BASELINE_RESULTS", "COMPLIANCE_METADATA", "COMPLIANCE_RESULTS", "COMPLIANCE_DOMAIN", "CLUSTER_HEALTH", "POLICY_CATEGORIES"], "type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SearchResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SearchService"]}, {"operationId": "SearchService_Autocomplete", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"explode": true, "in": "query", "name": "categories", "required": false, "schema": {"items": {"enum": ["SEARCH_UNSET", "ALERTS", "IMAGES", "IMAGE_COMPONENTS", "IMAGE_VULN_EDGE", "IMAGE_COMPONENT_EDGE", "POLICIES", "DEPLOYMENTS", "ACTIVE_COMPONENT", "PODS", "SECRETS", "PROCESS_INDICATORS", "COMPLIANCE", "CLUSTERS", "NAMESPACES", "NODES", "NODE_COMPONENTS", "NODE_VULN_EDGE", "NODE_COMPONENT_EDGE", "NODE_COMPONENT_CVE_EDGE", "COMPLIANCE_STANDARD", "COMPLIANCE_CONTROL_GROUP", "COMPLIANCE_CONTROL", "SERVICE_ACCOUNTS", "ROLES", "ROLEBINDINGS", "REPORT_CONFIGURATIONS", "PROCESS_BASELINES", "SUBJECTS", "RISKS", "VULNERABILITIES", "CLUSTER_VULNERABILITIES", "IMAGE_VULNERABILITIES", "NODE_VULNERABILITIES", "COMPONENT_VULN_EDGE", "CLUSTER_VULN_EDGE", "NETWORK_ENTITY", "VULN_REQUEST", "NETWORK_BASELINE", "NETWORK_POLICIES", "PROCESS_BASELINE_RESULTS", "COMPLIANCE_METADATA", "COMPLIANCE_RESULTS", "COMPLIANCE_DOMAIN", "CLUSTER_HEALTH", "POLICY_CATEGORIES"], "type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1AutocompleteResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SearchService"]}, {"operationId": "SearchService_Options", "parameters": [{"explode": true, "in": "query", "name": "categories", "required": false, "schema": {"items": {"enum": ["SEARCH_UNSET", "ALERTS", "IMAGES", "IMAGE_COMPONENTS", "IMAGE_VULN_EDGE", "IMAGE_COMPONENT_EDGE", "POLICIES", "DEPLOYMENTS", "ACTIVE_COMPONENT", "PODS", "SECRETS", "PROCESS_INDICATORS", "COMPLIANCE", "CLUSTERS", "NAMESPACES", "NODES", "NODE_COMPONENTS", "NODE_VULN_EDGE", "NODE_COMPONENT_EDGE", "NODE_COMPONENT_CVE_EDGE", "COMPLIANCE_STANDARD", "COMPLIANCE_CONTROL_GROUP", "COMPLIANCE_CONTROL", "SERVICE_ACCOUNTS", "ROLES", "ROLEBINDINGS", "REPORT_CONFIGURATIONS", "PROCESS_BASELINES", "SUBJECTS", "RISKS", "VULNERABILITIES", "CLUSTER_VULNERABILITIES", "IMAGE_VULNERABILITIES", "NODE_VULNERABILITIES", "COMPONENT_VULN_EDGE", "CLUSTER_VULN_EDGE", "NETWORK_ENTITY", "VULN_REQUEST", "NETWORK_BASELINE", "NETWORK_POLICIES", "PROCESS_BASELINE_RESULTS", "COMPLIANCE_METADATA", "COMPLIANCE_RESULTS", "COMPLIANCE_DOMAIN", "CLUSTER_HEALTH", "POLICY_CATEGORIES"], "type": "string"}, "type": "array"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1SearchOptionsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "tags": ["SearchService"]}], "method": "get"}

        
func (a Search) Search(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/search"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Search) Autocomplete(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("query_string,categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/search/autocomplete"

    tools.GetResource(&a.Client, uriPath, args)

}
func (a Search) Options(args map[string]interface{}) {

fmt.Printf("Running  Vaidation Failed")
    ok := tools.CheckFieldsValid("categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }
uriPath := "/v1/search/metadata/options"

    tools.GetResource(&a.Client, uriPath, args)

}                