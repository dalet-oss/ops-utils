package metadata

import (
	"fmt"

	"github.com/dalet-oss/ops-utils/clog"
	"gopkg.in/yaml.v3"
)

// PlatformMetadata is the root definition of a managed platform
type PlatformMetadata struct {
	Customer     PlatformMetadataCustomer       `yaml:"customer"`
	Infra        PlatformMetadataInfra          `yaml:"infra"`
	AWS          PlatformMetadataInfraAWS       `yaml:"aws,omitempty"`
	Azure        PlatformMetadataInfraAzure     `yaml:"azure,omitempty"`
	Kowabunga    PlatformMetadataInfraKowabunga `yaml:"kowabunga,omitempty"`
	Requirements PlatformMetadataRequirements   `yaml:"requirements,omitempty"`
	Product      PlatformMetadataProduct        `yaml:"product"`
	Secrets      PlatformMetadataSecrets        `yaml:"secrets"`
	Environments []PlatformMetadataEnvironment  `yaml:"environments"`
}

// PlatformMetadataCustomer contains customer-specific information
type PlatformMetadataCustomer struct {
	Name    string `yaml:"name"`
	Region  string `yaml:"region"`
	Country string `yaml:"country"`
}

// PlatformMetadataInfra contains infrastructure-specific information
type PlatformMetadataInfra struct {
	CostCenter    string                    `yaml:"cc,omitempty"`
	Cost          PlatformMetadataInfraCost `yaml:"cost,omitempty"`
	Provider      string                    `yaml:"provider"`
	DCH           bool                      `yaml:"dch"`
	RemoteAccess  string                    `yaml:"remote_access"`
	RemoteDetails string                    `yaml:"remote_details,omitempty"`
	Satellite     bool                      `yaml:"satellite"`
}

// PlatformMetadataInfra contains infrastructure-specific information
type PlatformMetadataInfraCost struct {
	Price    float64 `yaml:"price"`
	Currency string  `yaml:"currency"`
}

// PlatformMetadataInfraAWS contains AWS-infrastructure specific information
type PlatformMetadataInfraAWS struct {
	Region  string `yaml:"region,omitempty"`
	RoleARN string `yaml:"role_arn,omitempty"`
}

// PlatformMetadataInfraAzure contains Azure-infrastructure specific information
type PlatformMetadataInfraAzure struct {
	ResourceGroup string `yaml:"resource_group"`
}

// PlatformMetadataInfraKowabunga contains Kowabunga-infrastructure specific information
type PlatformMetadataInfraKowabunga struct {
	Endpoint string `yaml:"endpoint"` // e.g. http://kowabunga.admin.dalet.stg.ops:8080
	Owner    string `yaml:"owner"`
	Email    string `yaml:"email"`
	Region   string `yaml:"region"`
	Zone     string `yaml:"zone"`
}

// PlatformMetadataProduct contains product-related information
type PlatformMetadataProduct struct {
	Name         string                           `yaml:"name"`
	License      string                           `yaml:"license"`
	SupportLevel string                           `yaml:"support_level"`
	Monitored    PlatformMetadataProductMonitored `yaml:"monitored"`
}

// PlatformMetadataProductMonitored contains product-related monitoring status
type PlatformMetadataProductMonitored struct {
	Enabled      bool `yaml:"enabled"`
	ControlTower bool `yaml:"control_tower"`
}

// PlatformMetadataRequirements contains tools-related dependancies
type PlatformMetadataRequirements struct {
	Ansible PlatformMetadataRequirementsAnsible `yaml:"ansible,omitempty"`
	TF      PlatformMetadataRequirementsTF      `yaml:"terraform,omitempty"`
}

// PlatformMetadataRequirementsAnsible contains Ansible-related information
type PlatformMetadataRequirementsAnsible struct {
	MinVersion string `yaml:"min_version,omitempty"`
	Version    string `yaml:"version,omitempty"`
}

// PlatformMetadataRequirementsTF contains Terraform-related information
type PlatformMetadataRequirementsTF struct {
	Version string `yaml:"version,omitempty"`
}

