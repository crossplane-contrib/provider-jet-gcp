package accessapproval

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_folder_access_approval_settings", func(r *config.Resource) {
		r.Kind = "FolderSettings"
	})
	p.AddResourceConfigurator("google_organization_access_approval_settings", func(r *config.Resource) {
		r.Kind = "OrganizationSettings"
	})
	p.AddResourceConfigurator("google_project_access_approval_settings", func(r *config.Resource) {
		r.Kind = "ProjectSettings"
	})
}

// Resources in "Access Approval" group.
// Note(turkenh): The following resources are listed under "Access Approval"
// section in Terraform Documentation.
var Resources = map[string]struct{}{
	"google_folder_access_approval_settings":       {},
	"google_organization_access_approval_settings": {},
	"google_project_access_approval_settings":      {},
}
