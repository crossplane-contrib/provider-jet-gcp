module github.com/crossplane-contrib/provider-jet-gcp

go 1.16

replace github.com/crossplane/terrajet => ../../crossplane/terrajet

require (
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/terrajet v0.3.1
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/terraform-json v0.13.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1 // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)
