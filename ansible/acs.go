package acs

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

    alert "github.com/bailey84j/acs-go/acs/base/alert"
    apitoken "github.com/bailey84j/acs-go/acs/base/apitoken"
    auth "github.com/bailey84j/acs-go/acs/base/auth"
    authprovider "github.com/bailey84j/acs-go/acs/base/authprovider"
    externalbackup "github.com/bailey84j/acs-go/acs/base/externalbackup"
    centralhealth "github.com/bailey84j/acs-go/acs/base/centralhealth"
    clusterinit "github.com/bailey84j/acs-go/acs/base/clusterinit"
    clusters "github.com/bailey84j/acs-go/acs/base/clusters"
    compliancemanagement "github.com/bailey84j/acs-go/acs/base/compliancemanagement"
    compliance "github.com/bailey84j/acs-go/acs/base/compliance"
    config "github.com/bailey84j/acs-go/acs/base/config"
    credentialexpiry "github.com/bailey84j/acs-go/acs/base/credentialexpiry"
    clustercve "github.com/bailey84j/acs-go/acs/base/clustercve"
    cve "github.com/bailey84j/acs-go/acs/base/cve"
    imagecve "github.com/bailey84j/acs-go/acs/base/imagecve"
    nodecve "github.com/bailey84j/acs-go/acs/base/nodecve"
    db "github.com/bailey84j/acs-go/acs/base/db"
    debug "github.com/bailey84j/acs-go/acs/base/debug"
    deployment "github.com/bailey84j/acs-go/acs/base/deployment"
    detection "github.com/bailey84j/acs-go/acs/base/detection"
    featureflag "github.com/bailey84j/acs-go/acs/base/featureflag"
    group "github.com/bailey84j/acs-go/acs/base/group"
    imageintegration "github.com/bailey84j/acs-go/acs/base/imageintegration"
    image "github.com/bailey84j/acs-go/acs/base/image"
    integrationhealth "github.com/bailey84j/acs-go/acs/base/integrationhealth"
    license "github.com/bailey84j/acs-go/acs/base/license"
    metadata "github.com/bailey84j/acs-go/acs/base/metadata"
    mitreattack "github.com/bailey84j/acs-go/acs/base/mitreattack"
    namespace "github.com/bailey84j/acs-go/acs/base/namespace"
    networkbaseline "github.com/bailey84j/acs-go/acs/base/networkbaseline"
    networkgraph "github.com/bailey84j/acs-go/acs/base/networkgraph"
    networkpolicy "github.com/bailey84j/acs-go/acs/base/networkpolicy"
    node "github.com/bailey84j/acs-go/acs/base/node"
    notifier "github.com/bailey84j/acs-go/acs/base/notifier"
    ping "github.com/bailey84j/acs-go/acs/base/ping"
    pod "github.com/bailey84j/acs-go/acs/base/pod"
    policycategory "github.com/bailey84j/acs-go/acs/base/policycategory"
    policy "github.com/bailey84j/acs-go/acs/base/policy"
    probeupload "github.com/bailey84j/acs-go/acs/base/probeupload"
    processbaseline "github.com/bailey84j/acs-go/acs/base/processbaseline"
    process "github.com/bailey84j/acs-go/acs/base/process"
    rbac "github.com/bailey84j/acs-go/acs/base/rbac"
    reportconfiguration "github.com/bailey84j/acs-go/acs/base/reportconfiguration"
    report "github.com/bailey84j/acs-go/acs/base/report"
    role "github.com/bailey84j/acs-go/acs/base/role"
    search "github.com/bailey84j/acs-go/acs/base/search"
    secret "github.com/bailey84j/acs-go/acs/base/secret"
    sensorupgrade "github.com/bailey84j/acs-go/acs/base/sensorupgrade"
    account "github.com/bailey84j/acs-go/acs/base/account"
    identity "github.com/bailey84j/acs-go/acs/base/identity"
    signatureintegration "github.com/bailey84j/acs-go/acs/base/signatureintegration"
    summary "github.com/bailey84j/acs-go/acs/base/summary"
    telemetry "github.com/bailey84j/acs-go/acs/base/telemetry"
    user "github.com/bailey84j/acs-go/acs/base/user"
    vulnerabilityrequest "github.com/bailey84j/acs-go/acs/base/vulnerabilityrequest"
)

