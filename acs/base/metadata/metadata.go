package metadata

import (
    "fmt"

    client "github.com/bailey84j/acs-go/acs/client"
    tools "github.com/bailey84j/acs-go/acs/tools"
)

type Metadata struct {
	Client client.Client
}
func (a Metadata) GetMetadata(args map[string]interface{}) {

//  
// NOT Required 




    // GetResource()





}
func (a Metadata) TLSChallenge(args map[string]interface{}) {

//  
// NOT Required [{'name': 'challengeToken', 'description': 'generated challenge token by the service asking for TLS certs.', 'in': 'query', 'required': False, 'schema': {'type': 'string'}}]

    ok := tools.CheckFieldsValid("challengeToken_string",args)
    if !ok {
		fmt.Printf("Variable Vaidation Failed")
    }


    // TLSChallengeResource()





}
