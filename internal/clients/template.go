package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-tf-template/apis/v1alpha1"
)

const (
	keyUsername = "username"
	keyPassword = "password"
	keyHost     = "host"

	// Template credentials environment variable names
	envUsername = "HASHICUPS_USERNAME"
	envPassword = "HASHICUPS_PASSWORD"
)

const (
	fmtEnvVar = "%s=%s"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal template credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		templateCreds := map[string]string{}
		if err := json.Unmarshal(data, &templateCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{
			"host": templateCreds[keyHost],
		}
		// set environment variables for sensitive provider configuration
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envUsername, templateCreds[keyUsername]),
			fmt.Sprintf(fmtEnvVar, envPassword, templateCreds[keyPassword]),
		}
		return ps, nil
	}
}