// PlatformMetadataSecrets contains secrets management related information
type PlatformMetadataSecrets struct {
	Provider            string                    `yaml:"provider"`
	SM                  PlatformMetadataSecretsSM `yaml:"sm,omitempty"`
	ID                  string                    `yaml:"id"`
	VaultKeySuffix      string                    `yaml:"vault_key_suffix,omitempty"`
	SOPSCreatedAtSuffix string                    `yaml:"sops_created_at_suffix,omitempty"`
	SOPSPublicKeySuffix string                    `yaml:"sops_public_key_suffix,omitempty"`
	SOPSSecretKeySuffix string                    `yaml:"sops_secret_key_suffix,omitempty"`
}

// PlatformMetadataSecretsSM contains AWS Secrets Manager related information
type PlatformMetadataSecretsSM struct {
	Region        string `yaml:"region"`
	RoleARN       string `yaml:"role_arn,omitempty"`
	CustomProfile bool   `yaml:"custom_profile,omitempty"`
}

// PlatformMetadataEnvironment contains environment-related information
type PlatformMetadataEnvironment struct {
	Tag         string                                `yaml:"tag"`
	Name        string                                `yaml:"name"`
	Description string                                `yaml:"description"`
	Active      bool                                  `yaml:"active"`
	Version     string                                `yaml:"version"`
	FQDN        PlatformMetadataEnvironmentFQDN       `yaml:"fqdn,omitempty"`
	OS          PlatformMetadataEnvironmentOS         `yaml:"os"`
	Inventory   PlatformMetadataEnvironmentInventory  `yaml:"inventory"`
	StatusPage  PlatformMetadataEnvironmentStatusPage `yaml:"statuspage,omitempty"`
	Pagerduty   PlatformMetadataEnvironmentPagerduty  `yaml:"pagerduty"`
	Kubernetes  PlatformMetadataEnvironmentKubernetes `yaml:"kubernetes,omitempty"`
	Security    PlatformMetadataEnvironmentSecurity   `yaml:"security,omitempty"`
	Flex        PlatformMetadataEnvironmentFlex       `yaml:"flex,omitempty"`
	Pyramid     PlatformMetadataEnvironmentPyramid    `yaml:"pyramid,omitempty"`
}

// PlatformMetadataEnvironmentFQDN contains public domain information
type PlatformMetadataEnvironmentFQDN struct {
	Prefix string `yaml:"prefix"`
	Domain string `yaml:"domain"`
}

// PlatformMetadataEnvironmentOS contains Operating system related information
type PlatformMetadataEnvironmentOS struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version,omitempty"`
}

// PlatformMetadataEnvironmentInventory contains environment inventory related information
type PlatformMetadataEnvironmentInventory struct {
	Static bool `yaml:"static"`
}

// PlatformMetadataEnvironmentStatusPage contains environment status page related information
type PlatformMetadataEnvironmentStatusPage struct {
	Enabled     bool   `yaml:"enabled"`
	PageID      string `yaml:"page_id,omitempty"`
	ComponentID string `yaml:"component_id,omitempty"`
}

// PlatformMetadataEnvironmentPagerduty contains environment PagerDuty related information
type PlatformMetadataEnvironmentPagerduty struct {
	Enabled   bool   `yaml:"enabled"`
	ServiceID string `yaml:"service_id,omitempty"`
}

// PlatformMetadataEnvironmentKubernetes contains environment Kubernetes related information
type PlatformMetadataEnvironmentKubernetes struct {
	Enabled    bool                                            `yaml:"enabled"`
	Type       string                                          `yaml:"type,omitempty"`
	Version    string                                          `yaml:"version,omitempty"`
	Cluster    string                                          `yaml:"cluster,omitempty"`
	Controller PlatformMetadataEnvironmentKubernetesController `yaml:"controller,omitempty"`
}

// PlatformMetadataEnvironmentKubernetesController contains environment Kubernetes-controller related information
type PlatformMetadataEnvironmentKubernetesController struct {
	Endpoint string `yaml:"endpoint,omitempty"`
	Master   string `yaml:"master,omitempty"`
}

// PlatformMetadataEnvironmentSecurity contains environment security related information
type PlatformMetadataEnvironmentSecurity struct {
	Log4jCVE bool `yaml:"log4j_cve,omitempty"`
}

// PlatformMetadataEnvironmentFlex contains environment Flex related information
type PlatformMetadataEnvironmentFlex struct {
	URLs      PlatformMetadataEnvironmentURL           `yaml:"urls"`
	Transcode PlatformMetadataEnvironmentFlexTranscode `yaml:"transcode"`
	DBaaS     bool                                     `yaml:"dbaas"`
	LBaaS     bool                                     `yaml:"lbaas"`
}

