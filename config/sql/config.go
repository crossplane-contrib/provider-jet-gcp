package sql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_sql_database_instance", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/instances/%s", project, externalName), nil
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_sql_database", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, externalName), nil
		}

		r.References["instance"] = config.Reference{
			Type: "DatabaseInstance",
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_sql_source_representation_instance", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, externalName), nil
		}

		r.References["instance"] = config.Reference{
			Type: "DatabaseInstance",
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_sql_user", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"].(string)
			if !ok {
				return "", errors.New("cannot get the id field as string in tfstate")
			}
			words := strings.Split(id, "/")
			return words[0], nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}
			host, err := common.GetField(parameters, "host")
			if err != nil {
				host = ""
			}

			return fmt.Sprintf("%s/%s/%s", externalName, host, instance), nil
		}

		r.References["instance"] = config.Reference{
			Type: "DatabaseInstance",
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_sql_ssl_cert", func(r *config.Resource) {
		r.Version = common.VersionV1alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			if externalName == "" {
				return "", nil
			}
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/instances/%s/sslCerts/%s", project, instance, externalName), nil
		}

		r.References["instance"] = config.Reference{
			Type: "DatabaseInstance",
		}

		r.UseAsync = true
	})
}
