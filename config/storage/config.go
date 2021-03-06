package storage

import (
	"context"
	"path/filepath"
	"time"

	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_storage_bucket_object", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		// Note(turkenh): We have to modify schema of
		// "customer_encryption", since it is struct marked as sensitive.
		// https://github.com/crossplane/terrajet/issues/100#issuecomment-966892273
		r.TerraformResource.Schema["customer_encryption"].Sensitive = false
	})

	p.AddResourceConfigurator("google_storage_bucket", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			return filepath.Join(project, externalName), err
		}
		// Note(turkenh): Terraform provider added retries to bucket resource
		// for eventual consistency: https://github.com/hashicorp/terraform-provider-google/pull/10287
		// However, this causes refresh to keep retrying until timeout if the
		// bucket does not exist indeed. This causes the initial observe call to
		// hang on until timeout. We configure read timeout to a relatively
		// smaller value as a workaround/solution.
		// Related issue: https://github.com/crossplane-contrib/provider-jet-gcp/issues/12
		r.OperationTimeouts.Read = 1 * time.Minute
	})
}
