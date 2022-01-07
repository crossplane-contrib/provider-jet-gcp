package cloudplatform

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_project", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.TerraformResource.Schema["org_id"].Description =
			"The numeric ID of the organization this project belongs to."
	})
	p.AddResourceConfigurator("google_service_account_key", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		// Note(turkenh): We have to modify schema of "keepers", since it is a
		// map where elements configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["keepers"].Elem = schema.TypeString
	})
	p.AddResourceConfigurator("google_service_account", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.Kind = "ServiceAccount"
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["account_id"] = name
			},
			OmittedFields:     []string{"account_id"},
			GetExternalNameFn: config.IDAsExternalName,
			GetIDFn:           config.ExternalNameAsID,
		}
	})
}
