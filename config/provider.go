package config

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-provider-google/google"

	"github.com/crossplane-contrib/provider-jet-gcp/config/accessapproval"
	"github.com/crossplane-contrib/provider-jet-gcp/config/bigtable"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudfunctions"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudiot"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudplatform"
	"github.com/crossplane-contrib/provider-jet-gcp/config/compute"
	"github.com/crossplane-contrib/provider-jet-gcp/config/dataflow"
	"github.com/crossplane-contrib/provider-jet-gcp/config/dataproc"
	"github.com/crossplane-contrib/provider-jet-gcp/config/identityplatform"
	"github.com/crossplane-contrib/provider-jet-gcp/config/project"
	"github.com/crossplane-contrib/provider-jet-gcp/config/storage"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/crossplane-contrib/provider-jet-gcp"
)

var skipList = []string{
	// Note(turkenh): Following two resources conflicts their singular versions
	// "google_access_context_manager_access_level" and
	// "google_access_context_manager_service_perimeter". Skipping for now.
	"google_access_context_manager_access_levels$",
	"google_access_context_manager_service_perimeters$",
}

var includeList = []string{
	// Storage
	"google_storage_bucket$",

	// Compute
	"google_compute_network$",
	"google_compute_subnetwork$",
	"google_compute_address$",
	"google_compute_firewall$",
	"google_compute_instance$",
	"google_compute_managed_ssl_certificate$",
	"google_compute_router$",
	"google_compute_router_nat$",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	resourceMap := tf.Provider().ResourcesMap
	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(DefaultResource(
			groupOverrides(),
			externalNameConfig(),
		)),
		tjconfig.WithRootGroup("gcp.jet.crossplane.io"),
		tjconfig.WithShortName("gcpjet"),
		// Comment out the following line to generate all resources.
		tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		accessapproval.Configure,
		bigtable.Configure,
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		compute.Configure,
		dataflow.Configure,
		dataproc.Configure,
		identityplatform.Configure,
		project.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// DefaultResource returns a DefaultResourceFn that makes sure the original
// DefaultResource call is made with given options here.
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
