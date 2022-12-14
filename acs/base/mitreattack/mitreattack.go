package mitreattack

import (
    "fmt"
    tools "github.com/bailey84j/acs-go/acs/tools"
    client "github.com/bailey84j/acs-go/acs/client"
)

type MitreAttack struct {
	Client *client.Client
}
        // debug: {"detail": [{"operationId": "MitreAttackService_ListMitreAttackVectors", "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1ListMitreAttackVectorsResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "ListMitreAttackVectors returns all MITRE ATT&CK vectors.", "tags": ["MitreAttackService"]}, {"operationId": "MitreAttackService_GetMitreAttackVector", "parameters": [{"in": "path", "name": "id", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/v1GetMitreVectorResponse"}}}, "description": "A successful response."}, "default": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/runtimeError"}}}, "description": "An unexpected error response."}}, "summary": "GetMitreAttackVector returns the full MITRE ATT&CK vector for a tactic with all its techniques.", "tags": ["MitreAttackService"]}], "method": "get"}

        
// ListMitreAttackVectors returns all MITRE ATT&CK vectors.

func (a MitreAttack) ListMitreAttackVectors(args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/mitreattackvectors"

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}
// GetMitreAttackVector returns the full MITRE ATT&CK vector for a tactic with all its techniques.

func (a MitreAttack) GetMitreAttackVector(id string,args map[string]interface{}) (map[string]interface{}, error) {

    tools.LogPrint(tools.PrintLog{"Validating Fields", "INFO", a.Client.LogLevel})

uriPath := "/v1/mitreattackvectors/" + id + ""

    msi, err := tools.GetResource(a.Client, uriPath, args)
    if err != nil {
		return nil, fmt.Errorf("error: %d", err)
	}
    return msi, nil

}                