package identityplatform

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_identity_platform_default_supported_idp_config", func(r *config.Resource) {
		r.Kind = "DefaultSupportedIDPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_oauth_idp_config", func(r *config.Resource) {
		r.Kind = "OAuthIDPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_tenant_default_supported_idp_config", func(r *config.Resource) {
		r.Kind = "TenantDefaultSupportedIDPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_tenant_oauth_idp_config", func(r *config.Resource) {
		r.Kind = "TenantOAuthIDPConfig"
	})
}
