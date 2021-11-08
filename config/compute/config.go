package compute

import (
	"github.com/crossplane-contrib/provider-tf-gcp/config/common"
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_compute_managed_ssl_certificate", func(r *config.Resource) {
		r.Kind = "ManagedSSLCertificate"
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_compute_subnetwork", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_compute_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		r.References["subnetwork"] = config.Reference{
			Type: "Subnetwork",
		}
	})

	p.AddResourceConfigurator("google_compute_firewall", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
			Extractor: common.PathSelfLinkExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_router", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
			Extractor: common.PathSelfLinkExtractor,
		}
	})
	
	/*
	p.AddResourceConfigurator("google_compute_router_nat", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type: "Router",
		}
		r.References["subnetwork.name"] = config.Reference{
			Type: "Subnetwork",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.UseAsync = true
	})*/
}