// PlatformMetadataEnvironmentFlexTranscode contains environment Flex Transcoder related information
type PlatformMetadataEnvironmentFlexTranscode struct {
	Enabled bool `yaml:"enabled"`
	K8S     bool `yaml:"k8s_fsp"`
}

// PlatformMetadataEnvironmentPyramid contains environment Pyramid related information
type PlatformMetadataEnvironmentPyramid struct {
	URLs  PlatformMetadataEnvironmentURL `yaml:"urls"`
	DBaaS bool                           `yaml:"dbaas"`
	LBaaS bool                           `yaml:"lbaas"`
}

// PlatformMetadataEnvironmentURL contains environment URL related information
type PlatformMetadataEnvironmentURL struct {
	Core         string `yaml:"core"`
	Xymon        string `yaml:"xymon,omitempty"`
	Grafana      string `yaml:"grafana,omitempty"`
	Kibana       string `yaml:"kibana,omitempty"`
	Prometheus   string `yaml:"prometheus,omitempty"`
	AlertManager string `yaml:"alertmanager,omitempty"`
}

const (
	MetadataFile     = "META.yml"
	InvalidMetaField = "%s: empty or invalid %s in metadata file: '%s'"

	RegionUSWest = "US-WEST"
	RegionUSEast = "US-EAST"
	RegionCALA   = "CALA"
	RegionEMEA   = "EMEA"
	RegionAPAC   = "APAC"

	ProviderAWS       = "aws"
	ProviderAlibaba   = "ali"
	ProviderGoogle    = "gcp"
	ProviderAzure     = "azure"
	ProviderOVH       = "ovh"
	ProviderOnPrem    = "onprem"
	ProviderKowabunga = "kowabunga"

	SecretsVaultKeySuffixDefault      = "ansible_vault"
	SecretsSOPSCreatedAtSuffixDefault = "sops_created_at"
	SecretsSOPSPublicKeySuffixDefault = "sops_public_key"
	SecretsSOPSSecretKeySuffixDefault = "sops_secret_key"

	RemoteAccessNetgate         = "netgate"
	RemoteAccessNetgateNA       = "netgate-na"
	RemoteAccessNetgateEMEA     = "netgate-emea"
	RemoteAccessNetgateAPAC     = "netgate-apac"
	RemoteAccessCustomerVPN     = "customer-vpn"
	RemoteAccessCustomerBastion = "bastion"
	RemoteAccessRDP             = "rdp"

	ProductGalaxy            = "galaxy"
	ProductFlex              = "flex"
	ProductPyramid           = "pyramid"
	ProductAmberfin          = "amberfin"
	ProductMediaCortex       = "mediacortex"
	ProductCubeNG            = "cubeng"
	ProductBrioInStream      = "brio"
	ProductControlTower      = "control-tower"
	ProductIris              = "iris"
	ProductAtlas             = "atlas"
	ProductKvmVirtualization = "kvirt"

	LicensePerpetual    = "perpetual"
	LicenseSubscription = "subscription"
	LicenseDemo         = "demo"

	SupportNone     = "none"
	SupportStarter  = "starter"
	SupportPlus     = "plus"
	SupportUltimate = "ultimate"

	OsNameUbuntu = "Ubuntu"
	OsNameDebian = "Debian"
	OsNameRedHat = "RedHat"
	OsNameCentos = "Centos"
	OsNameAmazon = "AmazonLinux2"

	KubernetesTypeAuto  = "auto"
	KubernetesTypeEKS   = "eks"
	KubernetesTypeACK   = "ack"
	KubernetesTypeAKS   = "aks"
	KubernetesTypeGKE   = "gke"
	KubernetesTypeOKS   = "oks"
	KubernetesTypeAtlas = "atlas"
)

const (
	TfModuleAWSEKS               = "tf-aws-eks.git"
	TfModuleAWSEKSFieldName      = "cluster_name"
	TfModuleAWSEKSFieldVersion   = "cluster_version"
	TfModuleAWSEKSDefaultVersion = "1.23"
)

func isSupportedCustomerRegion(key string) bool {
	switch key {
	case
		RegionUSWest,
		RegionUSEast,
		RegionCALA,
		RegionEMEA,
		RegionAPAC:
		return true
	}
	return false
}

