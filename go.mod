module github.com/crossplane-contrib/provider-jet-gcp

go 1.16

require (
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/terrajet v0.3.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	// Commit for v4.0.0 tag https://github.com/hashicorp/terraform-provider-google/commit/f004d2d203fa86da2e344fc23dc6fe509e6bbfae
	github.com/hashicorp/terraform-provider-google v1.20.1-0.20211102210101-f004d2d203fa
	github.com/pkg/errors v0.9.1
	go.uber.org/multierr v1.7.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)
