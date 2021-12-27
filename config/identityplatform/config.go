package identityplatform

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_identity_platform_default_supported_idp_config", func(r *config.Resource) {
		r.Kind = "DefaultSupportedIdPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_oauth_idp_config", func(r *config.Resource) {
		r.Kind = "OAuthIdPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_tenant_default_supported_idp_config", func(r *config.Resource) {
		r.Kind = "TenantDefaultSupportedIdPConfig"
	})
	p.AddResourceConfigurator("google_identity_platform_tenant_oauth_idp_config", func(r *config.Resource) {
		r.Kind = "TenantOAuthIdPConfig"
	})
}
