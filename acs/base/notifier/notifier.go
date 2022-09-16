package notifier

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Notifier struct {
	Client client.Client
}
// GetNotifiers returns all notifier configurations.
func (a Notifier) GetNotifiers(args map[string]interface{}) {

//  
// NOT Required 




// GetResource()





}
// GetNotifier returns the notifier configuration given its ID.
func (a Notifier) GetNotifier(id string,args map[string]interface{}) {

//  
// NOT Required []




// GetResource()





}
// DeleteNotifier removes a notifier configuration given its ID.
func (a Notifier) DeleteNotifier(id string,args map[string]interface{}) {

//  
// NOT Required [{'name': 'force', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}]

ok := tools.CheckFieldsValid("force_boolean",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


// DeleteResource()





}
