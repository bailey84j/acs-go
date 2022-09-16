package image

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Image struct {
	Client client.Client
}
// ListImages returns all the images.
func (a Image) ListImages(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // ListResource()





}
// InvalidateScanAndRegistryCaches removes the image metadata cache.
func (a Image) InvalidateScanAndRegistryCaches(args map[string]interface{}) {

//  
// NOT Required 




    // InvalidateScanAndRegistryCachesResource()





}
// GetImage returns the image given its ID.
func (a Image) GetImage(id string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'includeSnoozed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'stripDescription', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("includeSnoozed_boolean,stripDescription_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
// ListImages returns all the images.
func (a Image) CountImages(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("query_string,pagination_limit_integer,pagination_offset_integer,pagination_sortOption_field_string,pagination_sortOption_reversed_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // CountResource()





}
// GetWatchedImages returns the list of image names that are currently
being watched.
func (a Image) GetWatchedImages(args map[string]interface{}) {

//  
// NOT Required 




    // GetWatchedResource()





}
// DeleteImage removes the images based on a query
func (a Image) DeleteImages(args map[string]interface{}) {

//  
// NOT Required [{'name': 'query.query', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.limit', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.offset', 'in': 'query', 'required': False, 'schema': {'type': 'integer', 'format': 'int32'}}, {'name': 'query.pagination.sortOption.field', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}, {'name': 'query.pagination.sortOption.reversed', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'confirm', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

    ok := tools.CheckFieldsValid("query_query_string,query_pagination_limit_integer,query_pagination_offset_integer,query_pagination_sortOption_field_string,query_pagination_sortOption_reversed_boolean,confirm_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // DeleteResource()





}
// UnwatchImage marks an image name to no longer be watched.
It returns successfully if the image is no longer being watched
after the call, irrespective of whether the image was already being watched.
func (a Image) UnwatchImage(args map[string]interface{}) {

//  
// NOT Required [{'name': 'name', 'description': 'The name of the image to unwatch.\nShould match the name of a previously watched image.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("name_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // UnwatchResource()





}