func isSupportedInfraProvider(key string) bool {
	switch key {
	case
		ProviderAWS,
		ProviderAlibaba,
		ProviderGoogle,
		ProviderAzure,
		ProviderOVH,
		ProviderOnPrem,
		ProviderKowabunga:
		return true
	}
	return false
}

func isSupportedInfraRemoteAccess(key string) bool {
	switch key {
	case
		RemoteAccessNetgate,
		RemoteAccessNetgateNA,
		RemoteAccessNetgateEMEA,
		RemoteAccessNetgateAPAC,
		RemoteAccessCustomerVPN,
		RemoteAccessCustomerBastion,
		RemoteAccessRDP:
		return true
	}
	return false
}

func isSupportedProductName(key string) bool {
	switch key {
	case
		ProductGalaxy,
		ProductFlex,
		ProductPyramid,
		ProductAmberfin,
		ProductMediaCortex,
		ProductCubeNG,
		ProductBrioInStream,
		ProductControlTower,
		ProductIris,
		ProductAtlas,
		ProductKvmVirtualization:
		return true
	}
	return false
}

func isSupportedProductLicense(key string) bool {
	switch key {
	case
		LicenseDemo,
		LicensePerpetual,
		LicenseSubscription:
		return true
	}
	return false
}

func isSupportedProductSupportLevel(key string) bool {
	switch key {
	case
		SupportNone,
		SupportStarter,
		SupportPlus,
		SupportUltimate:
		return true
	}
	return false
}

func isSupportedSecretsProvider(key string) bool {
	switch key {
	case
		ProviderAWS:
		return true
	}
	return false
}

func isSupportedOS(key string) bool {
	switch key {
	case
		OsNameUbuntu,
		OsNameDebian,
		OsNameRedHat,
		OsNameCentos,
		OsNameAmazon:
		return true
	}
	return false
}

func isSupportedKubernetesProvider(key string) bool {
	switch key {
	case
		KubernetesTypeAuto,
		KubernetesTypeEKS,
		KubernetesTypeACK,
		KubernetesTypeAKS,
		KubernetesTypeGKE,
		KubernetesTypeOKS,
		KubernetesTypeAtlas:
		return true
	}
	return false
}

func (p *PlatformMetadata) IsValid(ptf string) error {

	err := false
	type MetaValidateFunc func(string) bool
	type MetaParam struct {
		key     string
		f       MetaValidateFunc
		comment string
	}

	params := []MetaParam{
		MetaParam{p.Customer.Region, isSupportedCustomerRegion, "customer region"},
		MetaParam{p.Infra.Provider, isSupportedInfraProvider, "infra provider"},
		MetaParam{p.Infra.RemoteAccess, isSupportedInfraRemoteAccess, "infra remote access"},
		MetaParam{p.Product.Name, isSupportedProductName, "product name"},
		MetaParam{p.Product.License, isSupportedProductLicense, "product license"},
		MetaParam{p.Product.SupportLevel, isSupportedProductSupportLevel, "product support level"},
		MetaParam{p.Secrets.Provider, isSupportedSecretsProvider, "secrets provider"},
	}

	for _, pr := range params {
		if !pr.f(pr.key) {
			clog.Warningf(InvalidMetaField, ptf, pr.comment, pr.key)
			err = true
		}
	}

	for _, e := range p.Environments {
		if e.Kubernetes.Type != "" {
			if !isSupportedKubernetesProvider(e.Kubernetes.Type) {
				clog.Warningf(InvalidMetaField, ptf, "Kubernetes provider type", e.Kubernetes.Type)
				err = true
			}
		}
		if e.OS.Name != "" {
			if !isSupportedOS(e.OS.Name) {
				clog.Warningf(InvalidMetaField, ptf, "OS name", e.OS.Name)
				err = true
			}
		}
	}

	if err {
		return fmt.Errorf("invalid value")
	}

	return nil
}

func GetPlatformMetadata(ptf string, contents []byte) (PlatformMetadata, error) {
	var meta PlatformMetadata

	// unmarshal configuration
	err := yaml.Unmarshal(contents, &meta)
	if err != nil {
		e := fmt.Errorf("%s metadata: unable to unmarshal config (%s)", ptf, err)
		clog.Error(e)
		return meta, e
	}

	// check for valid configuration
	err = meta.IsValid(ptf)
	if err != nil {
		clog.Errorf("%s: metadata file seems to be invalid and with unsupported keys/values (%s)", ptf, err)
	}

	return meta, nil
}
