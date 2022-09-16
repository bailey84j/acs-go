package license

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type License struct {
	Client client.Client
}
func (a License) GetActiveLicenseKey(args map[string]interface{}) {

//  
// NOT Required 




    // GetActiveResource()





}
func (a License) GetActiveLicenseExpiration(args map[string]interface{}) {

//  
// NOT Required 




    // GetActiveResource()





}
func (a License) GetLicenses(args map[string]interface{}) {

//  
// NOT Required [{'name': 'active', 'in': 'query', 'required': False, 'schema': {'type': 'boolean'}}, {'name': 'statuses', 'in': 'query', 'required': False, 'explode': True, 'schema': {'type': 'array', 'items': {'type': 'string', 'enum': ['UNKNOWN', 'VALID', 'REVOKED', 'NOT_YET_VALID', 'EXPIRED', 'OTHER']}}}]

    ok := tools.CheckFieldsValid("active_boolean,statuses_array",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetResource()





}
