package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-tf-gcp/config/common"
)

var (
	// PathSelfLinkExtractor is the golang path to SelfLinkExtractor function
	// in this package.
	PathSelfLinkExtractor = SelfPackagePath + ".SelfLinkExtractor()"
)

// SelfLinkExtractor extracts URI of the resources from
// "status.atProvider.selfLink" which is quite common among all GCP resources.
func SelfLinkExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.selfLink")
		if err != nil {
			return ""
		}
		return r
	}
}
