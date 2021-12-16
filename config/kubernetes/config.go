package kubernetes

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-gcp/apis/rconfig"
	"github.com/crossplane-contrib/provider-jet-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_container_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"cluster_ipv4_cidr", "ip_allocation_policy"},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			location, err := common.GetField(parameters, "location")
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, externalName), nil
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_container_node_pool", func(r *config.Resource) {
		locationIndex, clusterNameIndex := 3, 5
		r.Kind = "NodePool"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			project, err := common.GetField(providerConfig, common.KeyProject)
			if err != nil {
				return "", err
			}
			clusterID, err := common.GetField(parameters, "cluster")
			if err != nil {
				return "", err
			}
			location := strings.Split(clusterID, "/")[locationIndex]
			cluster := strings.Split(clusterID, "/")[clusterNameIndex]
			return fmt.Sprintf("%s/%s/%s/%s", project, location, cluster, externalName), nil
		}
		r.References["cluster"] = config.Reference{
			Type:      "Cluster",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})
}
