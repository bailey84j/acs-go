package search

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Search struct {
	Client client.Client
}
func (a Search) Search(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'categories', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string', 'enum': ['SEARCH_UNSET', 'ALERTS', 'IMAGES', 'IMAGE_COMPONENTS', 'IMAGE_VULN_EDGE', 'IMAGE_COMPONENT_EDGE', 'POLICIES', 'DEPLOYMENTS', 'ACTIVE_COMPONENT', 'PODS', 'SECRETS', 'PROCESS_INDICATORS', 'COMPLIANCE', 'CLUSTERS', 'NAMESPACES', 'NODES', 'NODE_COMPONENTS', 'NODE_VULN_EDGE', 'NODE_COMPONENT_EDGE', 'NODE_COMPONENT_CVE_EDGE', 'COMPLIANCE_STANDARD', 'COMPLIANCE_CONTROL_GROUP', 'COMPLIANCE_CONTROL', 'SERVICE_ACCOUNTS', 'ROLES', 'ROLEBINDINGS', 'REPORT_CONFIGURATIONS', 'PROCESS_BASELINES', 'SUBJECTS', 'RISKS', 'VULNERABILITIES', 'CLUSTER_VULNERABILITIES', 'IMAGE_VULNERABILITIES', 'NODE_VULNERABILITIES', 'COMPONENT_VULN_EDGE', 'CLUSTER_VULN_EDGE', 'NETWORK_ENTITY', 'VULN_REQUEST', 'NETWORK_BASELINE', 'NETWORK_POLICIES', 'PROCESS_BASELINE_RESULTS', 'COMPLIANCE_METADATA', 'COMPLIANCE_RESULTS', 'COMPLIANCE_DOMAIN', 'CLUSTER_HEALTH', 'POLICY_CATEGORIES']}}}]

ok := tools.CheckFieldsValid("query_string,categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // Resource()





}
func (a Search) Autocomplete(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'categories', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string', 'enum': ['SEARCH_UNSET', 'ALERTS', 'IMAGES', 'IMAGE_COMPONENTS', 'IMAGE_VULN_EDGE', 'IMAGE_COMPONENT_EDGE', 'POLICIES', 'DEPLOYMENTS', 'ACTIVE_COMPONENT', 'PODS', 'SECRETS', 'PROCESS_INDICATORS', 'COMPLIANCE', 'CLUSTERS', 'NAMESPACES', 'NODES', 'NODE_COMPONENTS', 'NODE_VULN_EDGE', 'NODE_COMPONENT_EDGE', 'NODE_COMPONENT_CVE_EDGE', 'COMPLIANCE_STANDARD', 'COMPLIANCE_CONTROL_GROUP', 'COMPLIANCE_CONTROL', 'SERVICE_ACCOUNTS', 'ROLES', 'ROLEBINDINGS', 'REPORT_CONFIGURATIONS', 'PROCESS_BASELINES', 'SUBJECTS', 'RISKS', 'VULNERABILITIES', 'CLUSTER_VULNERABILITIES', 'IMAGE_VULNERABILITIES', 'NODE_VULNERABILITIES', 'COMPONENT_VULN_EDGE', 'CLUSTER_VULN_EDGE', 'NETWORK_ENTITY', 'VULN_REQUEST', 'NETWORK_BASELINE', 'NETWORK_POLICIES', 'PROCESS_BASELINE_RESULTS', 'COMPLIANCE_METADATA', 'COMPLIANCE_RESULTS', 'COMPLIANCE_DOMAIN', 'CLUSTER_HEALTH', 'POLICY_CATEGORIES']}}}]

ok := tools.CheckFieldsValid("query_string,categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // AutocompleteResource()





}
func (a Search) Options(args map[string]interface{}) {

//  
// NOT Required [{'name': 'categories', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string', 'enum': ['SEARCH_UNSET', 'ALERTS', 'IMAGES', 'IMAGE_COMPONENTS', 'IMAGE_VULN_EDGE', 'IMAGE_COMPONENT_EDGE', 'POLICIES', 'DEPLOYMENTS', 'ACTIVE_COMPONENT', 'PODS', 'SECRETS', 'PROCESS_INDICATORS', 'COMPLIANCE', 'CLUSTERS', 'NAMESPACES', 'NODES', 'NODE_COMPONENTS', 'NODE_VULN_EDGE', 'NODE_COMPONENT_EDGE', 'NODE_COMPONENT_CVE_EDGE', 'COMPLIANCE_STANDARD', 'COMPLIANCE_CONTROL_GROUP', 'COMPLIANCE_CONTROL', 'SERVICE_ACCOUNTS', 'ROLES', 'ROLEBINDINGS', 'REPORT_CONFIGURATIONS', 'PROCESS_BASELINES', 'SUBJECTS', 'RISKS', 'VULNERABILITIES', 'CLUSTER_VULNERABILITIES', 'IMAGE_VULNERABILITIES', 'NODE_VULNERABILITIES', 'COMPONENT_VULN_EDGE', 'CLUSTER_VULN_EDGE', 'NETWORK_ENTITY', 'VULN_REQUEST', 'NETWORK_BASELINE', 'NETWORK_POLICIES', 'PROCESS_BASELINE_RESULTS', 'COMPLIANCE_METADATA', 'COMPLIANCE_RESULTS', 'COMPLIANCE_DOMAIN', 'CLUSTER_HEALTH', 'POLICY_CATEGORIES']}}}]

ok := tools.CheckFieldsValid("categories_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // OptionsResource()





}
