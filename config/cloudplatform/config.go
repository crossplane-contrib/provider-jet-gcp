package cloudplatform

import (
	"context"
	"fmt"

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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			base["account_id"] = externalName
		}
		r.ExternalName.OmittedFields = []string{"account_id"}
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, err := common.GetField(tfstate, "account_id")
			if err != nil {
				return "", err
			}
			return id, nil
		}
		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", project, externalName, project), nil
		}
	})
}
