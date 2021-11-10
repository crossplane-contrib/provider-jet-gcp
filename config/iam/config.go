package iam

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_service_account_key", func(r *config.Resource) {
		r.Kind = "ServiceAccountKey"
		// Note(turkenh): We have to modify schema of "keepers", since it is a
		// map where elements configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["keepers"].Elem = schema.TypeString
	})

	p.AddResourceConfigurator("google_service_account_iam", func(r *config.Resource) {
		r.Kind = "ServiceAccountIAM"
	})

	p.AddResourceConfigurator("google_service_account", func(r *config.Resource) {
		r.Kind = "ServiceAccount"
	})
}