const (
	defaultBaseURL = "https://tower.fiscloudservices.com/"
	defaultAPIPath = "api/v1"
	userAgent      = "jarvis"
	logLevel       = "info"
)

type API struct {
	client               *client.Client
    Alert                *alert.Alert
    APIToken                *apitoken.APIToken
    Auth                *auth.Auth
    AuthProvider                *authprovider.AuthProvider
    ExternalBackup                *externalbackup.ExternalBackup
    CentralHealth                *centralhealth.CentralHealth
    ClusterInit                *clusterinit.ClusterInit
    Clusters                *clusters.Clusters
    ComplianceManagement                *compliancemanagement.ComplianceManagement
    Compliance                *compliance.Compliance
    Config                *config.Config
    CredentialExpiry                *credentialexpiry.CredentialExpiry
    ClusterCVE                *clustercve.ClusterCVE
    CVE                *cve.CVE
    ImageCVE                *imagecve.ImageCVE
    NodeCVE                *nodecve.NodeCVE
    DB                *db.DB
    Debug                *debug.Debug
    Deployment                *deployment.Deployment
    Detection                *detection.Detection
    FeatureFlag                *featureflag.FeatureFlag
    Group                *group.Group
    ImageIntegration                *imageintegration.ImageIntegration
    Image                *image.Image
    IntegrationHealth                *integrationhealth.IntegrationHealth
    License                *license.License
    Metadata                *metadata.Metadata
    MitreAttack                *mitreattack.MitreAttack
    Namespace                *namespace.Namespace
    NetworkBaseline                *networkbaseline.NetworkBaseline
    NetworkGraph                *networkgraph.NetworkGraph
    NetworkPolicy                *networkpolicy.NetworkPolicy
    Node                *node.Node
    Notifier                *notifier.Notifier
    Ping                *ping.Ping
    Pod                *pod.Pod
    PolicyCategory                *policycategory.PolicyCategory
    Policy                *policy.Policy
    ProbeUpload                *probeupload.ProbeUpload
    ProcessBaseline                *processbaseline.ProcessBaseline
    Process                *process.Process
    Rbac                *rbac.Rbac
    ReportConfiguration                *reportconfiguration.ReportConfiguration
    Report                *report.Report
    Role                *role.Role
    Search                *search.Search
    Secret                *secret.Secret
    SensorUpgrade                *sensorupgrade.SensorUpgrade
    Account                *account.Account
    Identity                *identity.Identity
    SignatureIntegration                *signatureintegration.SignatureIntegration
    Summary                *summary.Summary
    Telemetry                *telemetry.Telemetry
    User                *user.User
    VulnerabilityRequest                *vulnerabilityrequest.VulnerabilityRequest

// Authenticate is used to set and store Basic Auth for this client instance
func (a *API) Authenticate(username, password string) error {
	if len(username) > 0 && len(password) > 0 {
		a.client.Username = username
		a.client.Password = password
		return nil
	}
	return fmt.Errorf("no valid authentication credentials were provided")
}

func NewACSAPI(endpoint string) API {
	url, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("error parsing provided url: %s, Using default: %s", err.Error(), defaultBaseURL)
		url, _ = url.Parse(defaultBaseURL)
	}
	url.Path = path.Join(url.Path, defaultAPIPath)

	// This is initialized this way because of composition
	client := &client.Client{
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
        Alert : &alert.Alert {
			Client: *client,
		},
        APIToken : &apitoken.APIToken {
			Client: *client,
		},
        Auth : &auth.Auth {
			Client: *client,
		},
        AuthProvider : &authprovider.AuthProvider {
			Client: *client,
		},
        ExternalBackup : &externalbackup.ExternalBackup {
			Client: *client,
		},
        CentralHealth : &centralhealth.CentralHealth {
			Client: *client,
		},
        ClusterInit : &clusterinit.ClusterInit {
			Client: *client,
		},
        Clusters : &clusters.Clusters {
			Client: *client,
		},
        ComplianceManagement : &compliancemanagement.ComplianceManagement {
			Client: *client,
		},
        Compliance : &compliance.Compliance {
			Client: *client,
		},
        Config : &config.Config {
			Client: *client,
		},
        CredentialExpiry : &credentialexpiry.CredentialExpiry {
			Client: *client,
		},
        ClusterCVE : &clustercve.ClusterCVE {
			Client: *client,
		},
        CVE : &cve.CVE {
			Client: *client,
		},
        ImageCVE : &imagecve.ImageCVE {
			Client: *client,
		},
        NodeCVE : &nodecve.NodeCVE {
			Client: *client,
		},
        DB : &db.DB {
			Client: *client,
		},
        Debug : &debug.Debug {
			Client: *client,
		},
        Deployment : &deployment.Deployment {
			Client: *client,
		},
        Detection : &detection.Detection {
			Client: *client,
		},
        FeatureFlag : &featureflag.FeatureFlag {
			Client: *client,
		},
        Group : &group.Group {
			Client: *client,
		},
        ImageIntegration : &imageintegration.ImageIntegration {
			Client: *client,
		},
        Image : &image.Image {
			Client: *client,
		},
        IntegrationHealth : &integrationhealth.IntegrationHealth {
			Client: *client,
		},
        License : &license.License {
			Client: *client,
		},
        Metadata : &metadata.Metadata {
			Client: *client,
		},
        MitreAttack : &mitreattack.MitreAttack {
			Client: *client,
		},
        Namespace : &namespace.Namespace {
			Client: *client,
		},
        NetworkBaseline : &networkbaseline.NetworkBaseline {
			Client: *client,
		},
        NetworkGraph : &networkgraph.NetworkGraph {
			Client: *client,
		},
        NetworkPolicy : &networkpolicy.NetworkPolicy {
			Client: *client,
		},
        Node : &node.Node {
			Client: *client,
		},
        Notifier : &notifier.Notifier {
			Client: *client,
		},
        Ping : &ping.Ping {
			Client: *client,
		},
        Pod : &pod.Pod {
			Client: *client,
		},
        PolicyCategory : &policycategory.PolicyCategory {
			Client: *client,
		},
        Policy : &policy.Policy {
			Client: *client,
		},
        ProbeUpload : &probeupload.ProbeUpload {
			Client: *client,
		},
        ProcessBaseline : &processbaseline.ProcessBaseline {
			Client: *client,
		},
        Process : &process.Process {
			Client: *client,
		},
        Rbac : &rbac.Rbac {
			Client: *client,
		},
        ReportConfiguration : &reportconfiguration.ReportConfiguration {
			Client: *client,
		},
        Report : &report.Report {
			Client: *client,
		},
        Role : &role.Role {
			Client: *client,
		},
        Search : &search.Search {
			Client: *client,
		},
        Secret : &secret.Secret {
			Client: *client,
		},
        SensorUpgrade : &sensorupgrade.SensorUpgrade {
			Client: *client,
		},
        Account : &account.Account {
			Client: *client,
		},
        Identity : &identity.Identity {
			Client: *client,
		},
        SignatureIntegration : &signatureintegration.SignatureIntegration {
			Client: *client,
		},
        Summary : &summary.Summary {
			Client: *client,
		},
        Telemetry : &telemetry.Telemetry {
			Client: *client,
		},
        User : &user.User {
			Client: *client,
		},
        VulnerabilityRequest : &vulnerabilityrequest.VulnerabilityRequest {
			Client: *client,
		},
	}
}