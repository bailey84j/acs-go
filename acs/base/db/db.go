package db

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type DB struct {
	Client client.Client
}
func (a DB) GetExportCapabilities(args map[string]interface{}) {

//  
// NOT Required 




    // GetExportCapabilitiesResource()





}
func (a DB) GetActiveRestoreProcess(args map[string]interface{}) {

//  
// NOT Required 




    // GetActiveRestoreProcessResource()





}
func (a DB) CancelRestoreProcess(id string,args map[string]interface{}) {

//  
// NOT Required []




    // CancelRestoreProcessResource()





}
