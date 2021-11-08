package config

import (
	"github.com/crossplane-contrib/provider-tf-gcp/config/compute"
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-provider-google/google"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/crossplane-contrib/provider-tf-gcp"
)

var includeList = []string{
	"google_storage_bucket$",

	// Compute
	"google_compute_network$",
	"google_compute_subnetwork$",
	"google_compute_address$",
	"google_compute_firewall$",
	// Had to skip "google_compute_instance" due to
	// https://github.com/crossplane-contrib/terrajet/issues/134
	//"google_compute_instance$",
	"google_compute_managed_ssl_certificate$",
	"google_compute_router$",
	// Had to skip "google_compute_router_nat" due to
	// https://github.com/crossplane-contrib/terrajet/issues/135
	//"google_compute_router_nat$",


}
// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	resourceMap := tf.Provider().ResourcesMap
	// Comment out the line below instead of the above, if your Terraform
	// provider uses an old version (<v2) of github.com/hashicorp/terraform-plugin-sdk.
	// resourceMap := conversion.GetV2ResourceMap(tf.Provider())

	defaultResourceFn := func(name string, terraformResource *schema.Resource) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// GCP Resources has id in a format that contains some other parameter
		// like projectID, e.g. projects/{{project}}/global/networks/{{name}}
		// So, we cannot generate external name using the provided config until
		// https://github.com/crossplane-contrib/terrajet/issues/119 resolved.
		r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithGroupSuffix(".gcp.tf.crossplane.io"),
		tjconfig.WithShortName("tfgcp"),
		tjconfig.WithIncludeList(includeList))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		compute.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
