package config

import (
	"strings"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/iancoleman/strcase"

	"github.com/crossplane-contrib/provider-jet-gcp/config/accessapproval"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudplatform"
)

func groupOverrides() tjconfig.ResourceOption { //nolint: gocyclo
	// Note(turkenh): The todo below will fix this linter error, disabling until
	//  then
	return func(r *tjconfig.Resource) {
		// Todo(turkenh): Use an single map for all resource => group mapping
		words := strings.Split(r.Name, "_")
		if _, ok := cloudplatform.Resources[r.Name]; ok {
			r.ShortGroup = "cloudplatform"
		} else if _, ok := accessapproval.Resources[r.Name]; ok { //nolint: gocritic
			// Todo(turkenh): Check how to fix linter "ifElseChain: rewrite if-else"
			//  to switch statement that covers all checks here.
			r.ShortGroup = "accessapproval"
		} else if strings.HasPrefix(r.Name, "google_access_context_manager") ||
			strings.HasPrefix(r.Name, "google_data_loss_prevention") {
			r.ShortGroup = words[1] + words[2] + words[3]
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

			r.ShortGroup = words[1] + words[2]
			r.Kind = strcase.ToCamel(strings.Join(words[3:], "_"))
		}
	}
}
