package data

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_data_catalog_entry_group", func(r *config.Resource) {
		// Note(turkenh): We have to override the default kind here,
		// which is "CatalogEntryGroup", since it conflicts otherwise
		// with: CatalogEntryGroupKind redeclared in this block
		r.Kind = "DataCatalogEntryGroup"
	})
}
