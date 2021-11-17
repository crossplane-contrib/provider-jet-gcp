package config

import (
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-provider-google/google"
	"github.com/iancoleman/strcase"

	"github.com/crossplane-contrib/provider-tf-gcp/config/accessapproval"
	"github.com/crossplane-contrib/provider-tf-gcp/config/cloudfunctions"
	"github.com/crossplane-contrib/provider-tf-gcp/config/cloudiot"
	"github.com/crossplane-contrib/provider-tf-gcp/config/cloudplatform"
	"github.com/crossplane-contrib/provider-tf-gcp/config/compute"
	"github.com/crossplane-contrib/provider-tf-gcp/config/dataflow"
	"github.com/crossplane-contrib/provider-tf-gcp/config/dataproc"
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

		words := strings.Split(r.Name, "_")
		if _, ok := cloudplatform.Resources[r.Name]; ok {
			r.Group = "cloudplatform"
		} else if _, ok := accessapproval.Resources[r.Name]; ok {
			r.Group = "accessapproval"
		} else if strings.HasPrefix(r.Name, "google_access_context_manager") ||
			strings.HasPrefix(r.Name, "google_data_loss_prevention") {
			r.Group = words[1] + words[2] + words[3]
			r.Kind = strcase.ToCamel(strings.Join(words[4:], "_"))
		} else if strings.HasPrefix(r.Name, "google_active_directory") ||
			strings.HasPrefix(r.Name, "google_app_engine") ||
			strings.HasPrefix(r.Name, "google_assured_workloads") ||
			strings.HasPrefix(r.Name, "google_binary_authorization") ||
			strings.HasPrefix(r.Name, "google_deployment_manager") ||
			strings.HasPrefix(r.Name, "google_essential_contacts") ||
			strings.HasPrefix(r.Name, "google_game_services") ||
			strings.HasPrefix(r.Name, "google_gke_hub") ||
			strings.HasPrefix(r.Name, "google_identity_platform") ||
			strings.HasPrefix(r.Name, "google_ml_engine") ||
			strings.HasPrefix(r.Name, "google_network_management") ||
			strings.HasPrefix(r.Name, "google_network_services") ||
			strings.HasPrefix(r.Name, "google_service_networking") ||
			strings.HasPrefix(r.Name, "google_resource_manager") ||
			strings.HasPrefix(r.Name, "google_secret_manager") ||
			strings.HasPrefix(r.Name, "google_org_policy") ||
			strings.HasPrefix(r.Name, "google_vpc_access_connector") ||
			// Examples: google_cloud_identity, google_cloud_run,
			// google_cloud_asset, google_data_fusion, google_data_source...
			strings.HasPrefix(r.Name, "google_cloud_") || strings.HasPrefix(r.Name, "google_data_") {

			r.Group = words[1] + words[2]
			r.Kind = strcase.ToCamel(strings.Join(words[3:], "_"))
		}
		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithGroupSuffix(".gcp.tf.crossplane.io"),
		tjconfig.WithShortName("tfgcp"),
		// Comment out the following line to generate all resources.
		tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		accessapproval.Configure,
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		compute.Configure,
		dataflow.Configure,
		dataproc.Configure,
		project.Configure,
		storage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
