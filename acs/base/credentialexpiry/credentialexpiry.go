package credentialexpiry

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type CredentialExpiry struct {
	Client client.Client
}
// GetCertExpiry returns information related to the expiry component mTLS certificate.
func (a CredentialExpiry) GetCertExpiry(args map[string]interface{}) {

//  
// NOT Required [{'name': 'component', 'in': 'query', 'required': False, 'schema': {'type': 'string', 'enum': ['UNKNOWN', 'CENTRAL', 'SCANNER'], 'default': 'UNKNOWN'}}]

    ok := tools.CheckFieldsValid("component_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // GetCertExpiryResource()





}
