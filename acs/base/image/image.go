package image

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type Image struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "ImageService_ListImages", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListImagesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListImages returns all the images.", "tags": ["ImageService"]}, {"operationId": "ImageService_InvalidateScanAndRegistryCaches", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "InvalidateScanAndRegistryCaches removes the image metadata cache.", "tags": ["ImageService"]}, {"operationId": "ImageService_GetImage", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}, {"in": "query", "name": "includeSnoozed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "stripDescription", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageImage"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetImage returns the image given its ID.", "tags": ["ImageService"]}, {"operationId": "ImageService_CountImages", "parameters": [{"in": "query", "name": "query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1CountImagesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListImages returns all the images.", "tags": ["ImageService"]}, {"operationId": "ImageService_GetWatchedImages", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetWatchedImagesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetWatchedImages returns the list of image names that are currently\nbeing watched.", "tags": ["ImageService"]}], "method": "get"}

        
// ListImages returns all the images.

func (a Image) ListImages(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/images"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// InvalidateScanAndRegistryCaches removes the image metadata cache.

func (a Image) InvalidateScanAndRegistryCaches(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/images/cache/invalidate"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetImage returns the image given its ID.

func (a Image) GetImage(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("includeSnoozed-boolean,stripDescription-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/images/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// ListImages returns all the images.

func (a Image) CountImages(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query-string,pagination_limit-integer,pagination_offset-integer,pagination_sortOption_field-string,pagination_sortOption_reversed-boolean",args)
    if !ok {
		return nil, fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/imagescount"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetWatchedImages returns the list of image names that are currently being watched.

func (a Image) GetWatchedImages(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/watchedimages"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}        // debug: {"detail": [{"operationId": "ImageService_DeleteImages", "parameters": [{"in": "query", "name": "query.query", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.limit", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.offset", "required": false, "schema": {"format": "int32", "type": "integer"}}, {"in": "query", "name": "query.pagination.sortOption.field", "required": false, "schema": {"type": "string"}}, {"in": "query", "name": "query.pagination.sortOption.reversed", "required": false, "schema": {"type": "boolean"}}, {"in": "query", "name": "confirm", "required": false, "schema": {"type": "boolean"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1DeleteImagesResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "DeleteImage removes the images based on a query", "tags": ["ImageService"]}, {"operationId": "ImageService_UnwatchImage", "parameters": [{"description": "The name of the image to unwatch.\nShould match the name of a previously watched image.", "in": "query", "name": "name", "required": false, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1Empty"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "UnwatchImage marks an image name to no longer be watched.\nIt returns successfully if the image is no longer being watched\nafter the call, irrespective of whether the image was already being watched.", "tags": ["ImageService"]}], "method": "delete"}

        
// DeleteImage removes the images based on a query

func (a Image) DeleteImages(args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("query_query-string,query_pagination_limit-integer,query_pagination_offset-integer,query_pagination_sortOption_field-string,query_pagination_sortOption_reversed-boolean,confirm-boolean",args)
    if !ok {
		return  fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/images"

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}
// UnwatchImage marks an image name to no longer be watched. It returns successfully if the image is no longer being watched after the call, irrespective of whether the image was already being watched.

func (a Image) UnwatchImage(args map[string]interface{}) ( error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})
    ok := tools.CheckFieldsValid("name-string",args)
    if !ok {
		return  fmt.Errorf("Variable Vaidation Failed")
    }
uriPath := "/v1/watchedimages"

     err := tools.DeleteResource(a.Client, uriPath, args)
    if err != nil {
		return  fmt.Errorf("error: %d", err)
	}
    return  nil

}            // debug: {"detail": [{"operationId": "ImageService_ScanImage", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ScanImageRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/storageImage"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ScanImage scans a single image and returns the result", "tags": ["ImageService"]}, {"operationId": "ImageService_WatchImage", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1WatchImageRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1WatchImageResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "WatchImage marks an image name as to be watched.", "tags": ["ImageService"]}], "method": "post"}

        
// ScanImage scans a single image and returns the result

func (a Image) ScanImage(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/images/scan"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// WatchImage marks an image name as to be watched.

func (a Image) WatchImage(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/watchedimages"

    msi, err := tools.PostResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}    