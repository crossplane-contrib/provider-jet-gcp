package config

import (
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-provider-google/google"

	"github.com/crossplane-contrib/provider-tf-gcp/config/cloudfunctions"
	"github.com/crossplane-contrib/provider-tf-gcp/config/cloudiot"
	"github.com/crossplane-contrib/provider-tf-gcp/config/compute"
	"github.com/crossplane-contrib/provider-tf-gcp/config/data"
	"github.com/crossplane-contrib/provider-tf-gcp/config/dataflow"
	"github.com/crossplane-contrib/provider-tf-gcp/config/dataproc"
	"github.com/crossplane-contrib/provider-tf-gcp/config/iam"
	"github.com/crossplane-contrib/provider-tf-gcp/config/monitoring"
	"github.com/crossplane-contrib/provider-tf-gcp/config/project"
	"github.com/crossplane-contrib/provider-tf-gcp/config/storage"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/crossplane-contrib/provider-tf-gcp"
)

var skipList = []string{
	// Note(turkenh): Following two resources conflicts their singular versions
	// "google_access_context_manager_access_level" and
	// "google_access_context_manager_service_perimeter". Skipping for now.
	"google_access_context_manager_access_levels$",
	"google_access_context_manager_service_perimeters$",
}

/*var includeList = []string{
	"google_compute_subnetwork$",
	"google_compute_router_nat$",
}*/

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

		if strings.HasPrefix(r.Name, "google_service") {
			r.Group = "iam"
		}
		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithGroupSuffix(".gcp.tf.crossplane.io"),
		tjconfig.WithShortName("tfgcp"),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		cloudfunctions.Configure,
		cloudiot.Configure,
		compute.Configure,
		data.Configure,
		dataflow.Configure,
		dataproc.Configure,
		iam.Configure,
		monitoring.Configure,
		project.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
