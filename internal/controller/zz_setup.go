/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	address "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/address"
	firewall "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/firewall"
	instance "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instance"
	managedsslcertificate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/network"
	router "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/router"
	routernat "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/routernat"
	subnetwork "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/subnetwork"
	providerconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/providerconfig"
	bucket "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		address.Setup,
		firewall.Setup,
		instance.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		router.Setup,
		routernat.Setup,
		subnetwork.Setup,
		providerconfig.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}