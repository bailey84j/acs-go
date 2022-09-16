package compliance

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Compliance struct {
	Client client.Client
}
func (a Compliance) GetAggregatedResults(args map[string]interface{}) {

//  
// NOT Required [{'name': 'groupBy', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string', 'enum': ['UNKNOWN', 'STANDARD', 'CLUSTER', 'CATEGORY', 'CONTROL', 'NAMESPACE', 'NODE', 'DEPLOYMENT', 'CHECK']}}}, {'name': 'unit', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'enum': ['UNKNOWN', 'STANDARD', 'CLUSTER', 'CATEGORY', 'CONTROL', 'NAMESPACE', 'NODE', 'DEPLOYMENT', 'CHECK'], 'default': 'UNKNOWN'}}, {'name': 'where.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'where.pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'where.pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'where.pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'where.pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("groupBy_array,unit_string,where_query_string,where_pagination_limit_integer,where_pagination_offset_integer,where_pagination_sortOption_field_string,where_pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetAggregatedResultsResource()





}
func (a Compliance) GetRunResults(args map[string]interface{}) {

//  
// NOT Required [{'name': 'clusterId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'standardId', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'runId', 'description': 'Specifies the run ID for which to return results. If empty, the most recent run is returned.\nCAVEAT: Setting this field circumvents the results cache on the server-side, which may lead to significantly\n        increased memory pressure and decreased performance.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("clusterId_string,standardId_string,runId_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetRunResultsResource()





}
func (a Compliance) GetStandards(args map[string]interface{}) {

//  
// NOT Required 




    // GetStandardsResource()





}
func (a Compliance) GetStandard(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetStandardResource()





}
