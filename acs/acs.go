package acs

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

	alert "github.com/bailey84j/acs-go/base/alert"
	apitoken "github.com/bailey84j/acs-go/base/apitoken"
	authprovider "github.com/bailey84j/acs-go/base/authprovider"
	centralhealth "github.com/bailey84j/acs-go/base/centralhealth"
	clustercve "github.com/bailey84j/acs-go/base/clustercve"
	clusterinit "github.com/bailey84j/acs-go/base/clusterinit"
	clusters "github.com/bailey84j/acs-go/base/clusters"
	compliance "github.com/bailey84j/acs-go/base/compliance"
	compliancemanagement "github.com/bailey84j/acs-go/base/compliancemanagement"
	config "github.com/bailey84j/acs-go/base/config"
	credentialexpiry "github.com/bailey84j/acs-go/base/credentialexpiry"
	cve "github.com/bailey84j/acs-go/base/cve"
	db "github.com/bailey84j/acs-go/base/db"
	debug "github.com/bailey84j/acs-go/base/debug"
	deployment "github.com/bailey84j/acs-go/base/deployment"
	detection "github.com/bailey84j/acs-go/base/detection"
	externalbackup "github.com/bailey84j/acs-go/base/externalbackup"
	featureflag "github.com/bailey84j/acs-go/base/featureflag"
	group "github.com/bailey84j/acs-go/base/group"
	image "github.com/bailey84j/acs-go/base/image"
	imagecve "github.com/bailey84j/acs-go/base/imagecve"
	imageintegration "github.com/bailey84j/acs-go/base/imageintegration"
	integrationhealth "github.com/bailey84j/acs-go/base/integrationhealth"
	license "github.com/bailey84j/acs-go/base/license"
	metadata "github.com/bailey84j/acs-go/base/metadata"
	mitreattack "github.com/bailey84j/acs-go/base/mitreattack"
	namespace "github.com/bailey84j/acs-go/base/namespace"
	networkbaseline "github.com/bailey84j/acs-go/base/networkbaseline"
	networkgraph "github.com/bailey84j/acs-go/base/networkgraph"
	networkpolicy "github.com/bailey84j/acs-go/base/networkpolicy"
	node "github.com/bailey84j/acs-go/base/node"
	nodecve "github.com/bailey84j/acs-go/base/nodecve"
	notifier "github.com/bailey84j/acs-go/base/notifier"
	ping "github.com/bailey84j/acs-go/base/ping"
	pod "github.com/bailey84j/acs-go/base/pod"
	policy "github.com/bailey84j/acs-go/base/policy"
	policycategory "github.com/bailey84j/acs-go/base/policycategory"
	probeupload "github.com/bailey84j/acs-go/base/probeupload"
	process "github.com/bailey84j/acs-go/base/process"
	processbaseline "github.com/bailey84j/acs-go/base/processbaseline"
	rbac "github.com/bailey84j/acs-go/base/rbac"
	report "github.com/bailey84j/acs-go/base/report"
	reportconfiguration "github.com/bailey84j/acs-go/base/reportconfiguration"
	role "github.com/bailey84j/acs-go/base/role"
	search "github.com/bailey84j/acs-go/base/search"
	sensorupgrade "github.com/bailey84j/acs-go/base/sensorupgrade"
	serviceaccount "github.com/bailey84j/acs-go/base/serviceaccount"
	serviceidenity "github.com/bailey84j/acs-go/base/serviceidenity"
	signatureintegration "github.com/bailey84j/acs-go/base/signatureintegration"
	summary "github.com/bailey84j/acs-go/base/summary"
	telemetry "github.com/bailey84j/acs-go/base/telemetry"
	user "github.com/bailey84j/acs-go/base/user"
	vulnerabilityrequest "github.com/bailey84j/acs-go/base/vulnerabilityrequest"
)

const (
	defaultBaseURL = "https://tower.fiscloudservices.com/"
	defaultAPIPath = "api/v1"
	userAgent      = "jarvis"
	logLevel       = "info"
)

