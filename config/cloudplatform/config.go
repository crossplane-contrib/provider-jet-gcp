package cloudplatform

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_billing_subaccount", func(r *config.Resource) {
		r.Kind = "BillingSubaccount"
	})
	p.AddResourceConfigurator("google_folder", func(r *config.Resource) {
		r.Kind = "Folder"
	})
	p.AddResourceConfigurator("google_folder_iam", func(r *config.Resource) {
		r.Kind = "FolderIAM"
	})
	p.AddResourceConfigurator("google_folder_organization_policy", func(r *config.Resource) {
		r.Kind = "FolderOrganizationPolicy"
	})
	p.AddResourceConfigurator("google_organization_iam", func(r *config.Resource) {
		r.Kind = "OrganizationIAM"
	})
	p.AddResourceConfigurator("google_organization_iam_custom_role", func(r *config.Resource) {
		r.Kind = "OrganizationIAMCustomRole"
	})
	p.AddResourceConfigurator("google_organization_policy", func(r *config.Resource) {
		r.Kind = "OrganizationPolicy"
	})
	p.AddResourceConfigurator("google_project", func(r *config.Resource) {
		r.Kind = "Project"
	})
	p.AddResourceConfigurator("google_project_default_service_accounts", func(r *config.Resource) {
		r.Kind = "ProjectDefaultServiceAccounts"
	})
	p.AddResourceConfigurator("google_project_iam", func(r *config.Resource) {
		r.Kind = "ProjectIAM"
	})
	p.AddResourceConfigurator("google_project_iam_custom_role", func(r *config.Resource) {
		r.Kind = "ProjectIAMCustomRole"
	})
	p.AddResourceConfigurator("google_project_organization_policy", func(r *config.Resource) {
		r.Kind = "ProjectOrganizationPolicy"
	})
	p.AddResourceConfigurator("google_project_service", func(r *config.Resource) {
		r.Kind = "ProjectService"
	})
	p.AddResourceConfigurator("google_service_account", func(r *config.Resource) {
		r.Kind = "ServiceAccount"
	})
	p.AddResourceConfigurator("google_service_account_iam", func(r *config.Resource) {
		r.Kind = "ServiceAccountIAM"
	})
	p.AddResourceConfigurator("google_service_account_key", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of "keepers", since it is a
		// map where elements configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["keepers"].Elem = schema.TypeString
		r.Kind = "ServiceAccountKey"
	})
	p.AddResourceConfigurator("google_service_networking_peered_dns_domain", func(r *config.Resource) {
		r.Kind = "ServiceNetworkingPeeredDNSDomain"
	})
	p.AddResourceConfigurator("google_project_service_identity", func(r *config.Resource) {
		r.Kind = "ProjectServiceIdentity"
	})

	p.AddResourceConfigurator("google_organization_iam_audit_config", func(r *config.Resource) {
		r.Kind = "OrganizationIAMAuditConfig"
	})
	p.AddResourceConfigurator("google_organization_iam_binding", func(r *config.Resource) {
		r.Kind = "OrganizationIAMBinding"
	})
	p.AddResourceConfigurator("google_organization_iam_member", func(r *config.Resource) {
		r.Kind = "OrganizationIAMMember"
	})
	p.AddResourceConfigurator("google_organization_iam_policy", func(r *config.Resource) {
		r.Kind = "OrganizationIAMPolicy"
	})

	p.AddResourceConfigurator("google_project_iam_audit_config", func(r *config.Resource) {
		r.Kind = "ProjectIAMAuditConfig"
	})
	p.AddResourceConfigurator("google_project_iam_binding", func(r *config.Resource) {
		r.Kind = "ProjectIAMBinding"
	})
	p.AddResourceConfigurator("google_project_iam_member", func(r *config.Resource) {
		r.Kind = "ProjectIAMMember"
	})
	p.AddResourceConfigurator("google_project_iam_policy", func(r *config.Resource) {
		r.Kind = "ProjectIAMPolicy"
	})
	p.AddResourceConfigurator("google_project_usage_export_bucket", func(r *config.Resource) {
		r.Kind = "ProjectUsageExportBucket"
	})

	p.AddResourceConfigurator("google_folder_iam_audit_config", func(r *config.Resource) {
		r.Kind = "FolderIAMAuditConfig"
	})
	p.AddResourceConfigurator("google_folder_iam_binding", func(r *config.Resource) {
		r.Kind = "FolderIAMBinding"
	})
	p.AddResourceConfigurator("google_folder_iam_member", func(r *config.Resource) {
		r.Kind = "FolderIAMMember"
	})
	p.AddResourceConfigurator("google_folder_iam_policy", func(r *config.Resource) {
		r.Kind = "FolderIAMPolicy"
	})

	p.AddResourceConfigurator("google_service_account_iam_binding", func(r *config.Resource) {
		r.Kind = "ServiceAccountIAMBinding"
	})
	p.AddResourceConfigurator("google_service_account_iam_member", func(r *config.Resource) {
		r.Kind = "ServiceAccountIAMMember"
	})
	p.AddResourceConfigurator("google_service_account_iam_policy", func(r *config.Resource) {
		r.Kind = "ServiceAccountIAMPolicy"
	})
}

// Resources in "Cloud Platform" group.
// Note(turkenh): The following resources are listed under "Cloud Platform"
// section in Terraform Documentation.
var Resources = map[string]struct{}{
	"google_billing_subaccount":                   {},
	"google_folder":                               {},
	"google_folder_iam":                           {},
	"google_folder_organization_policy":           {},
	"google_organization_iam":                     {},
	"google_organization_iam_custom_role":         {},
	"google_organization_policy":                  {},
	"google_project":                              {},
	"google_project_default_service_accounts":     {},
	"google_project_iam":                          {},
	"google_project_iam_custom_role":              {},
	"google_project_organization_policy":          {},
	"google_project_service":                      {},
	"google_service_account":                      {},
	"google_service_account_iam":                  {},
	"google_service_account_key":                  {},
	"google_service_networking_peered_dns_domain": {},
	"google_project_service_identity":             {},

	"google_organization_iam_audit_config": {},
	"google_organization_iam_binding":      {},
	"google_organization_iam_member":       {},
	"google_organization_iam_policy":       {},

	"google_project_iam_audit_config":    {},
	"google_project_iam_binding":         {},
	"google_project_iam_member":          {},
	"google_project_iam_policy":          {},
	"google_project_usage_export_bucket": {},

	"google_folder_iam_audit_config": {},
	"google_folder_iam_binding":      {},
	"google_folder_iam_member":       {},
	"google_folder_iam_policy":       {},

	"google_service_account_iam_binding": {},
	"google_service_account_iam_member":  {},
	"google_service_account_iam_policy":  {},
}
