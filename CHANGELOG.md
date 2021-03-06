## 1.2.4 (March 11, 2020)

ENHANCEMENTS:

* Provider: Added new parameter `zone` to support power virtual resources and data sources to work in multi-zone environment ([#1141](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1141))
* resource/ibm_pi_volume: Updated the list of volume types for power virtual volume ([#1149](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1149))
* resource/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* data/ibm_container_vpc_cluster : Added support for `ingress_hostname` and `ingress_secret` attributes ([#1167](https://github.com/IBM-Cloud/terraform-provider-ibm/issue/1167))
* resource/ibm_is_floating_ip : Handle the case when floating IP is deleted manually ([#1160](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1160))

BUG FIXES:

* resources : Handle the case where the resource might be already deleted (manually) for ibm_iam_access_policies, ibm_iam_authorization_policies, ibm_iam_service_policies ([#1162](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1162))
* resource/ibm_is_inetwork_acl: Fix the order of creation of network acl ([#1123](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1123))
* resource/ibm_container_vpc_cluster: Added new attribute `wait_till` to control the cluster creation. Now user can control the cluster creation until master is ready / any one worker node is ready / ingress_hostname is  
  assigned.  ([#1143](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1143))
* resource/ibm_pi_instance: Fix the timeout configuration for create ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))
* doc/ibm_cis_ip_addresses : Fix the description of data source ([#1178](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1178))

## 1.2.3 (March 03, 2020)

BUG FIXES:

* data/ibm_container_cluster_config : Fix the error to download the cluster config for VPC clusters ([#1150](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1150))

## 1.2.2 (February 26, 2020)

ENHANCEMENTS:

* resource/ibm_is_vpc: Improved error message for VPC creation ([#1106](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1106))
* resource/ibm_is_ssh_key: Improved error message for VPC SSH Key creation ([#1105](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1105))
* resource/ibm_container_cluster : Added gateway feature support for IKS clusters. This feature helps to create a cluster with a gateway worker pool of two gateway worker nodes that are connected to public and private VLANs to provide limited public access, and a compute worker pool of compute worker nodes that are connected to the private VLAN only. 
([#1125](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1125))
* data/ibm_conatiner_cluster_config : Extended the data source to provide additional attribute like admin_key, admin_certificate, ca_certificate, host and token. This attributes helps to connect to other providers like Kubernetes and Helm without loading cluster config file. ([#895](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/895))

BUG FIXES:

* doc/ibm_certificate_manager_order: Changed the type of rotate_key from string to bool ([#1110](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1110))
* resource/ibm_is_instance: Fix for updating security group for primary network interface for vpc instance. Now users can add or delete security groups([#1078](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1078))
* doc/ibm_resource_key : Provided an example in the docs as a workaround to create credentials using serviceID parameter ([#1121](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1121))
* resource/ibm_is_network_acl : Fix for crash during the update of rules. Fix for the order of rules creation. Now users can add or delete rules for network_acl ([#1117](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1117))
* resource/ibm_is_public_gateway : Added support for resource group and tags parameters ([#1102](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1102))
* resource/ibm_is_floating_ip : Added support for tags parameters ([#1131](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1131))
* resource/ibm_database : Parameters remote_leader_id, key_protect_instance and key_protect_key can’t be updated after creation. ([#1111](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1111))
* example/ibm-key-protect : Updated example to create an authorisation policy between COS and Key Protect instance([#1133](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1133))
* resource/ibm_resource_group: Removed suppression of error during deletion ([#1108](https://github.com/IBM-Cloud/terraform-provider-ibm/pull/1108))
* resource/ibm_iam_user_invite : Fix for inviting user from IBM Cloud lite account. ([#1114](https://github.com/IBM-Cloud/terraform-provider-ibm/issues/1114))


