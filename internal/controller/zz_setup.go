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

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	foldersettings "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accessapproval/foldersettings"
	organizationsettings "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accessapproval/organizationsettings"
	projectsettings "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accessapproval/projectsettings"
	accesslevel "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/accesspolicy"
	gcpuseraccessbinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/gcpuseraccessbinding"
	serviceperimeter "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/accesscontextmanager/serviceperimeterresource"
	domain "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/activedirectory/domain"
	domaintrust "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/activedirectory/domaintrust"
	envgroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/envgroup"
	envgroupattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/envgroupattachment"
	environment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/environment"
	instance "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/instance"
	instanceattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/instanceattachment"
	organization "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/apigee/organization"
	application "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/application"
	applicationurldispatchrules "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/applicationurldispatchrules"
	domainmapping "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/domainmapping"
	firewallrule "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/firewallrule"
	flexibleappversion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/flexibleappversion"
	servicenetworksettings "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/servicenetworksettings"
	servicesplittraffic "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/servicesplittraffic"
	standardappversion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/appengine/standardappversion"
	workload "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/assuredworkloads/workload"
	dataset "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/dataset"
	datasetaccess "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/datasetaccess"
	datasetiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/datasetiambinding"
	datasetiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/datasetiammember"
	datasetiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/datasetiampolicy"
	datatransferconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/datatransferconfig"
	job "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/job"
	reservation "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/reservation"
	routine "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/routine"
	table "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/table"
	tableiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/tableiambinding"
	tableiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/tableiammember"
	tableiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigquery/tableiampolicy"
	appprofile "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/appprofile"
	garbagecollectionpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/garbagecollectionpolicy"
	instancebigtable "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/instance"
	instanceiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/instanceiambinding"
	instanceiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/instanceiammember"
	instanceiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/instanceiampolicy"
	tablebigtable "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/table"
	tableiambindingbigtable "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/tableiambinding"
	tableiammemberbigtable "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/tableiammember"
	tableiampolicybigtable "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/bigtable/tableiampolicy"
	accountiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/billing/accountiambinding"
	accountiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/billing/accountiammember"
	accountiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/billing/accountiampolicy"
	budget "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/billing/budget"
	attestor "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/binaryauthorization/attestor"
	attestoriambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/binaryauthorization/attestoriambinding"
	attestoriammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/binaryauthorization/attestoriammember"
	attestoriampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/binaryauthorization/attestoriampolicy"
	policy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/binaryauthorization/policy"
	folderfeed "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudasset/folderfeed"
	organizationfeed "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudasset/organizationfeed"
	projectfeed "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudasset/projectfeed"
	trigger "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudbuild/trigger"
	function "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudfunctions/function"
	functioniambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudfunctions/functioniambinding"
	functioniammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudfunctions/functioniammember"
	functioniampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudfunctions/functioniampolicy"
	group "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudidentity/group"
	groupmembership "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudidentity/groupmembership"
	device "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudiot/device"
	registry "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudiot/registry"
	billingsubaccount "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/billingsubaccount"
	folder "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folder"
	folderiamauditconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folderiamauditconfig"
	folderiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folderiambinding"
	folderiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folderiammember"
	folderiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folderiampolicy"
	folderorganizationpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/folderorganizationpolicy"
	organizationiamauditconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationiamauditconfig"
	organizationiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationiambinding"
	organizationiamcustomrole "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationiammember"
	organizationiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationiampolicy"
	organizationpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/organizationpolicy"
	project "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectiambinding"
	projectiamcustomrole "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectiamcustomrole"
	projectiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectiammember"
	projectiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectiampolicy"
	projectorganizationpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectorganizationpolicy"
	projectservice "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectservice"
	projectusageexportbucket "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/serviceaccountiambinding"
	serviceaccountiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/serviceaccountiampolicy"
	serviceaccountkey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmappingcloudrun "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudrun/domainmapping"
	service "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudrun/service"
	serviceiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudrun/serviceiambinding"
	serviceiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudrun/serviceiammember"
	serviceiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudrun/serviceiampolicy"
	jobcloudscheduler "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudscheduler/job"
	queue "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/cloudtasks/queue"
	environmentcomposer "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/composer/environment"
	address "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/address"
	attacheddisk "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/attacheddisk"
	autoscaler "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/autoscaler"
	backendbucket "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/backendbucket"
	backendbucketsignedurlkey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/backendbucketsignedurlkey"
	backendservice "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/backendservice"
	backendservicesignedurlkey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/backendservicesignedurlkey"
	disk "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/disk"
	diskiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/diskiambinding"
	diskiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/diskiammember"
	diskiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/diskiampolicy"
	diskresourcepolicyattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/externalvpngateway"
	firewall "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/firewall"
	firewallpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/firewallpolicy"
	firewallpolicyassociation "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/firewallpolicyassociation"
	firewallpolicyrule "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/firewallpolicyrule"
	forwardingrule "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/forwardingrule"
	globaladdress "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/globalforwardingrule"
	globalnetworkendpoint "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/globalnetworkendpointgroup"
	havpngateway "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/havpngateway"
	healthcheck "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/httpshealthcheck"
	image "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/image"
	imageiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/imageiambinding"
	imageiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/imageiammember"
	imageiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/imageiampolicy"
	instancecompute "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instance"
	instancefromtemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instancefromtemplate"
	instancegroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instancegroup"
	instancegroupmanager "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instancegroupmanager"
	instancegroupnamedport "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instancegroupnamedport"
	instanceiambindingcompute "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instanceiambinding"
	instanceiammembercompute "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instanceiammember"
	instanceiampolicycompute "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instanceiampolicy"
	instancetemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/instancetemplate"
	interconnectattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/interconnectattachment"
	managedsslcertificate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/network"
	networkendpoint "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/networkendpoint"
	networkendpointgroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/networkendpointgroup"
	networkpeering "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/networkpeering"
	networkpeeringroutesconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/networkpeeringroutesconfig"
	nodegroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/nodegroup"
	nodetemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/nodetemplate"
	packetmirroring "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/packetmirroring"
	perinstanceconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/perinstanceconfig"
	projectdefaultnetworktier "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/projectdefaultnetworktier"
	projectmetadata "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/projectmetadata"
	projectmetadataitem "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/projectmetadataitem"
	regionautoscaler "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionautoscaler"
	regionbackendservice "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionbackendservice"
	regiondisk "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiondisk"
	regiondiskiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiondiskiambinding"
	regiondiskiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiondiskiammember"
	regiondiskiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiondiskiampolicy"
	regiondiskresourcepolicyattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiondiskresourcepolicyattachment"
	regionhealthcheck "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regioninstancegroupmanager"
	regionnetworkendpointgroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionnetworkendpointgroup"
	regionperinstanceconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionperinstanceconfig"
	regionsslcertificate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/regionurlmap"
	reservationcompute "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/reservation"
	resourcepolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/resourcepolicy"
	route "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/route"
	router "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/router"
	routerinterface "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/routerinterface"
	routernat "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/routernat"
	routerpeer "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/routerpeer"
	securitypolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/securitypolicy"
	serviceattachment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/serviceattachment"
	sharedvpchostproject "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/sharedvpchostproject"
	sharedvpcserviceproject "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/sharedvpcserviceproject"
	snapshot "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/snapshot"
	sslcertificate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/sslcertificate"
	sslpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/sslpolicy"
	subnetwork "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/subnetwork"
	subnetworkiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/subnetworkiambinding"
	subnetworkiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/subnetworkiammember"
	subnetworkiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/subnetworkiampolicy"
	targetgrpcproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targetgrpcproxy"
	targethttpproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targethttpproxy"
	targethttpsproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targethttpsproxy"
	targetinstance "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targetinstance"
	targetpool "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targetpool"
	targetsslproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targetsslproxy"
	targettcpproxy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/targettcpproxy"
	urlmap "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/urlmap"
	vpngateway "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/vpngateway"
	vpntunnel "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/compute/vpntunnel"
	cluster "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/container/cluster"
	nodepool "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/container/nodepool"
	registrycontainer "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/container/registry"
	note "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/containeranalysis/note"
	occurrence "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/containeranalysis/occurrence"
	entry "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/entry"
	entrygroup "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/entrygroup"
	entrygroupiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/entrygroupiambinding"
	entrygroupiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/entrygroupiammember"
	entrygroupiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/entrygroupiampolicy"
	tag "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/tag"
	tagtemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/tagtemplate"
	tagtemplateiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/tagtemplateiambinding"
	tagtemplateiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/tagtemplateiammember"
	tagtemplateiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datacatalog/tagtemplateiampolicy"
	jobdataflow "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataflow/job"
	deidentifytemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datalossprevention/deidentifytemplate"
	inspecttemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datalossprevention/inspecttemplate"
	jobtrigger "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datalossprevention/jobtrigger"
	storedinfotype "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datalossprevention/storedinfotype"
	autoscalingpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/autoscalingpolicy"
	clusterdataproc "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/cluster"
	clusteriambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/clusteriambinding"
	clusteriammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/clusteriammember"
	clusteriampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/clusteriampolicy"
	jobdataproc "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/job"
	jobiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/jobiambinding"
	jobiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/jobiammember"
	jobiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/jobiampolicy"
	workflowtemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dataproc/workflowtemplate"
	index "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/datastore/index"
	deployment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/deploymentmanager/deployment"
	agent "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflow/agent"
	entitytype "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflow/entitytype"
	fulfillment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflow/fulfillment"
	intent "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflow/intent"
	agentdialogflowcx "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/agent"
	entitytypedialogflowcx "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/entitytype"
	environmentdialogflowcx "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/environment"
	flow "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/flow"
	intentdialogflowcx "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/intent"
	page "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/page"
	version "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dialogflowcx/version"
	managedzone "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dns/managedzone"
	policydns "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dns/policy"
	recordset "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/dns/recordset"
	serviceendpoints "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/endpoints/service"
	serviceiambindingendpoints "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/endpoints/serviceiambinding"
	serviceiammemberendpoints "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/endpoints/serviceiammember"
	serviceiampolicyendpoints "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/endpoints/serviceiampolicy"
	contact "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/essentialcontacts/contact"
	triggereventarc "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/eventarc/trigger"
	instancefilestore "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/filestore/instance"
	document "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/firestore/document"
	indexfirestore "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/firestore/index"
	gameservercluster "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gameservices/gameservercluster"
	gameserverconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gameservices/gameserverconfig"
	gameserverdeployment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gameservices/gameserverdeployment"
	gameserverdeploymentrollout "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gameservices/gameserverdeploymentrollout"
	realm "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gameservices/realm"
	membership "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/gkehub/membership"
	consentstore "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/consentstore"
	consentstoreiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/consentstoreiambinding"
	consentstoreiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/consentstoreiammember"
	consentstoreiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/consentstoreiampolicy"
	datasethealthcare "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/dataset"
	datasetiambindinghealthcare "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/datasetiambinding"
	datasetiammemberhealthcare "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/datasetiammember"
	datasetiampolicyhealthcare "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/datasetiampolicy"
	dicomstore "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/dicomstore"
	dicomstoreiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/dicomstoreiambinding"
	dicomstoreiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/dicomstoreiammember"
	dicomstoreiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/dicomstoreiampolicy"
	fhirstore "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/fhirstore"
	fhirstoreiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/fhirstoreiambinding"
	fhirstoreiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/fhirstoreiammember"
	fhirstoreiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/fhirstoreiampolicy"
	hl7v2store "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/hl7v2store"
	hl7v2storeiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/hl7v2storeiambinding"
	hl7v2storeiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/hl7v2storeiammember"
	hl7v2storeiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/healthcare/hl7v2storeiampolicy"
	appengineserviceiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineserviceiambinding"
	appengineserviceiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineserviceiammember"
	appengineserviceiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineserviceiampolicy"
	appengineversioniambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineversioniambinding"
	appengineversioniammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineversioniammember"
	appengineversioniampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/appengineversioniampolicy"
	brand "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/brand"
	client "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/client"
	tunneliambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunneliambinding"
	tunneliammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunneliammember"
	tunneliampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunneliampolicy"
	tunnelinstanceiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunnelinstanceiambinding"
	tunnelinstanceiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunnelinstanceiammember"
	tunnelinstanceiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/tunnelinstanceiampolicy"
	webbackendserviceiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webbackendserviceiambinding"
	webbackendserviceiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webbackendserviceiammember"
	webbackendserviceiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webbackendserviceiampolicy"
	webiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webiambinding"
	webiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webiammember"
	webiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webiampolicy"
	webtypeappengineiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypeappengineiambinding"
	webtypeappengineiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypeappengineiammember"
	webtypeappengineiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypeappengineiampolicy"
	webtypecomputeiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypecomputeiambinding"
	webtypecomputeiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypecomputeiammember"
	webtypecomputeiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/iap/webtypecomputeiampolicy"
	defaultsupportedidpconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/oauthidpconfig"
	tenant "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/cryptokey"
	cryptokeyiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/cryptokeyiambinding"
	cryptokeyiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/cryptokeyiammember"
	cryptokeyiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/cryptokeyiampolicy"
	keyring "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/keyring"
	keyringiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/keyringiambinding"
	keyringiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/keyringiammember"
	keyringiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/keyringiampolicy"
	keyringimportjob "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/kms/secretciphertext"
	billingaccountbucketconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/billingaccountbucketconfig"
	billingaccountexclusion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/billingaccountexclusion"
	billingaccountsink "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/billingaccountsink"
	folderbucketconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/folderbucketconfig"
	folderexclusion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/folderexclusion"
	foldersink "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/foldersink"
	metric "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/metric"
	organizationbucketconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/organizationbucketconfig"
	organizationexclusion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/organizationexclusion"
	organizationsink "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/organizationsink"
	projectbucketconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/projectbucketconfig"
	projectexclusion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/projectexclusion"
	projectsink "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/logging/projectsink"
	instancememcache "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/memcache/instance"
	model "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/mlengine/model"
	alertpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/alertpolicy"
	customservice "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/customservice"
	dashboard "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/dashboard"
	groupmonitoring "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/group"
	metricdescriptor "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/metricdescriptor"
	notificationchannel "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/notificationchannel"
	slo "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/slo"
	uptimecheckconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/monitoring/uptimecheckconfig"
	connectivitytest "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/networkmanagement/connectivitytest"
	edgecachekeyset "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/networkservices/edgecachekeyset"
	edgecacheorigin "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/networkservices/edgecacheorigin"
	edgecacheservice "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/networkservices/edgecacheservice"
	environmentnotebooks "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/environment"
	instancenotebooks "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/instance"
	instanceiambindingnotebooks "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/instanceiambinding"
	instanceiammembernotebooks "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/instanceiammember"
	instanceiampolicynotebooks "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/instanceiampolicy"
	location "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/notebooks/location"
	policyorgpolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/orgpolicy/policy"
	patchdeployment "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/osconfig/patchdeployment"
	sshpublickey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/oslogin/sshpublickey"
	capool "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/capool"
	capooliambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/capooliambinding"
	capooliammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/capooliammember"
	capooliampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/capooliampolicy"
	certificate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/certificate"
	certificateauthority "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/certificateauthority"
	certificatetemplate "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/privateca/certificatetemplate"
	providerconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/providerconfig"
	litereservation "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/litereservation"
	litesubscription "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/litesubscription"
	litetopic "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/litetopic"
	schema "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/schema"
	subscription "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/subscription"
	subscriptioniambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/subscriptioniambinding"
	subscriptioniammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/subscriptioniammember"
	subscriptioniampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/subscriptioniampolicy"
	topic "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/topic"
	topiciambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/topiciambinding"
	topiciammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/topiciammember"
	topiciampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/pubsub/topiciampolicy"
	instanceredis "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/redis/instance"
	lien "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/resourcemanager/lien"
	notificationconfig "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/scc/notificationconfig"
	source "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/scc/source"
	secret "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/secretmanager/secret"
	secretiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/secretmanager/secretiambinding"
	secretiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/secretmanager/secretiammember"
	secretiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/secretmanager/secretiampolicy"
	secretversion "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/secretmanager/secretversion"
	connection "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/servicenetworking/connection"
	repository "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sourcerepo/repository"
	repositoryiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sourcerepo/repositoryiambinding"
	repositoryiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sourcerepo/repositoryiammember"
	repositoryiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sourcerepo/repositoryiampolicy"
	database "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/database"
	databaseiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/databaseiambinding"
	databaseiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/databaseiammember"
	databaseiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/databaseiampolicy"
	instancespanner "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/instance"
	instanceiambindingspanner "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/instanceiambinding"
	instanceiammemberspanner "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/instanceiammember"
	instanceiampolicyspanner "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/spanner/instanceiampolicy"
	databasesql "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sql/database"
	databaseinstance "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sql/sslcert"
	user "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/sql/user"
	bucket "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketaccesscontrol"
	bucketacl "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketacl"
	bucketiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketiambinding"
	bucketiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketiammember"
	bucketiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketiampolicy"
	bucketobject "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/defaultobjectacl"
	hmackey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/hmackey"
	notification "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/notification"
	objectaccesscontrol "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/objectaccesscontrol"
	objectacl "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storage/objectacl"
	jobstoragetransfer "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/storagetransfer/job"
	tagbinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagbinding"
	tagkey "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagkey"
	tagkeyiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagkeyiambinding"
	tagkeyiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagkeyiammember"
	tagkeyiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagkeyiampolicy"
	tagvalue "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagvalue"
	tagvalueiambinding "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagvalueiambinding"
	tagvalueiammember "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagvalueiammember"
	tagvalueiampolicy "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tags/tagvalueiampolicy"
	node "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/tpu/node"
	datasetvertexai "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/vertexai/dataset"
	connector "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/vpcaccess/connector"
	workflow "github.com/crossplane-contrib/provider-jet-gcp/internal/controller/workflows/workflow"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		foldersettings.Setup,
		organizationsettings.Setup,
		projectsettings.Setup,
		accesslevel.Setup,
		accesslevelcondition.Setup,
		accesspolicy.Setup,
		gcpuseraccessbinding.Setup,
		serviceperimeter.Setup,
		serviceperimeterresource.Setup,
		domain.Setup,
		domaintrust.Setup,
		envgroup.Setup,
		envgroupattachment.Setup,
		environment.Setup,
		instance.Setup,
		instanceattachment.Setup,
		organization.Setup,
		application.Setup,
		applicationurldispatchrules.Setup,
		domainmapping.Setup,
		firewallrule.Setup,
		flexibleappversion.Setup,
		servicenetworksettings.Setup,
		servicesplittraffic.Setup,
		standardappversion.Setup,
		workload.Setup,
		dataset.Setup,
		datasetaccess.Setup,
		datasetiambinding.Setup,
		datasetiammember.Setup,
		datasetiampolicy.Setup,
		datatransferconfig.Setup,
		job.Setup,
		reservation.Setup,
		routine.Setup,
		table.Setup,
		tableiambinding.Setup,
		tableiammember.Setup,
		tableiampolicy.Setup,
		appprofile.Setup,
		garbagecollectionpolicy.Setup,
		instancebigtable.Setup,
		instanceiambinding.Setup,
		instanceiammember.Setup,
		instanceiampolicy.Setup,
		tablebigtable.Setup,
		tableiambindingbigtable.Setup,
		tableiammemberbigtable.Setup,
		tableiampolicybigtable.Setup,
		accountiambinding.Setup,
		accountiammember.Setup,
		accountiampolicy.Setup,
		budget.Setup,
		attestor.Setup,
		attestoriambinding.Setup,
		attestoriammember.Setup,
		attestoriampolicy.Setup,
		policy.Setup,
		folderfeed.Setup,
		organizationfeed.Setup,
		projectfeed.Setup,
		trigger.Setup,
		function.Setup,
		functioniambinding.Setup,
		functioniammember.Setup,
		functioniampolicy.Setup,
		group.Setup,
		groupmembership.Setup,
		device.Setup,
		registry.Setup,
		billingsubaccount.Setup,
		folder.Setup,
		folderiamauditconfig.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		folderiampolicy.Setup,
		folderorganizationpolicy.Setup,
		organizationiamauditconfig.Setup,
		organizationiambinding.Setup,
		organizationiamcustomrole.Setup,
		organizationiammember.Setup,
		organizationiampolicy.Setup,
		organizationpolicy.Setup,
		project.Setup,
		projectdefaultserviceaccounts.Setup,
		projectiamauditconfig.Setup,
		projectiambinding.Setup,
		projectiamcustomrole.Setup,
		projectiammember.Setup,
		projectiampolicy.Setup,
		projectorganizationpolicy.Setup,
		projectservice.Setup,
		projectusageexportbucket.Setup,
		serviceaccount.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
		domainmappingcloudrun.Setup,
		service.Setup,
		serviceiambinding.Setup,
		serviceiammember.Setup,
		serviceiampolicy.Setup,
		jobcloudscheduler.Setup,
		queue.Setup,
		environmentcomposer.Setup,
		address.Setup,
		attacheddisk.Setup,
		autoscaler.Setup,
		backendbucket.Setup,
		backendbucketsignedurlkey.Setup,
		backendservice.Setup,
		backendservicesignedurlkey.Setup,
		disk.Setup,
		diskiambinding.Setup,
		diskiammember.Setup,
		diskiampolicy.Setup,
		diskresourcepolicyattachment.Setup,
		externalvpngateway.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		firewallpolicyassociation.Setup,
		firewallpolicyrule.Setup,
		forwardingrule.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		globalnetworkendpoint.Setup,
		globalnetworkendpointgroup.Setup,
		havpngateway.Setup,
		healthcheck.Setup,
		httphealthcheck.Setup,
		httpshealthcheck.Setup,
		image.Setup,
		imageiambinding.Setup,
		imageiammember.Setup,
		imageiampolicy.Setup,
		instancecompute.Setup,
		instancefromtemplate.Setup,
		instancegroup.Setup,
		instancegroupmanager.Setup,
		instancegroupnamedport.Setup,
		instanceiambindingcompute.Setup,
		instanceiammembercompute.Setup,
		instanceiampolicycompute.Setup,
		instancetemplate.Setup,
		interconnectattachment.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		networkendpoint.Setup,
		networkendpointgroup.Setup,
		networkpeering.Setup,
		networkpeeringroutesconfig.Setup,
		nodegroup.Setup,
		nodetemplate.Setup,
		packetmirroring.Setup,
		perinstanceconfig.Setup,
		projectdefaultnetworktier.Setup,
		projectmetadata.Setup,
		projectmetadataitem.Setup,
		regionautoscaler.Setup,
		regionbackendservice.Setup,
		regiondisk.Setup,
		regiondiskiambinding.Setup,
		regiondiskiammember.Setup,
		regiondiskiampolicy.Setup,
		regiondiskresourcepolicyattachment.Setup,
		regionhealthcheck.Setup,
		regioninstancegroupmanager.Setup,
		regionnetworkendpointgroup.Setup,
		regionperinstanceconfig.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regionurlmap.Setup,
		reservationcompute.Setup,
		resourcepolicy.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routernat.Setup,
		routerpeer.Setup,
		securitypolicy.Setup,
		serviceattachment.Setup,
		sharedvpchostproject.Setup,
		sharedvpcserviceproject.Setup,
		snapshot.Setup,
		sslcertificate.Setup,
		sslpolicy.Setup,
		subnetwork.Setup,
		subnetworkiambinding.Setup,
		subnetworkiammember.Setup,
		subnetworkiampolicy.Setup,
		targetgrpcproxy.Setup,
		targethttpproxy.Setup,
		targethttpsproxy.Setup,
		targetinstance.Setup,
		targetpool.Setup,
		targetsslproxy.Setup,
		targettcpproxy.Setup,
		urlmap.Setup,
		vpngateway.Setup,
		vpntunnel.Setup,
		cluster.Setup,
		nodepool.Setup,
		registrycontainer.Setup,
		note.Setup,
		occurrence.Setup,
		entry.Setup,
		entrygroup.Setup,
		entrygroupiambinding.Setup,
		entrygroupiammember.Setup,
		entrygroupiampolicy.Setup,
		tag.Setup,
		tagtemplate.Setup,
		tagtemplateiambinding.Setup,
		tagtemplateiammember.Setup,
		tagtemplateiampolicy.Setup,
		jobdataflow.Setup,
		deidentifytemplate.Setup,
		inspecttemplate.Setup,
		jobtrigger.Setup,
		storedinfotype.Setup,
		autoscalingpolicy.Setup,
		clusterdataproc.Setup,
		clusteriambinding.Setup,
		clusteriammember.Setup,
		clusteriampolicy.Setup,
		jobdataproc.Setup,
		jobiambinding.Setup,
		jobiammember.Setup,
		jobiampolicy.Setup,
		workflowtemplate.Setup,
		index.Setup,
		deployment.Setup,
		agent.Setup,
		entitytype.Setup,
		fulfillment.Setup,
		intent.Setup,
		agentdialogflowcx.Setup,
		entitytypedialogflowcx.Setup,
		environmentdialogflowcx.Setup,
		flow.Setup,
		intentdialogflowcx.Setup,
		page.Setup,
		version.Setup,
		managedzone.Setup,
		policydns.Setup,
		recordset.Setup,
		serviceendpoints.Setup,
		serviceiambindingendpoints.Setup,
		serviceiammemberendpoints.Setup,
		serviceiampolicyendpoints.Setup,
		contact.Setup,
		triggereventarc.Setup,
		instancefilestore.Setup,
		document.Setup,
		indexfirestore.Setup,
		gameservercluster.Setup,
		gameserverconfig.Setup,
		gameserverdeployment.Setup,
		gameserverdeploymentrollout.Setup,
		realm.Setup,
		membership.Setup,
		consentstore.Setup,
		consentstoreiambinding.Setup,
		consentstoreiammember.Setup,
		consentstoreiampolicy.Setup,
		datasethealthcare.Setup,
		datasetiambindinghealthcare.Setup,
		datasetiammemberhealthcare.Setup,
		datasetiampolicyhealthcare.Setup,
		dicomstore.Setup,
		dicomstoreiambinding.Setup,
		dicomstoreiammember.Setup,
		dicomstoreiampolicy.Setup,
		fhirstore.Setup,
		fhirstoreiambinding.Setup,
		fhirstoreiammember.Setup,
		fhirstoreiampolicy.Setup,
		hl7v2store.Setup,
		hl7v2storeiambinding.Setup,
		hl7v2storeiammember.Setup,
		hl7v2storeiampolicy.Setup,
		appengineserviceiambinding.Setup,
		appengineserviceiammember.Setup,
		appengineserviceiampolicy.Setup,
		appengineversioniambinding.Setup,
		appengineversioniammember.Setup,
		appengineversioniampolicy.Setup,
		brand.Setup,
		client.Setup,
		tunneliambinding.Setup,
		tunneliammember.Setup,
		tunneliampolicy.Setup,
		tunnelinstanceiambinding.Setup,
		tunnelinstanceiammember.Setup,
		tunnelinstanceiampolicy.Setup,
		webbackendserviceiambinding.Setup,
		webbackendserviceiammember.Setup,
		webbackendserviceiampolicy.Setup,
		webiambinding.Setup,
		webiammember.Setup,
		webiampolicy.Setup,
		webtypeappengineiambinding.Setup,
		webtypeappengineiammember.Setup,
		webtypeappengineiampolicy.Setup,
		webtypecomputeiambinding.Setup,
		webtypecomputeiammember.Setup,
		webtypecomputeiampolicy.Setup,
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
		cryptokey.Setup,
		cryptokeyiambinding.Setup,
		cryptokeyiammember.Setup,
		cryptokeyiampolicy.Setup,
		keyring.Setup,
		keyringiambinding.Setup,
		keyringiammember.Setup,
		keyringiampolicy.Setup,
		keyringimportjob.Setup,
		secretciphertext.Setup,
		billingaccountbucketconfig.Setup,
		billingaccountexclusion.Setup,
		billingaccountsink.Setup,
		folderbucketconfig.Setup,
		folderexclusion.Setup,
		foldersink.Setup,
		metric.Setup,
		organizationbucketconfig.Setup,
		organizationexclusion.Setup,
		organizationsink.Setup,
		projectbucketconfig.Setup,
		projectexclusion.Setup,
		projectsink.Setup,
		instancememcache.Setup,
		model.Setup,
		alertpolicy.Setup,
		customservice.Setup,
		dashboard.Setup,
		groupmonitoring.Setup,
		metricdescriptor.Setup,
		notificationchannel.Setup,
		slo.Setup,
		uptimecheckconfig.Setup,
		connectivitytest.Setup,
		edgecachekeyset.Setup,
		edgecacheorigin.Setup,
		edgecacheservice.Setup,
		environmentnotebooks.Setup,
		instancenotebooks.Setup,
		instanceiambindingnotebooks.Setup,
		instanceiammembernotebooks.Setup,
		instanceiampolicynotebooks.Setup,
		location.Setup,
		policyorgpolicy.Setup,
		patchdeployment.Setup,
		sshpublickey.Setup,
		capool.Setup,
		capooliambinding.Setup,
		capooliammember.Setup,
		capooliampolicy.Setup,
		certificate.Setup,
		certificateauthority.Setup,
		certificatetemplate.Setup,
		providerconfig.Setup,
		litereservation.Setup,
		litesubscription.Setup,
		litetopic.Setup,
		schema.Setup,
		subscription.Setup,
		subscriptioniambinding.Setup,
		subscriptioniammember.Setup,
		subscriptioniampolicy.Setup,
		topic.Setup,
		topiciambinding.Setup,
		topiciammember.Setup,
		topiciampolicy.Setup,
		instanceredis.Setup,
		lien.Setup,
		notificationconfig.Setup,
		source.Setup,
		secret.Setup,
		secretiambinding.Setup,
		secretiammember.Setup,
		secretiampolicy.Setup,
		secretversion.Setup,
		connection.Setup,
		repository.Setup,
		repositoryiambinding.Setup,
		repositoryiammember.Setup,
		repositoryiampolicy.Setup,
		database.Setup,
		databaseiambinding.Setup,
		databaseiammember.Setup,
		databaseiampolicy.Setup,
		instancespanner.Setup,
		instanceiambindingspanner.Setup,
		instanceiammemberspanner.Setup,
		instanceiampolicyspanner.Setup,
		databasesql.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
		bucket.Setup,
		bucketaccesscontrol.Setup,
		bucketacl.Setup,
		bucketiambinding.Setup,
		bucketiammember.Setup,
		bucketiampolicy.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
		hmackey.Setup,
		notification.Setup,
		objectaccesscontrol.Setup,
		objectacl.Setup,
		jobstoragetransfer.Setup,
		tagbinding.Setup,
		tagkey.Setup,
		tagkeyiambinding.Setup,
		tagkeyiammember.Setup,
		tagkeyiampolicy.Setup,
		tagvalue.Setup,
		tagvalueiambinding.Setup,
		tagvalueiammember.Setup,
		tagvalueiampolicy.Setup,
		node.Setup,
		datasetvertexai.Setup,
		connector.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
