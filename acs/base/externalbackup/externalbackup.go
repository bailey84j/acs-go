package externalbackup

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type ExternalBackup struct {
	Client client.Client
}
// GetExternalBackups returns all external backup configurations.
func (a ExternalBackup) GetExternalBackups(args map[string]interface{}) {

//  
// NOT Required 




    // GetResource()





}
// GetExternalBackup returns the external backup configuration given its ID.
func (a ExternalBackup) GetExternalBackup(id string,args map[string]interface{}) {

//  
// NOT Required []




    // GetResource()





}
// DeleteExternalBackup removes an external backup configuration given its ID.
func (a ExternalBackup) DeleteExternalBackup(id string,args map[string]interface{}) {

//  
// NOT Required []




    // DeleteResource()





}
