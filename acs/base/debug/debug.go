package debug

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Debug struct {
	Client client.Client
}
// Stream authorization traces for all incoming requests.
func (a Debug) StreamAuthzTraces(args map[string]interface{}) {

//  
// NOT Required 




    // StreamAuthzTracesResource()





}
// Get the current logging level for StackRox services.
func (a Debug) GetLogLevel(args map[string]interface{}) {

//  
// NOT Required [{'name': 'modules', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string'}}}]

    ok := tools.CheckFieldsValid("modules_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetLogLevelResource()





}
