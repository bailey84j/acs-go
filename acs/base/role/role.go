package role

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Role struct {
	Client client.Client
}
func (a Role) GetMyPermissions(args map[string]interface{}) {

//  
// NOT Required 




    // GetMyPermissionsResource()





}
func (a Role) ListPermissionSets(args map[string]interface{}) {

//  
// NOT Required 




    // ListPermissionSetsResource()





}
func (a Role) GetPermissionSet(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetPermissionSetResource()





}
func (a Role) GetResources(args map[string]interface{}) {

//  
// NOT Required 




    // GetResourcesResource()





}
func (a Role) GetRoles(args map[string]interface{}) {

//  
// NOT Required 




    // GetResource()





}
func (a Role) GetRole(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
func (a Role) ListSimpleAccessScopes(args map[string]interface{}) {

//  
// NOT Required 




    // ListSimpleAccessScopesResource()





}
func (a Role) GetSimpleAccessScope(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetSimpleAccessScopeResource()





}
func (a Role) DeletePermissionSet(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeletePermissionSetResource()





}
func (a Role) DeleteRole(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteResource()





}
func (a Role) DeleteSimpleAccessScope(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteSimpleAccessScopeResource()





}
