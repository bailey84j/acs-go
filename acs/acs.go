package acs

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	defaultBaseURL = "https://tower.fiscloudservices.com/"
	defaultAPIPath = "api/v1"
	userAgent      = "jarvis"
	logLevel       = "info"
)

type API struct {
	client *Client

	Attach         *Attach
	Tables         *Tables
	ServiceCatalog *ServiceCatalog
	ITSM           *ITSM
}

// Authenticate is used to set and store Basic Auth for this client instance
func (a *API) Authenticate(username, password string) error {
	if len(username) > 0 && len(password) > 0 {
		a.client.Username = username
		a.client.Password = password
		return nil
	}
	return fmt.Errorf("no valid authentication credentials were provided")
}

func NewSnowAPI(endpoint string) API {
	url, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("error parsing provided url: %s, Using default: %s", err.Error(), defaultBaseURL)
		url, _ = url.Parse(defaultBaseURL)
	}
	url.Path = path.Join(url.Path, defaultAPIPath)

	// This is initialized this way because of composition
	client := &Client{
		&http.Client{},
		userAgent,
		"",
		"",
		url,
		3,
		"3s",
		logLevel,
	}

	return API{
		client: client,
		Alert: &Alert{
			Client: client,
		},
		APIToken: &APIToken{
			Client: client,
		},
		AuthProvider: &AuthProvider{
			Client: client,
		},
		ExternalBackup: &ExternalBackup{
			Client: client,
		},
		CentralHealth: &CentralHealth{
			Client: client,
		},
		ClusterInit: &ClusterInit{
			Client: client,
		},
		Clusters: &Clusters{
			Client: client,
		},
		ComplianceManagement: &ComplianceManagement{
			Client: client,
		},
		Compliance: &Compliance{
			Client: client,
		},
		Config: &Config{
			Client: client,
		},
		CredentialExpiry: &CredentialExpiry{
			Client: client,
		},
		ClusterCVE: &ClusterCVE{
			Client: client,
		},
		CVE: &CVE{
			Client: client,
		},
		ImageCVE: &ImageCVE{
			Client: client,
		},
		NodeCVE: &NodeCVE{
			Client: client,
		},
		DB: &DB{
			Client: client,
		},
		Debug: &Debug{
			Client: client,
		},
		Deployment: &Deployment{
			Client: client,
		},
		Detection: &Detection{
			Client: client,
		},
		FeatureFlag: &FeatureFlag{
			Client: client,
		},
		Group: &Group{
			Client: client,
		},
		ImageIntegration: &ImageIntegration{
			Client: client,
		},
		Image: &Image{
			Client: client,
		},
		IntegrationHealth: &IntegrationHealth{
			Client: client,
		},
		License: &License{
			Client: client,
		},
		Metadata: &Metadata{
			Client: client,
		},
		MitreAttack: &MitreAttack{
			Client: client,
		},
		Namespace: &Namespace{
			Client: client,
		},
		NetworkBaseline: &NetworkBaseline{
			Client: client,
		},
		NetworkGraph: &NetworkGraph{
			Client: client,
		},
		NetworkPolicy: &NetworkPolicy{
			Client: client,
		},
		Node: &Node{
			Client: client,
		},
		Notifier: &Notifier{
			Client: client,
		},
		Ping: &Ping{
			Client: client,
		},
		Pod: &Pod{
			Client: client,
		},
		PolicyCategory: &PolicyCategory{
			Client: client,
		},
		Policy: &Policy{
			Client: client,
		},
		ProbeUpload: &ProbeUpload{
			Client: client,
		},
		ProcessBaseline: &ProcessBaseline{
			Client: client,
		},
		Process: &Process{
			Client: client,
		},
		RBAC: &RBAC{
			Client: client,
		},
		ReportConfiguration: &ReportConfiguration{
			Client: client,
		},
		Report: &Report{
			Client: client,
		},
		Role: &Role{
			Client: client,
		},
		Search: &Search{
			Client: client,
		},
		SensorUpgrade: &SensorUpgrade{
			Client: client,
		},
		ServiceAccount: &ServiceAccount{
			Client: client,
		},
		ServiceIdenity: &ServiceIdenity{
			Client: client,
		},
		SignatureIntegration: &SignatureIntegration{
			Client: client,
		},
		Summary: &Summary{
			Client: client,
		},
		Telemetry: &Telemetry{
			Client: client,
		},
		User: &User{
			Client: client,
		},
		VulnerabilityRequest: &VulnerabilityRequest{
			Client: client,
		},
	}
}