type API struct {
	client *Client

	APIToken             *apitoken.APIToken
	AuthProvider         *authprovider.AuthProvider
	ExternalBackup       *externalbackup.ExternalBackup
	CentralHealth        *centralhealth.CentralHealth
	ClusterInit          *clusterinit.ClusterInit
	Clusters             *clusters.Clusters
	ComplianceManagement *compliancemanagement.ComplianceManagement
	Compliance           *compliance.Compliance
	Config               *config.Config
	CredentialExpiry     *credentialexpiry.CredentialExpiry
	ClusterCVE           *clustercve.ClusterCVE
	CVE                  *cve.CVE
	ImageCVE             *imagecve.ImageCVE
	NodeCVE              *nodecve.NodeCVE
	DB                   *db.DB
	Debug                *debug.Debug
	Deployment           *deployment.Deployment
	Detection            *detection.Detection
	FeatureFlag          *featureflag.FeatureFlag
	Group                *group.Group
	ImageIntegration     *imageintegration.ImageIntegration
	Image                *image.Image
	IntegrationHealth    *integrationhealth.IntegrationHealth
	License              *license.License
	Metadata             *metadata.Metadata
	MitreAttack          *mitreattack.MitreAttack
	Namespace            *namespace.Namespace
	NetworkBaseline      *networkbaseline.NetworkBaseline
	NetworkGraph         *networkgraph.NetworkGraph
	NetworkPolicy        *networkpolicy.NetworkPolicy
	Node                 *node.Node
	Notifier             *notifier.Notifier
	Ping                 *ping.Ping
	Pod                  *pod.Pod
	PolicyCategory       *policycategory.PolicyCategory
	Policy               *policy.Policy
	ProbeUpload          *probeupload.ProbeUpload
	ProcessBaseline      *processbaseline.ProcessBaseline
	Process              *process.Process
	RBAC                 *rbac.RBAC
	ReportConfiguration  *reportconfiguration.ReportConfiguration
	Report               *report.Report
	Role                 *role.Role
	Search               *search.Search
	SensorUpgrade        *sensorupgrade.SensorUpgrade
	ServiceAccount       *serviceaccount.ServiceAccount
	ServiceIdenity       *serviceidenity.ServiceIdenity
	SignatureIntegration *signatureintegration.SignatureIntegration
	Summary              *summary.Summary
	Telemetry            *telemetry.Telemetry
	User                 *user.User
	VulnerabilityRequest *vulnerabilityrequest.VulnerabilityRequest
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
		Alert: &alert.Alert{
			Client: client,
		},
		APIToken: &apitoken.APIToken{
			Client: client,
		},
		APIToken: &apitoken.APIToken{
			Client: client,
		},
		AuthProvider: &authprovider.AuthProvider{
			Client: client,
		},
		ExternalBackup: &externalbackup.ExternalBackup{
			Client: client,
		},
		CentralHealth: &centralhealth.CentralHealth{
			Client: client,
		},
		ClusterInit: &clusterinit.ClusterInit{
			Client: client,
		},
		Clusters: &clusters.Clusters{
			Client: client,
		},
		ComplianceManagement: &compliancemanagement.ComplianceManagement{
			Client: client,
		},
		Compliance: &compliance.Compliance{
			Client: client,
		},
		Config: &config.Config{
			Client: client,
		},
		CredentialExpiry: &credentialexpiry.CredentialExpiry{
			Client: client,
		},
		ClusterCVE: &clustercve.ClusterCVE{
			Client: client,
		},
		CVE: &cve.CVE{
			Client: client,
		},
		ImageCVE: &imagecve.ImageCVE{
			Client: client,
		},
		NodeCVE: &nodecve.NodeCVE{
			Client: client,
		},
		DB: &db.DB{
			Client: client,
		},
		Debug: &debug.Debug{
			Client: client,
		},
		Deployment: &deployment.Deployment{
			Client: client,
		},
		Detection: &detection.Detection{
			Client: client,
		},
		FeatureFlag: &featureflag.FeatureFlag{
			Client: client,
		},
		Group: &group.Group{
			Client: client,
		},
		ImageIntegration: &imageintegration.ImageIntegration{
			Client: client,
		},
		Image: &image.Image{
			Client: client,
		},
		IntegrationHealth: &integrationhealth.IntegrationHealth{
			Client: client,
		},
		License: &license.License{
			Client: client,
		},
		Metadata: &metadata.Metadata{
			Client: client,
		},
		MitreAttack: &mitreattack.MitreAttack{
			Client: client,
		},
		Namespace: &namespace.Namespace{
			Client: client,
		},
		NetworkBaseline: &networkbaseline.NetworkBaseline{
			Client: client,
		},
		NetworkGraph: &networkgraph.NetworkGraph{
			Client: client,
		},
		NetworkPolicy: &networkpolicy.NetworkPolicy{
			Client: client,
		},
		Node: &node.Node{
			Client: client,
		},
		Notifier: &notifier.Notifier{
			Client: client,
		},
		Ping: &ping.Ping{
			Client: client,
		},
		Pod: &pod.Pod{
			Client: client,
		},
		PolicyCategory: &policycategory.PolicyCategory{
			Client: client,
		},
		Policy: &policy.Policy{
			Client: client,
		},
		ProbeUpload: &probeupload.ProbeUpload{
			Client: client,
		},
		ProcessBaseline: &processbaseline.ProcessBaseline{
			Client: client,
		},
		Process: &process.Process{
			Client: client,
		},
		RBAC: &rbac.RBAC{
			Client: client,
		},
		ReportConfiguration: &reportconfiguration.ReportConfiguration{
			Client: client,
		},
		Report: &report.Report{
			Client: client,
		},
		Role: &role.Role{
			Client: client,
		},
		Search: &search.Search{
			Client: client,
		},
		SensorUpgrade: &sensorupgrade.SensorUpgrade{
			Client: client,
		},
		ServiceAccount: &serviceaccount.ServiceAccount{
			Client: client,
		},
		ServiceIdenity: &serviceidenity.ServiceIdenity{
			Client: client,
		},
		SignatureIntegration: &signatureintegration.SignatureIntegration{
			Client: client,
		},
		Summary: &summary.Summary{
			Client: client,
		},
		Telemetry: &telemetry.Telemetry{
			Client: client,
		},
		User: &user.User{
			Client: client,
		},
		VulnerabilityRequest: &vulnerabilityrequest.VulnerabilityRequest{
			Client: client,
		},
	}
}
