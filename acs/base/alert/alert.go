package alert

import (
	"fmt"

	client "github.com/bailey84j/acs-go/acs/client"
	tools "github.com/bailey84j/acs-go/acs/tools"
)

type Alert struct {
	Client client.Client
}

// List returns the slim list version of the alerts.
func (a Alert) ListAlerts(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

	ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// ListResource()

}

// GetAlertsCounts returns the number of alerts in the requested cluster or category.
func (a Alert) GetAlertsCounts(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'request.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'request.pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'request.pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'request.pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'request.pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'groupBy', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'enum': ['UNSET', 'CATEGORY', 'CLUSTER'], 'default': 'UNSET'}}]

	ok := tools.CheckFieldsValid("request_query_string,request_pagination_limit_integer,request_pagination_offset_integer,request_pagination_sortOption_field_string,request_pagination_sortOption_reversed_boolean,groupBy_string", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// GetResource()

}

// GetAlertsGroup returns alerts grouped by policy.
func (a Alert) GetAlertsGroup(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

	ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// GetResource()

}

// GetAlertTimeseries returns the alerts sorted by time.
func (a Alert) GetAlertTimeseries(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

	ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// GetResource()

}

// GetAlert returns the alert given its id.
func (a Alert) GetAlert(id string, args map[string]interface{}) {

	//
	// NOT Required []

	// GetResource()

}

// CountAlerts counts how many alerts match the get request.
func (a Alert) CountAlerts(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

	ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// CountResource()

}
func (a Alert) DeleteAlerts(args map[string]interface{}) {

	//
	// NOT Required [{'name': 'query.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'confirm', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

	ok := tools.CheckFieldsValid("query_query_string,query_pagination_limit_integer,query_pagination_offset_integer,query_pagination_sortOption_field_string,query_pagination_sortOption_reversed_boolean,confirm_boolean", args)
	if !ok {
		fmt.Printf("Variable Vaidation Failed")
	}

	// DeleteResource()

}
