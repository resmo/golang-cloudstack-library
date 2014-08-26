package cloudstack

const waitInterval = 1

var isAsyncMap = map[string]bool{
	"listNetworkACLs":                      false,
	"reconnectHost":                        true,
	"createCondition":                      true,
	"copyTemplate":                         true,
	"listRouters":                          false,
	"listNiciraNvpDeviceNetworks":          false,
	"addNicToVirtualMachine":               true,
	"extractVolume":                        true,
	"listCiscoAsa1000vResources":           false,
	"addAccountToProject":                  true,
	"listNetworkServiceProviders":          false,
	"deleteCiscoNexusVSM":                  true,
	"deleteEgressFirewallRule":             true,
	"listNetworkOfferings":                 false,
	"addCluster":                           false,
	"listInternalLoadBalancerElements":     false,
	"listCiscoVnmcResources":               false,
	"listHypervisors":                      false,
	"updateConfiguration":                  false,
	"createVpnConnection":                  true,
	"listVolumes":                          false,
	"suspendProject":                       true,
	"deleteLoadBalancer":                   true,
	"authorizeSecurityGroupIngress":        true,
	"listLoadBalancers":                    false,
	"listTrafficTypeImplementors":          false,
	"addNetscalerLoadBalancer":             true,
	"importLdapUsers":                      false,
	"deleteDomain":                         true,
	"listF5LoadBalancers":                  false,
	"addTrafficMonitor":                    false,
	"createPortableIpRange":                true,
	"configureNetscalerLoadBalancer":       true,
	"createTemplate":                       true,
	"listLoadBalancerRuleInstances":        false,
	"migrateVolume":                        true,
	"deleteLBHealthCheckPolicy":            true,
	"updatePhysicalNetwork":                true,
	"deleteStaticRoute":                    true,
	"deletePaloAltoFirewall":               true,
	"registerSSHKeyPair":                   false,
	"listTrafficMonitors":                  false,
	"updateCloudToUseObjectStore":          false,
	"listAutoScaleVmGroups":                false,
	"getVMPassword":                        false,
	"listInstanceGroups":                   false,
	"addSecondaryStorage":                  false,
	"createNetwork":                        false,
	"listProjects":                         false,
	"enableAccount":                        false,
	"destroySystemVm":                      true,
	"listPublicIpAddresses":                false,
	"listGuestOsMapping":                   false,
	"updateRemoteAccessVpn":                true,
	"enableStorageMaintenance":             true,
	"removeFromGlobalLoadBalancerRule":     true,
	"updateLoadBalancer":                   true,
	"listVpnGateways":                      false,
	"stopRouter":                           true,
	"listClusters":                         false,
	"listDedicatedPods":                    false,
	"attachVolume":                         true,
	"updateVPCOffering":                    true,
	"resetSSHKeyForVirtualMachine":         true,
	"addCiscoAsa1000vResource":             false,
	"listAutoScalePolicies":                false,
	"cleanVMReservations":                  true,
	"createAffinityGroup":                  true,
	"addExternalLoadBalancer":              false,
	"listDeploymentPlanners":               false,
	"listAlerts":                           false,
	"deleteExternalLoadBalancer":           false,
	"deleteTags":                           true,
	"deleteAccountFromProject":             true,
	"listCiscoNexusVSMs":                   false,
	"addBaremetalPxePingServer":            true,
	"listPrivateGateways":                  false,
	"deletePortableIpRange":                true,
	"updateRegion":                         false,
	"updateVolume":                         true,
	"listUcsManagers":                      false,
	"listNetworks":                         false,
	"uploadCustomCertificate":              true,
	"listImageStores":                      false,
	"listCapacity":                         false,
	"createAutoScaleVmProfile":             true,
	"createSecurityGroup":                  false,
	"releaseDedicatedGuestVlanRange":       true,
	"createSSHKeyPair":                     false,
	"cancelHostMaintenance":                true,
	"updateServiceOffering":                false,
	"listLdapUsers":                        false,
	"releaseDedicatedHost":                 true,
	"updateStoragePool":                    false,
	"deleteStorageNetworkIpRange":          true,
	"listS3s":                              false,
	"listVirtualRouterElements":            false,
	"createInternalLoadBalancerElement":    true,
	"deleteFirewallRule":                   true,
	"addBaremetalHost":                     false,
	"deleteNiciraNvpDevice":                true,
	"listApis":                             false,
	"deleteProject":                        true,
	"removeIpFromNic":                      true,
	"updateHostPassword":                   false,
	"deleteIpForwardingRule":               true,
	"getVirtualMachineUserData":            false,
	"addSwift":                             false,
	"createGlobalLoadBalancerRule":         true,
	"resizeVolume":                         true,
	"listSSHKeyPairs":                      false,
	"createStaticRoute":                    true,
	"listNiciraNvpDevices":                 false,
	"deleteGlobalLoadBalancerRule":         true,
	"activateProject":                      true,
	"releaseDedicatedZone":                 true,
	"createVMSnapshot":                     true,
	"configureF5LoadBalancer":              true,
	"enableStaticNat":                      false,
	"createIpForwardingRule":               true,
	"updateIpAddress":                      true,
	"listStorageProviders":                 false,
	"listRegions":                          false,
	"updateDiskOffering":                   false,
	"listNetworkACLLists":                  false,
	"listUcsProfiles":                      false,
	"addPaloAltoFirewall":                  true,
	"recoverVirtualMachine":                false,
	"listCapabilities":                     false,
	"releaseDedicatedCluster":              true,
	"updateVPC":                            true,
	"associateUcsProfileToBlade":           true,
	"startInternalLoadBalancerVM":          true,
	"listProjectAccounts":                  false,
	"updateAutoScaleVmProfile":             true,
	"updatePortForwardingRule":             true,
	"listDedicatedHosts":                   false,
	"listPortForwardingRules":              false,
	"listTemplatePermissions":              false,
	"createStorageNetworkIpRange":          true,
	"uploadSslCert":                        false,
	"queryAsyncJobResult":                  false,
	"createLoadBalancer":                   true,
	"removeVmwareDc":                       false,
	"cancelStorageMaintenance":             true,
	"deployVirtualMachine":                 true,
	"deletePod":                            false,
	"revokeSecurityGroupEgress":            true,
	"deleteCondition":                      true,
	"createNetworkACLList":                 true,
	"createPortForwardingRule":             true,
	"deleteSecondaryStagingStore":          false,
	"listNetscalerLoadBalancers":           false,
	"createVPCOffering":                    true,
	"createEgressFirewallRule":             true,
	"destroyRouter":                        true,
	"deleteAlerts":                         false,
	"updateLBStickinessPolicy":             true,
	"listUsageRecords":                     false,
	"assignToGlobalLoadBalancerRule":       true,
	"listOvsElements":                      false,
	"updateTrafficType":                    true,
	"listPods":                             false,
	"listLdapConfigurations":               false,
	"listOsCategories":                     false,
	"enableUser":                           false,
	"addSrxFirewall":                       true,
	"addNiciraNvpDevice":                   true,
	"addIpToNic":                           true,
	"createInstanceGroup":                  false,
	"updateNetworkACLItem":                 true,
	"deleteNetworkACLList":                 true,
	"updateAccount":                        false,
	"deleteNetworkDevice":                  false,
	"listSnapshotPolicies":                 false,
	"configureInternalLoadBalancerElement": true,
	"releaseHostReservation":               true,
	"createDomain":                         false,
	"deleteSSHKeyPair":                     false,
	"deleteSnapshotPolicies":               false,
	"listLBHealthCheckPolicies":            false,
	"listEvents":                           false,
	"assignCertToLoadBalancer":             true,
	"addHost":                              false,
	"listDedicatedClusters":                false,
	"deleteVpnGateway":                     true,
	"expungeVirtualMachine":                true,
	"addNetworkDevice":                     false,
	"createAutoScaleVmGroup":               true,
	"deleteNetworkServiceProvider":         true,
	"rebootRouter":                         true,
	"createLBHealthCheckPolicy":            true,
	"archiveEvents":                        false,
	"listConfigurations":                   false,
	"updateHost":                           false,
	"listProjectInvitations":               false,
	"deleteIso":                            true,
	"removeGuestOsMapping":                 true,
	"createVpnCustomerGateway":             true,
	"listTrafficTypes":                     false,
	"updateResourceLimit":                  false,
	"lockAccount":                          false,
	"deleteTrafficMonitor":                 false,
	"createUser":                           false,
	"deleteAutoScalePolicy":                true,
	"deleteSrxFirewall":                    true,
	"updateNetworkACLList":                 true,
	"listNetworkIsolationMethods":          false,
	"listDiskOfferings":                    false,
	"detachVolume":                         true,
	"deleteUser":                           false,
	"deleteNetworkACL":                     true,
	"listSnapshots":                        false,
	"deleteVPC":                            true,
	"deleteSecurityGroup":                  false,
	"addS3":                                false,
	"listCounters":                         false,
	"deleteSslCert":                        false,
	"updateHypervisorCapabilities":         false,
	"createPhysicalNetwork":                true,
	"updateLoadBalancerRule":               true,
	"deleteTemplate":                       true,
	"listHypervisorCapabilities":           false,
	"listTags":                             false,
	"deleteVpnCustomerGateway":             true,
	"deleteCiscoVnmcResource":              false,
	"listIsoPermissions":                   false,
	"createVirtualRouterElement":           true,
	"updateAutoScalePolicy":                true,
	"releaseDedicatedPod":                  true,
	"addBigSwitchVnsDevice":                true,
	"createVpnGateway":                     true,
	"dedicateZone":                         true,
	"dedicateCluster":                      true,
	"deleteCounter":                        true,
	"addExternalFirewall":                  false,
	"updateStorageNetworkIpRange":          true,
	"listAsyncJobs":                        false,
	"listUsageTypes":                       false,
	"deleteBigSwitchVnsDevice":             true,
	"listSecondaryStagingStores":           false,
	"listF5LoadBalancerNetworks":           false,
	"createVlanIpRange":                    false,
	"addImageStore":                        false,
	"ldapCreateAccount":                    false,
	"findHostsForMigration":                false,
	"dedicatePod":                          true,
	"addF5LoadBalancer":                    true,
	"deleteAutoScaleVmGroup":               true,
	"listRemoteAccessVpns":                 false,
	"updateGlobalLoadBalancerRule":         true,
	"listAffinityGroupTypes":               false,
	"registerTemplate":                     false,
	"createServiceInstance":                true,
	"deleteNetworkOffering":                false,
	"authorizeSecurityGroupEgress":         true,
	"disableAutoScaleVmGroup":              true,
	"listVpnCustomerGateways":              false,
	"createAccount":                        false,
	"prepareHostForMaintenance":            true,
	"addOpenDaylightController":            true,
	"updateEgressFirewallRule":             true,
	"deletePrivateGateway":                 true,
	"updateGuestOs":                        true,
	"stopInternalLoadBalancerVM":           true,
	"createSecondaryStagingStore":          false,
	"deleteTrafficType":                    true,
	"generateUsageRecords":                 false,
	"deleteLoadBalancerRule":               true,
	"getApiLimit":                          false,
	"attachIso":                            true,
	"deletePortForwardingRule":             true,
	"listDomains":                          false,
	"getUser":                              false,
	"listLoadBalancerRules":                false,
	"listPaloAltoFirewalls":                false,
	"deleteEvents":                         false,
	"removeCertFromLoadBalancer":           true,
	"configureSrxFirewall":                 true,
	"deleteZone":                           false,
	"updateProjectInvitation":              true,
	"deleteVolume":                         false,
	"createTags":                           true,
	"enableAutoScaleVmGroup":               true,
	"listOpenDaylightControllers":          false,
	"listResourceDetails":                  false,
	"deleteLdapConfiguration":              false,
	"updateCluster":                        false,
	"listSrxFirewalls":                     false,
	"removeVpnUser":                        true,
	"scaleVirtualMachine":                  true,
	"listPaloAltoFirewallNetworks":         false,
	"releasePublicIpRange":                 false,
	"listTemplates":                        false,
	"listDedicatedGuestVlanRanges":         false,
	"updateVpnCustomerGateway":             true,
	"updateIsoPermissions":                 false,
	"stopSystemVm":                         true,
	"restartNetwork":                       true,
	"prepareTemplate":                      false,
	"removeGuestOs":                        true,
	"changeServiceForRouter":               false,
	"rebootVirtualMachine":                 true,
	"listGlobalLoadBalancerRules":          false,
	"listLBStickinessPolicies":             false,
	"updateZone":                           false,
	"enableCiscoNexusVSM":                  true,
	"listSecurityGroups":                   false,
	"dedicateGuestVlanRange":               false,
	"listFirewallRules":                    false,
	"updateVMAffinityGroup":                true,
	"configurePaloAltoFirewall":            true,
	"deleteVpnConnection":                  true,
	"listAffinityGroups":                   false,
	"updateUser":                           false,
	"updateVpnConnection":                  true,
	"deleteDiskOffering":                   false,
	"updateVpnGateway":                     true,
	"startSystemVm":                        true,
	"deleteF5LoadBalancer":                 true,
	"updateProject":                        true,
	"archiveAlerts":                        false,
	"createZone":                           false,
	"listBaremetalPxeServers":              false,
	"deleteNetwork":                        true,
	"listStaticRoutes":                     false,
	"changeServiceForSystemVm":             false,
	"addCiscoVnmcResource":                 false,
	"listNics":                             false,
	"deleteNetscalerLoadBalancer":          true,
	"listSrxFirewallNetworks":              false,
	"createStoragePool":                    false,
	"removeRegion":                         false,
	"listBaremetalDhcp":                    false,
	"addTrafficType":                       true,
	"listSwifts":                           false,
	"updateTemplate":                       false,
	"disableUser":                          true,
	"configureVirtualRouterElement":        true,
	"deleteProjectInvitation":              true,
	"createSnapshotPolicy":                 false,
	"updateInstanceGroup":                  false,
	"migrateSystemVm":                      true,
	"createServiceOffering":                false,
	"removeNicFromVirtualMachine":          true,
	"deleteCiscoAsa1000vResource":          false,
	"revokeSecurityGroupIngress":           true,
	"updateDefaultNicForVirtualMachine":    true,
	"listExternalFirewalls":                false,
	"disableStaticNat":                     true,
	"createNetworkACL":                     true,
	"createPod":                            false,
	"createVPC":                            true,
	"listOsTypes":                          false,
	"addBaremetalPxeKickStartServer":       true,
	"registerIso":                          false,
	"addResourceDetail":                    true,
	"listSslCerts":                         false,
	"disassociateIpAddress":                true,
	"listZones":                            false,
	"listExternalLoadBalancers":            false,
	"resetPasswordForVirtualMachine":       true,
	"createVolume":                         true,
	"assignToLoadBalancerRule":             true,
	"listUcsBlades":                        false,
	"startRouter":                          true,
	"updateGuestOsMapping":                 true,
	"extractIso":                           true,
	"removeResourceDetail":                 true,
	"changeServiceForVirtualMachine":       false,
	"deleteRemoteAccessVpn":                true,
	"addVmwareDc":                          false,
	"deleteImageStore":                     false,
	"deleteVlanIpRange":                    false,
	"listStoragePools":                     false,
	"resetVpnConnection":                   true,
	"createRemoteAccessVpn":                true,
	"startVirtualMachine":                  true,
	"extractTemplate":                      true,
	"listSystemVms":                        false,
	"detachIso":                            true,
	"deleteServiceOffering":                false,
	"deleteAccount":                        true,
	"associateIpAddress":                   true,
	"listNetworkDevice":                    false,
	"disableAccount":                       true,
	"migrateVirtualMachine":                true,
	"listVMSnapshot":                       false,
	"updateDomain":                         false,
	"listDedicatedZones":                   false,
	"removeFromLoadBalancerRule":           true,
	"resetApiLimit":                        false,
	"registerUserKeys":                     false,
	"addVpnUser":                           true,
	"listVPCs":                             false,
	"updateLBHealthCheckPolicy":            true,
	"assignVirtualMachine":                 false,
	"updateFirewallRule":                   true,
	"listConditions":                       false,
	"updateVirtualMachine":                 false,
	"createPrivateGateway":                 true,
	"deleteLBStickinessPolicy":             true,
	"listResourceLimits":                   false,
	"listServiceOfferings":                 false,
	"disableCiscoNexusVSM":                 true,
	"addUcsManager":                        false,
	"deleteVMSnapshot":                     true,
	"deleteAutoScaleVmProfile":             true,
	"deleteStoragePool":                    false,
	"deleteSnapshot":                       true,
	"createProject":                        true,
	"createLoadBalancerRule":               true,
	"createAutoScalePolicy":                true,
	"restoreVirtualMachine":                true,
	"listEventTypes":                       false,
	"createNetworkOffering":                false,
	"listBigSwitchVnsDevices":              false,
	"deleteAffinityGroup":                  true,
	"copyIso":                              true,
	"dedicatePublicIpRange":                false,
	"addGuestOsMapping":                    true,
	"listDomainChildren":                   false,
	"uploadVolume":                         true,
	"listAutoScaleVmProfiles":              false,
	"createLBStickinessPolicy":             true,
	"migrateVirtualMachineWithVolume":      true,
	"stopVirtualMachine":                   true,
	"createCounter":                        true,
	"listAccounts":                         false,
	"createSnapshot":                       true,
	"updateIso":                            false,
	"listPortableIpRanges":                 false,
	"configureOvsElement":                  true,
	"listIpForwardingRules":                false,
	"destroyVirtualMachine":                true,
	"updateNetwork":                        true,
	"dedicateHost":                         true,
	"addRegion":                            false,
	"createDiskOffering":                   false,
	"getCloudIdentifier":                   false,
	"listNetscalerLoadBalancerNetworks":    false,
	"deleteExternalFirewall":               false,
	"listInternalLoadBalancerVMs":          false,
	"createFirewallRule":                   true,
	"updateResourceCount":                  false,
	"addNetworkServiceProvider":            true,
	"rebootSystemVm":                       true,
	"revertToVMSnapshot":                   true,
	"lockUser":                             false,
	"markDefaultZoneForAccount":            true,
	"addLdapConfiguration":                 false,
	"listVirtualMachines":                  false,
	"restartVPC":                           true,
	"replaceNetworkACLList":                true,
	"generateAlert":                        true,
	"scaleSystemVm":                        true,
	"listEgressFirewallRules":              false,
	"updateAutoScaleVmGroup":               true,
	"listHosts":                            false,
	"updateTemplatePermissions":            false,
	"listVlanIpRanges":                     false,
	"listPhysicalNetworks":                 false,
	"listStorageNetworkIpRange":            false,
	"listVpnConnections":                   false,
	"deleteHost":                           false,
	"listVPCOfferings":                     false,
	"updateNetworkOffering":                false,
	"deletePhysicalNetwork":                true,
	"deleteInstanceGroup":                  false,
	"addBaremetalDhcp":                     true,
	"deleteVPCOffering":                    true,
	"listVpnUsers":                         false,
	"updatePod":                            false,
	"listVmwareDcs":                        false,
	"deleteOpenDaylightController":         true,
	"deleteCluster":                        false,
	"updateNetworkServiceProvider":         true,
	"listUsers":                            false,
	"findStoragePoolsForMigration":         false,
	"upgradeRouterTemplate":                false,
	"listSupportedNetworkServices":         false,
	"addStratosphereSsp":                   false,
	"addGuestOs":                           true,
	"listIsos":                             false,
	"revertSnapshot":                       true,
}