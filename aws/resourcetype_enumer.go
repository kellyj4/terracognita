// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "aws_instanceaws_amiaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_cloudwatch_dashboardaws_cloudwatch_log_groupaws_cloudwatch_log_streamaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_connectionaws_dx_gatewayaws_dx_private_virtual_interfaceaws_dx_public_virtual_interfaceaws_dx_lagaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_snapshotaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_managed_prefix_listaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_network_aclaws_networkfirewall_firewallaws_networkfirewall_firewall_policyaws_networkfirewall_rule_groupaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_dhcp_optionsaws_vpc_endpointaws_vpc_ipamaws_vpc_ipam_poolaws_vpc_ipam_scopeaws_vpc_peering_connectionaws_vpn_gateway"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 26, 42, 70, 91, 111, 142, 168, 192, 216, 237, 257, 278, 300, 324, 348, 375, 412, 437, 464, 488, 512, 537, 552, 567, 589, 608, 639, 667, 684, 698, 730, 761, 771, 796, 814, 830, 844, 859, 874, 897, 924, 947, 985, 1020, 1060, 1102, 1153, 1198, 1227, 1274, 1321, 1368, 1387, 1394, 1409, 1432, 1465, 1498, 1522, 1553, 1560, 1575, 1601, 1626, 1648, 1666, 1687, 1718, 1731, 1755, 1775, 1806, 1830, 1861, 1875, 1887, 1906, 1936, 1957, 1983, 1995, 2024, 2043, 2073, 2093, 2113, 2125, 2143, 2162, 2186, 2205, 2211, 2242, 2257, 2284, 2304, 2323, 2353, 2375, 2400, 2413, 2428, 2443, 2471, 2506, 2536, 2555, 2570, 2592, 2612, 2638, 2662, 2683, 2701, 2730, 2767, 2783, 2811, 2826, 2839, 2857, 2888, 2913, 2932, 2955, 2979, 3014, 3036, 3056, 3080, 3096, 3109, 3135, 3145, 3166, 3173, 3193, 3209, 3221, 3238, 3256, 3282, 3297}

const _ResourceTypeLowerName = "aws_instanceaws_amiaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_cloudwatch_dashboardaws_cloudwatch_log_groupaws_cloudwatch_log_streamaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_connectionaws_dx_gatewayaws_dx_private_virtual_interfaceaws_dx_public_virtual_interfaceaws_dx_lagaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_snapshotaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_managed_prefix_listaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_network_aclaws_networkfirewall_firewallaws_networkfirewall_firewall_policyaws_networkfirewall_rule_groupaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_dhcp_optionsaws_vpc_endpointaws_vpc_ipamaws_vpc_ipam_poolaws_vpc_ipam_scopeaws_vpc_peering_connectionaws_vpn_gateway"

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i+1)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[Instance-(1)]
	_ = x[AMI-(2)]
	_ = x[ALB-(3)]
	_ = x[ALBListener-(4)]
	_ = x[ALBListenerCertificate-(5)]
	_ = x[ALBListenerRule-(6)]
	_ = x[ALBTargetGroup-(7)]
	_ = x[ALBTargetGroupAttachment-(8)]
	_ = x[APIGatewayDeployment-(9)]
	_ = x[APIGatewayResource-(10)]
	_ = x[APIGatewayRestAPI-(11)]
	_ = x[APIGatewayStage-(12)]
	_ = x[AthenaWorkgroup-(13)]
	_ = x[AutoscalingGroup-(14)]
	_ = x[AutoscalingPolicy-(15)]
	_ = x[AutoscalingSchedule-(16)]
	_ = x[BatchJobDefinition-(17)]
	_ = x[CloudfrontDistribution-(18)]
	_ = x[CloudfrontOriginAccessIdentity-(19)]
	_ = x[CloudfrontPublicKey-(20)]
	_ = x[CloudwatchMetricAlarm-(21)]
	_ = x[CloudwatchDashboard-(22)]
	_ = x[CloudwatchLogGroup-(23)]
	_ = x[CloudwatchLogStream-(24)]
	_ = x[DaxCluster-(25)]
	_ = x[DBInstance-(26)]
	_ = x[DBParameterGroup-(27)]
	_ = x[DBSubnetGroup-(28)]
	_ = x[DirectoryServiceDirectory-(29)]
	_ = x[DmsReplicationInstance-(30)]
	_ = x[DXConnection-(31)]
	_ = x[DXGateway-(32)]
	_ = x[DXPrivateVirtualInterface-(33)]
	_ = x[DXPublicVirtualInterface-(34)]
	_ = x[DXLag-(35)]
	_ = x[DynamodbGlobalTable-(36)]
	_ = x[DynamodbTable-(37)]
	_ = x[EBSSnapshot-(38)]
	_ = x[EBSVolume-(39)]
	_ = x[ECSCluster-(40)]
	_ = x[ECSService-(41)]
	_ = x[ECSTaskDefinition-(42)]
	_ = x[EC2ManagedPrefixList-(43)]
	_ = x[EC2TransitGateway-(44)]
	_ = x[EC2TransitGatewayVPCAttachment-(45)]
	_ = x[EC2TransitGatewayRouteTable-(46)]
	_ = x[EC2TransitGatewayMulticastDomain-(47)]
	_ = x[EC2TransitGatewayPeeringAttachment-(48)]
	_ = x[EC2TransitGatewayPeeringAttachmentAccepter-(49)]
	_ = x[EC2TransitGatewayPrefixListReference-(50)]
	_ = x[EC2TransitGatewayRoute-(51)]
	_ = x[EC2TransitGatewayRouteTableAssociation-(52)]
	_ = x[EC2TransitGatewayRouteTablePropagation-(53)]
	_ = x[EC2TransitGatewayVPCAttachmentAccepter-(54)]
	_ = x[EFSFileSystem-(55)]
	_ = x[EIP-(56)]
	_ = x[EKSCluster-(57)]
	_ = x[ElasticacheCluster-(58)]
	_ = x[ElasticacheReplicationGroup-(59)]
	_ = x[ElasticBeanstalkApplication-(60)]
	_ = x[ElasticsearchDomain-(61)]
	_ = x[ElasticsearchDomainPolicy-(62)]
	_ = x[ELB-(63)]
	_ = x[EMRCluster-(64)]
	_ = x[FsxLustreFileSystem-(65)]
	_ = x[GlueCatalogDatabase-(66)]
	_ = x[GlueCatalogTable-(67)]
	_ = x[IAMAccessKey-(68)]
	_ = x[IAMAccountAlias-(69)]
	_ = x[IAMAccountPasswordPolicy-(70)]
	_ = x[IAMGroup-(71)]
	_ = x[IAMGroupMembership-(72)]
	_ = x[IAMGroupPolicy-(73)]
	_ = x[IAMGroupPolicyAttachment-(74)]
	_ = x[IAMInstanceProfile-(75)]
	_ = x[IAMOpenidConnectProvider-(76)]
	_ = x[IAMPolicy-(77)]
	_ = x[IAMRole-(78)]
	_ = x[IAMRolePolicy-(79)]
	_ = x[IAMRolePolicyAttachment-(80)]
	_ = x[IAMSAMLProvider-(81)]
	_ = x[IAMServerCertificate-(82)]
	_ = x[IAMUser-(83)]
	_ = x[IAMUserGroupMembership-(84)]
	_ = x[IAMUserPolicy-(85)]
	_ = x[IAMUserPolicyAttachment-(86)]
	_ = x[IAMUserSSHKey-(87)]
	_ = x[InternetGateway-(88)]
	_ = x[KeyPair-(89)]
	_ = x[KinesisStream-(90)]
	_ = x[LambdaFunction-(91)]
	_ = x[LaunchConfiguration-(92)]
	_ = x[LaunchTemplate-(93)]
	_ = x[LB-(94)]
	_ = x[LBCookieStickinessPolicy-(95)]
	_ = x[LBListener-(96)]
	_ = x[LBListenerCertificate-(97)]
	_ = x[LBListenerRule-(98)]
	_ = x[LBTargetGroup-(99)]
	_ = x[LBTargetGroupAttachment-(100)]
	_ = x[LightsailInstance-(101)]
	_ = x[MediaStoreContainer-(102)]
	_ = x[MQBroker-(103)]
	_ = x[NatGateway-(104)]
	_ = x[NetworkAcl-(105)]
	_ = x[NetworkfirewallFirewall-(106)]
	_ = x[NetworkfirewallFirewallPolicy-(107)]
	_ = x[NetworkfirewallRuleGroup-(108)]
	_ = x[NeptuneCluster-(109)]
	_ = x[RDSCluster-(110)]
	_ = x[RDSGlobalCluster-(111)]
	_ = x[RedshiftCluster-(112)]
	_ = x[Route53DelegationSet-(113)]
	_ = x[Route53HealthCheck-(114)]
	_ = x[Route53QueryLog-(115)]
	_ = x[Route53Record-(116)]
	_ = x[Route53ResolverEndpoint-(117)]
	_ = x[Route53ResolverRuleAssociation-(118)]
	_ = x[Route53Zone-(119)]
	_ = x[Route53ZoneAssociation-(120)]
	_ = x[RouteTable-(121)]
	_ = x[S3Bucket-(122)]
	_ = x[SecurityGroup-(123)]
	_ = x[SESActiveReceiptRuleSet-(124)]
	_ = x[SESConfigurationSet-(125)]
	_ = x[SESDomainDKIM-(126)]
	_ = x[SESDomainIdentity-(127)]
	_ = x[SESDomainMailFrom-(128)]
	_ = x[SESIdentityNotificationTopic-(129)]
	_ = x[SESReceiptFilter-(130)]
	_ = x[SESReceiptRule-(131)]
	_ = x[SESReceiptRuleSet-(132)]
	_ = x[SESTemplate-(133)]
	_ = x[SQSQueue-(134)]
	_ = x[StoragegatewayGateway-(135)]
	_ = x[Subnet-(136)]
	_ = x[VolumeAttachment-(137)]
	_ = x[VPC-(138)]
	_ = x[VPCDhcpOptions-(139)]
	_ = x[VPCEndpoint-(140)]
	_ = x[VPCIpam-(141)]
	_ = x[VPCIpamPool-(142)]
	_ = x[VPCIpamScope-(143)]
	_ = x[VPCPeeringConnection-(144)]
	_ = x[VPNGateway-(145)]
}

var _ResourceTypeValues = []ResourceType{Instance, AMI, ALB, ALBListener, ALBListenerCertificate, ALBListenerRule, ALBTargetGroup, ALBTargetGroupAttachment, APIGatewayDeployment, APIGatewayResource, APIGatewayRestAPI, APIGatewayStage, AthenaWorkgroup, AutoscalingGroup, AutoscalingPolicy, AutoscalingSchedule, BatchJobDefinition, CloudfrontDistribution, CloudfrontOriginAccessIdentity, CloudfrontPublicKey, CloudwatchMetricAlarm, CloudwatchDashboard, CloudwatchLogGroup, CloudwatchLogStream, DaxCluster, DBInstance, DBParameterGroup, DBSubnetGroup, DirectoryServiceDirectory, DmsReplicationInstance, DXConnection, DXGateway, DXPrivateVirtualInterface, DXPublicVirtualInterface, DXLag, DynamodbGlobalTable, DynamodbTable, EBSSnapshot, EBSVolume, ECSCluster, ECSService, ECSTaskDefinition, EC2ManagedPrefixList, EC2TransitGateway, EC2TransitGatewayVPCAttachment, EC2TransitGatewayRouteTable, EC2TransitGatewayMulticastDomain, EC2TransitGatewayPeeringAttachment, EC2TransitGatewayPeeringAttachmentAccepter, EC2TransitGatewayPrefixListReference, EC2TransitGatewayRoute, EC2TransitGatewayRouteTableAssociation, EC2TransitGatewayRouteTablePropagation, EC2TransitGatewayVPCAttachmentAccepter, EFSFileSystem, EIP, EKSCluster, ElasticacheCluster, ElasticacheReplicationGroup, ElasticBeanstalkApplication, ElasticsearchDomain, ElasticsearchDomainPolicy, ELB, EMRCluster, FsxLustreFileSystem, GlueCatalogDatabase, GlueCatalogTable, IAMAccessKey, IAMAccountAlias, IAMAccountPasswordPolicy, IAMGroup, IAMGroupMembership, IAMGroupPolicy, IAMGroupPolicyAttachment, IAMInstanceProfile, IAMOpenidConnectProvider, IAMPolicy, IAMRole, IAMRolePolicy, IAMRolePolicyAttachment, IAMSAMLProvider, IAMServerCertificate, IAMUser, IAMUserGroupMembership, IAMUserPolicy, IAMUserPolicyAttachment, IAMUserSSHKey, InternetGateway, KeyPair, KinesisStream, LambdaFunction, LaunchConfiguration, LaunchTemplate, LB, LBCookieStickinessPolicy, LBListener, LBListenerCertificate, LBListenerRule, LBTargetGroup, LBTargetGroupAttachment, LightsailInstance, MediaStoreContainer, MQBroker, NatGateway, NetworkAcl, NetworkfirewallFirewall, NetworkfirewallFirewallPolicy, NetworkfirewallRuleGroup, NeptuneCluster, RDSCluster, RDSGlobalCluster, RedshiftCluster, Route53DelegationSet, Route53HealthCheck, Route53QueryLog, Route53Record, Route53ResolverEndpoint, Route53ResolverRuleAssociation, Route53Zone, Route53ZoneAssociation, RouteTable, S3Bucket, SecurityGroup, SESActiveReceiptRuleSet, SESConfigurationSet, SESDomainDKIM, SESDomainIdentity, SESDomainMailFrom, SESIdentityNotificationTopic, SESReceiptFilter, SESReceiptRule, SESReceiptRuleSet, SESTemplate, SQSQueue, StoragegatewayGateway, Subnet, VolumeAttachment, VPC, VPCDhcpOptions, VPCEndpoint, VPCIpam, VPCIpamPool, VPCIpamScope, VPCPeeringConnection, VPNGateway}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:12]:           Instance,
	_ResourceTypeLowerName[0:12]:      Instance,
	_ResourceTypeName[12:19]:          AMI,
	_ResourceTypeLowerName[12:19]:     AMI,
	_ResourceTypeName[19:26]:          ALB,
	_ResourceTypeLowerName[19:26]:     ALB,
	_ResourceTypeName[26:42]:          ALBListener,
	_ResourceTypeLowerName[26:42]:     ALBListener,
	_ResourceTypeName[42:70]:          ALBListenerCertificate,
	_ResourceTypeLowerName[42:70]:     ALBListenerCertificate,
	_ResourceTypeName[70:91]:          ALBListenerRule,
	_ResourceTypeLowerName[70:91]:     ALBListenerRule,
	_ResourceTypeName[91:111]:         ALBTargetGroup,
	_ResourceTypeLowerName[91:111]:    ALBTargetGroup,
	_ResourceTypeName[111:142]:        ALBTargetGroupAttachment,
	_ResourceTypeLowerName[111:142]:   ALBTargetGroupAttachment,
	_ResourceTypeName[142:168]:        APIGatewayDeployment,
	_ResourceTypeLowerName[142:168]:   APIGatewayDeployment,
	_ResourceTypeName[168:192]:        APIGatewayResource,
	_ResourceTypeLowerName[168:192]:   APIGatewayResource,
	_ResourceTypeName[192:216]:        APIGatewayRestAPI,
	_ResourceTypeLowerName[192:216]:   APIGatewayRestAPI,
	_ResourceTypeName[216:237]:        APIGatewayStage,
	_ResourceTypeLowerName[216:237]:   APIGatewayStage,
	_ResourceTypeName[237:257]:        AthenaWorkgroup,
	_ResourceTypeLowerName[237:257]:   AthenaWorkgroup,
	_ResourceTypeName[257:278]:        AutoscalingGroup,
	_ResourceTypeLowerName[257:278]:   AutoscalingGroup,
	_ResourceTypeName[278:300]:        AutoscalingPolicy,
	_ResourceTypeLowerName[278:300]:   AutoscalingPolicy,
	_ResourceTypeName[300:324]:        AutoscalingSchedule,
	_ResourceTypeLowerName[300:324]:   AutoscalingSchedule,
	_ResourceTypeName[324:348]:        BatchJobDefinition,
	_ResourceTypeLowerName[324:348]:   BatchJobDefinition,
	_ResourceTypeName[348:375]:        CloudfrontDistribution,
	_ResourceTypeLowerName[348:375]:   CloudfrontDistribution,
	_ResourceTypeName[375:412]:        CloudfrontOriginAccessIdentity,
	_ResourceTypeLowerName[375:412]:   CloudfrontOriginAccessIdentity,
	_ResourceTypeName[412:437]:        CloudfrontPublicKey,
	_ResourceTypeLowerName[412:437]:   CloudfrontPublicKey,
	_ResourceTypeName[437:464]:        CloudwatchMetricAlarm,
	_ResourceTypeLowerName[437:464]:   CloudwatchMetricAlarm,
	_ResourceTypeName[464:488]:        CloudwatchDashboard,
	_ResourceTypeLowerName[464:488]:   CloudwatchDashboard,
	_ResourceTypeName[488:512]:        CloudwatchLogGroup,
	_ResourceTypeLowerName[488:512]:   CloudwatchLogGroup,
	_ResourceTypeName[512:537]:        CloudwatchLogStream,
	_ResourceTypeLowerName[512:537]:   CloudwatchLogStream,
	_ResourceTypeName[537:552]:        DaxCluster,
	_ResourceTypeLowerName[537:552]:   DaxCluster,
	_ResourceTypeName[552:567]:        DBInstance,
	_ResourceTypeLowerName[552:567]:   DBInstance,
	_ResourceTypeName[567:589]:        DBParameterGroup,
	_ResourceTypeLowerName[567:589]:   DBParameterGroup,
	_ResourceTypeName[589:608]:        DBSubnetGroup,
	_ResourceTypeLowerName[589:608]:   DBSubnetGroup,
	_ResourceTypeName[608:639]:        DirectoryServiceDirectory,
	_ResourceTypeLowerName[608:639]:   DirectoryServiceDirectory,
	_ResourceTypeName[639:667]:        DmsReplicationInstance,
	_ResourceTypeLowerName[639:667]:   DmsReplicationInstance,
	_ResourceTypeName[667:684]:        DXConnection,
	_ResourceTypeLowerName[667:684]:   DXConnection,
	_ResourceTypeName[684:698]:        DXGateway,
	_ResourceTypeLowerName[684:698]:   DXGateway,
	_ResourceTypeName[698:730]:        DXPrivateVirtualInterface,
	_ResourceTypeLowerName[698:730]:   DXPrivateVirtualInterface,
	_ResourceTypeName[730:761]:        DXPublicVirtualInterface,
	_ResourceTypeLowerName[730:761]:   DXPublicVirtualInterface,
	_ResourceTypeName[761:771]:        DXLag,
	_ResourceTypeLowerName[761:771]:   DXLag,
	_ResourceTypeName[771:796]:        DynamodbGlobalTable,
	_ResourceTypeLowerName[771:796]:   DynamodbGlobalTable,
	_ResourceTypeName[796:814]:        DynamodbTable,
	_ResourceTypeLowerName[796:814]:   DynamodbTable,
	_ResourceTypeName[814:830]:        EBSSnapshot,
	_ResourceTypeLowerName[814:830]:   EBSSnapshot,
	_ResourceTypeName[830:844]:        EBSVolume,
	_ResourceTypeLowerName[830:844]:   EBSVolume,
	_ResourceTypeName[844:859]:        ECSCluster,
	_ResourceTypeLowerName[844:859]:   ECSCluster,
	_ResourceTypeName[859:874]:        ECSService,
	_ResourceTypeLowerName[859:874]:   ECSService,
	_ResourceTypeName[874:897]:        ECSTaskDefinition,
	_ResourceTypeLowerName[874:897]:   ECSTaskDefinition,
	_ResourceTypeName[897:924]:        EC2ManagedPrefixList,
	_ResourceTypeLowerName[897:924]:   EC2ManagedPrefixList,
	_ResourceTypeName[924:947]:        EC2TransitGateway,
	_ResourceTypeLowerName[924:947]:   EC2TransitGateway,
	_ResourceTypeName[947:985]:        EC2TransitGatewayVPCAttachment,
	_ResourceTypeLowerName[947:985]:   EC2TransitGatewayVPCAttachment,
	_ResourceTypeName[985:1020]:       EC2TransitGatewayRouteTable,
	_ResourceTypeLowerName[985:1020]:  EC2TransitGatewayRouteTable,
	_ResourceTypeName[1020:1060]:      EC2TransitGatewayMulticastDomain,
	_ResourceTypeLowerName[1020:1060]: EC2TransitGatewayMulticastDomain,
	_ResourceTypeName[1060:1102]:      EC2TransitGatewayPeeringAttachment,
	_ResourceTypeLowerName[1060:1102]: EC2TransitGatewayPeeringAttachment,
	_ResourceTypeName[1102:1153]:      EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeLowerName[1102:1153]: EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeName[1153:1198]:      EC2TransitGatewayPrefixListReference,
	_ResourceTypeLowerName[1153:1198]: EC2TransitGatewayPrefixListReference,
	_ResourceTypeName[1198:1227]:      EC2TransitGatewayRoute,
	_ResourceTypeLowerName[1198:1227]: EC2TransitGatewayRoute,
	_ResourceTypeName[1227:1274]:      EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeLowerName[1227:1274]: EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeName[1274:1321]:      EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeLowerName[1274:1321]: EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeName[1321:1368]:      EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeLowerName[1321:1368]: EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeName[1368:1387]:      EFSFileSystem,
	_ResourceTypeLowerName[1368:1387]: EFSFileSystem,
	_ResourceTypeName[1387:1394]:      EIP,
	_ResourceTypeLowerName[1387:1394]: EIP,
	_ResourceTypeName[1394:1409]:      EKSCluster,
	_ResourceTypeLowerName[1394:1409]: EKSCluster,
	_ResourceTypeName[1409:1432]:      ElasticacheCluster,
	_ResourceTypeLowerName[1409:1432]: ElasticacheCluster,
	_ResourceTypeName[1432:1465]:      ElasticacheReplicationGroup,
	_ResourceTypeLowerName[1432:1465]: ElasticacheReplicationGroup,
	_ResourceTypeName[1465:1498]:      ElasticBeanstalkApplication,
	_ResourceTypeLowerName[1465:1498]: ElasticBeanstalkApplication,
	_ResourceTypeName[1498:1522]:      ElasticsearchDomain,
	_ResourceTypeLowerName[1498:1522]: ElasticsearchDomain,
	_ResourceTypeName[1522:1553]:      ElasticsearchDomainPolicy,
	_ResourceTypeLowerName[1522:1553]: ElasticsearchDomainPolicy,
	_ResourceTypeName[1553:1560]:      ELB,
	_ResourceTypeLowerName[1553:1560]: ELB,
	_ResourceTypeName[1560:1575]:      EMRCluster,
	_ResourceTypeLowerName[1560:1575]: EMRCluster,
	_ResourceTypeName[1575:1601]:      FsxLustreFileSystem,
	_ResourceTypeLowerName[1575:1601]: FsxLustreFileSystem,
	_ResourceTypeName[1601:1626]:      GlueCatalogDatabase,
	_ResourceTypeLowerName[1601:1626]: GlueCatalogDatabase,
	_ResourceTypeName[1626:1648]:      GlueCatalogTable,
	_ResourceTypeLowerName[1626:1648]: GlueCatalogTable,
	_ResourceTypeName[1648:1666]:      IAMAccessKey,
	_ResourceTypeLowerName[1648:1666]: IAMAccessKey,
	_ResourceTypeName[1666:1687]:      IAMAccountAlias,
	_ResourceTypeLowerName[1666:1687]: IAMAccountAlias,
	_ResourceTypeName[1687:1718]:      IAMAccountPasswordPolicy,
	_ResourceTypeLowerName[1687:1718]: IAMAccountPasswordPolicy,
	_ResourceTypeName[1718:1731]:      IAMGroup,
	_ResourceTypeLowerName[1718:1731]: IAMGroup,
	_ResourceTypeName[1731:1755]:      IAMGroupMembership,
	_ResourceTypeLowerName[1731:1755]: IAMGroupMembership,
	_ResourceTypeName[1755:1775]:      IAMGroupPolicy,
	_ResourceTypeLowerName[1755:1775]: IAMGroupPolicy,
	_ResourceTypeName[1775:1806]:      IAMGroupPolicyAttachment,
	_ResourceTypeLowerName[1775:1806]: IAMGroupPolicyAttachment,
	_ResourceTypeName[1806:1830]:      IAMInstanceProfile,
	_ResourceTypeLowerName[1806:1830]: IAMInstanceProfile,
	_ResourceTypeName[1830:1861]:      IAMOpenidConnectProvider,
	_ResourceTypeLowerName[1830:1861]: IAMOpenidConnectProvider,
	_ResourceTypeName[1861:1875]:      IAMPolicy,
	_ResourceTypeLowerName[1861:1875]: IAMPolicy,
	_ResourceTypeName[1875:1887]:      IAMRole,
	_ResourceTypeLowerName[1875:1887]: IAMRole,
	_ResourceTypeName[1887:1906]:      IAMRolePolicy,
	_ResourceTypeLowerName[1887:1906]: IAMRolePolicy,
	_ResourceTypeName[1906:1936]:      IAMRolePolicyAttachment,
	_ResourceTypeLowerName[1906:1936]: IAMRolePolicyAttachment,
	_ResourceTypeName[1936:1957]:      IAMSAMLProvider,
	_ResourceTypeLowerName[1936:1957]: IAMSAMLProvider,
	_ResourceTypeName[1957:1983]:      IAMServerCertificate,
	_ResourceTypeLowerName[1957:1983]: IAMServerCertificate,
	_ResourceTypeName[1983:1995]:      IAMUser,
	_ResourceTypeLowerName[1983:1995]: IAMUser,
	_ResourceTypeName[1995:2024]:      IAMUserGroupMembership,
	_ResourceTypeLowerName[1995:2024]: IAMUserGroupMembership,
	_ResourceTypeName[2024:2043]:      IAMUserPolicy,
	_ResourceTypeLowerName[2024:2043]: IAMUserPolicy,
	_ResourceTypeName[2043:2073]:      IAMUserPolicyAttachment,
	_ResourceTypeLowerName[2043:2073]: IAMUserPolicyAttachment,
	_ResourceTypeName[2073:2093]:      IAMUserSSHKey,
	_ResourceTypeLowerName[2073:2093]: IAMUserSSHKey,
	_ResourceTypeName[2093:2113]:      InternetGateway,
	_ResourceTypeLowerName[2093:2113]: InternetGateway,
	_ResourceTypeName[2113:2125]:      KeyPair,
	_ResourceTypeLowerName[2113:2125]: KeyPair,
	_ResourceTypeName[2125:2143]:      KinesisStream,
	_ResourceTypeLowerName[2125:2143]: KinesisStream,
	_ResourceTypeName[2143:2162]:      LambdaFunction,
	_ResourceTypeLowerName[2143:2162]: LambdaFunction,
	_ResourceTypeName[2162:2186]:      LaunchConfiguration,
	_ResourceTypeLowerName[2162:2186]: LaunchConfiguration,
	_ResourceTypeName[2186:2205]:      LaunchTemplate,
	_ResourceTypeLowerName[2186:2205]: LaunchTemplate,
	_ResourceTypeName[2205:2211]:      LB,
	_ResourceTypeLowerName[2205:2211]: LB,
	_ResourceTypeName[2211:2242]:      LBCookieStickinessPolicy,
	_ResourceTypeLowerName[2211:2242]: LBCookieStickinessPolicy,
	_ResourceTypeName[2242:2257]:      LBListener,
	_ResourceTypeLowerName[2242:2257]: LBListener,
	_ResourceTypeName[2257:2284]:      LBListenerCertificate,
	_ResourceTypeLowerName[2257:2284]: LBListenerCertificate,
	_ResourceTypeName[2284:2304]:      LBListenerRule,
	_ResourceTypeLowerName[2284:2304]: LBListenerRule,
	_ResourceTypeName[2304:2323]:      LBTargetGroup,
	_ResourceTypeLowerName[2304:2323]: LBTargetGroup,
	_ResourceTypeName[2323:2353]:      LBTargetGroupAttachment,
	_ResourceTypeLowerName[2323:2353]: LBTargetGroupAttachment,
	_ResourceTypeName[2353:2375]:      LightsailInstance,
	_ResourceTypeLowerName[2353:2375]: LightsailInstance,
	_ResourceTypeName[2375:2400]:      MediaStoreContainer,
	_ResourceTypeLowerName[2375:2400]: MediaStoreContainer,
	_ResourceTypeName[2400:2413]:      MQBroker,
	_ResourceTypeLowerName[2400:2413]: MQBroker,
	_ResourceTypeName[2413:2428]:      NatGateway,
	_ResourceTypeLowerName[2413:2428]: NatGateway,
	_ResourceTypeName[2428:2443]:      NetworkAcl,
	_ResourceTypeLowerName[2428:2443]: NetworkAcl,
	_ResourceTypeName[2443:2471]:      NetworkfirewallFirewall,
	_ResourceTypeLowerName[2443:2471]: NetworkfirewallFirewall,
	_ResourceTypeName[2471:2506]:      NetworkfirewallFirewallPolicy,
	_ResourceTypeLowerName[2471:2506]: NetworkfirewallFirewallPolicy,
	_ResourceTypeName[2506:2536]:      NetworkfirewallRuleGroup,
	_ResourceTypeLowerName[2506:2536]: NetworkfirewallRuleGroup,
	_ResourceTypeName[2536:2555]:      NeptuneCluster,
	_ResourceTypeLowerName[2536:2555]: NeptuneCluster,
	_ResourceTypeName[2555:2570]:      RDSCluster,
	_ResourceTypeLowerName[2555:2570]: RDSCluster,
	_ResourceTypeName[2570:2592]:      RDSGlobalCluster,
	_ResourceTypeLowerName[2570:2592]: RDSGlobalCluster,
	_ResourceTypeName[2592:2612]:      RedshiftCluster,
	_ResourceTypeLowerName[2592:2612]: RedshiftCluster,
	_ResourceTypeName[2612:2638]:      Route53DelegationSet,
	_ResourceTypeLowerName[2612:2638]: Route53DelegationSet,
	_ResourceTypeName[2638:2662]:      Route53HealthCheck,
	_ResourceTypeLowerName[2638:2662]: Route53HealthCheck,
	_ResourceTypeName[2662:2683]:      Route53QueryLog,
	_ResourceTypeLowerName[2662:2683]: Route53QueryLog,
	_ResourceTypeName[2683:2701]:      Route53Record,
	_ResourceTypeLowerName[2683:2701]: Route53Record,
	_ResourceTypeName[2701:2730]:      Route53ResolverEndpoint,
	_ResourceTypeLowerName[2701:2730]: Route53ResolverEndpoint,
	_ResourceTypeName[2730:2767]:      Route53ResolverRuleAssociation,
	_ResourceTypeLowerName[2730:2767]: Route53ResolverRuleAssociation,
	_ResourceTypeName[2767:2783]:      Route53Zone,
	_ResourceTypeLowerName[2767:2783]: Route53Zone,
	_ResourceTypeName[2783:2811]:      Route53ZoneAssociation,
	_ResourceTypeLowerName[2783:2811]: Route53ZoneAssociation,
	_ResourceTypeName[2811:2826]:      RouteTable,
	_ResourceTypeLowerName[2811:2826]: RouteTable,
	_ResourceTypeName[2826:2839]:      S3Bucket,
	_ResourceTypeLowerName[2826:2839]: S3Bucket,
	_ResourceTypeName[2839:2857]:      SecurityGroup,
	_ResourceTypeLowerName[2839:2857]: SecurityGroup,
	_ResourceTypeName[2857:2888]:      SESActiveReceiptRuleSet,
	_ResourceTypeLowerName[2857:2888]: SESActiveReceiptRuleSet,
	_ResourceTypeName[2888:2913]:      SESConfigurationSet,
	_ResourceTypeLowerName[2888:2913]: SESConfigurationSet,
	_ResourceTypeName[2913:2932]:      SESDomainDKIM,
	_ResourceTypeLowerName[2913:2932]: SESDomainDKIM,
	_ResourceTypeName[2932:2955]:      SESDomainIdentity,
	_ResourceTypeLowerName[2932:2955]: SESDomainIdentity,
	_ResourceTypeName[2955:2979]:      SESDomainMailFrom,
	_ResourceTypeLowerName[2955:2979]: SESDomainMailFrom,
	_ResourceTypeName[2979:3014]:      SESIdentityNotificationTopic,
	_ResourceTypeLowerName[2979:3014]: SESIdentityNotificationTopic,
	_ResourceTypeName[3014:3036]:      SESReceiptFilter,
	_ResourceTypeLowerName[3014:3036]: SESReceiptFilter,
	_ResourceTypeName[3036:3056]:      SESReceiptRule,
	_ResourceTypeLowerName[3036:3056]: SESReceiptRule,
	_ResourceTypeName[3056:3080]:      SESReceiptRuleSet,
	_ResourceTypeLowerName[3056:3080]: SESReceiptRuleSet,
	_ResourceTypeName[3080:3096]:      SESTemplate,
	_ResourceTypeLowerName[3080:3096]: SESTemplate,
	_ResourceTypeName[3096:3109]:      SQSQueue,
	_ResourceTypeLowerName[3096:3109]: SQSQueue,
	_ResourceTypeName[3109:3135]:      StoragegatewayGateway,
	_ResourceTypeLowerName[3109:3135]: StoragegatewayGateway,
	_ResourceTypeName[3135:3145]:      Subnet,
	_ResourceTypeLowerName[3135:3145]: Subnet,
	_ResourceTypeName[3145:3166]:      VolumeAttachment,
	_ResourceTypeLowerName[3145:3166]: VolumeAttachment,
	_ResourceTypeName[3166:3173]:      VPC,
	_ResourceTypeLowerName[3166:3173]: VPC,
	_ResourceTypeName[3173:3193]:      VPCDhcpOptions,
	_ResourceTypeLowerName[3173:3193]: VPCDhcpOptions,
	_ResourceTypeName[3193:3209]:      VPCEndpoint,
	_ResourceTypeLowerName[3193:3209]: VPCEndpoint,
	_ResourceTypeName[3209:3221]:      VPCIpam,
	_ResourceTypeLowerName[3209:3221]: VPCIpam,
	_ResourceTypeName[3221:3238]:      VPCIpamPool,
	_ResourceTypeLowerName[3221:3238]: VPCIpamPool,
	_ResourceTypeName[3238:3256]:      VPCIpamScope,
	_ResourceTypeLowerName[3238:3256]: VPCIpamScope,
	_ResourceTypeName[3256:3282]:      VPCPeeringConnection,
	_ResourceTypeLowerName[3256:3282]: VPCPeeringConnection,
	_ResourceTypeName[3282:3297]:      VPNGateway,
	_ResourceTypeLowerName[3282:3297]: VPNGateway,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:12],
	_ResourceTypeName[12:19],
	_ResourceTypeName[19:26],
	_ResourceTypeName[26:42],
	_ResourceTypeName[42:70],
	_ResourceTypeName[70:91],
	_ResourceTypeName[91:111],
	_ResourceTypeName[111:142],
	_ResourceTypeName[142:168],
	_ResourceTypeName[168:192],
	_ResourceTypeName[192:216],
	_ResourceTypeName[216:237],
	_ResourceTypeName[237:257],
	_ResourceTypeName[257:278],
	_ResourceTypeName[278:300],
	_ResourceTypeName[300:324],
	_ResourceTypeName[324:348],
	_ResourceTypeName[348:375],
	_ResourceTypeName[375:412],
	_ResourceTypeName[412:437],
	_ResourceTypeName[437:464],
	_ResourceTypeName[464:488],
	_ResourceTypeName[488:512],
	_ResourceTypeName[512:537],
	_ResourceTypeName[537:552],
	_ResourceTypeName[552:567],
	_ResourceTypeName[567:589],
	_ResourceTypeName[589:608],
	_ResourceTypeName[608:639],
	_ResourceTypeName[639:667],
	_ResourceTypeName[667:684],
	_ResourceTypeName[684:698],
	_ResourceTypeName[698:730],
	_ResourceTypeName[730:761],
	_ResourceTypeName[761:771],
	_ResourceTypeName[771:796],
	_ResourceTypeName[796:814],
	_ResourceTypeName[814:830],
	_ResourceTypeName[830:844],
	_ResourceTypeName[844:859],
	_ResourceTypeName[859:874],
	_ResourceTypeName[874:897],
	_ResourceTypeName[897:924],
	_ResourceTypeName[924:947],
	_ResourceTypeName[947:985],
	_ResourceTypeName[985:1020],
	_ResourceTypeName[1020:1060],
	_ResourceTypeName[1060:1102],
	_ResourceTypeName[1102:1153],
	_ResourceTypeName[1153:1198],
	_ResourceTypeName[1198:1227],
	_ResourceTypeName[1227:1274],
	_ResourceTypeName[1274:1321],
	_ResourceTypeName[1321:1368],
	_ResourceTypeName[1368:1387],
	_ResourceTypeName[1387:1394],
	_ResourceTypeName[1394:1409],
	_ResourceTypeName[1409:1432],
	_ResourceTypeName[1432:1465],
	_ResourceTypeName[1465:1498],
	_ResourceTypeName[1498:1522],
	_ResourceTypeName[1522:1553],
	_ResourceTypeName[1553:1560],
	_ResourceTypeName[1560:1575],
	_ResourceTypeName[1575:1601],
	_ResourceTypeName[1601:1626],
	_ResourceTypeName[1626:1648],
	_ResourceTypeName[1648:1666],
	_ResourceTypeName[1666:1687],
	_ResourceTypeName[1687:1718],
	_ResourceTypeName[1718:1731],
	_ResourceTypeName[1731:1755],
	_ResourceTypeName[1755:1775],
	_ResourceTypeName[1775:1806],
	_ResourceTypeName[1806:1830],
	_ResourceTypeName[1830:1861],
	_ResourceTypeName[1861:1875],
	_ResourceTypeName[1875:1887],
	_ResourceTypeName[1887:1906],
	_ResourceTypeName[1906:1936],
	_ResourceTypeName[1936:1957],
	_ResourceTypeName[1957:1983],
	_ResourceTypeName[1983:1995],
	_ResourceTypeName[1995:2024],
	_ResourceTypeName[2024:2043],
	_ResourceTypeName[2043:2073],
	_ResourceTypeName[2073:2093],
	_ResourceTypeName[2093:2113],
	_ResourceTypeName[2113:2125],
	_ResourceTypeName[2125:2143],
	_ResourceTypeName[2143:2162],
	_ResourceTypeName[2162:2186],
	_ResourceTypeName[2186:2205],
	_ResourceTypeName[2205:2211],
	_ResourceTypeName[2211:2242],
	_ResourceTypeName[2242:2257],
	_ResourceTypeName[2257:2284],
	_ResourceTypeName[2284:2304],
	_ResourceTypeName[2304:2323],
	_ResourceTypeName[2323:2353],
	_ResourceTypeName[2353:2375],
	_ResourceTypeName[2375:2400],
	_ResourceTypeName[2400:2413],
	_ResourceTypeName[2413:2428],
	_ResourceTypeName[2428:2443],
	_ResourceTypeName[2443:2471],
	_ResourceTypeName[2471:2506],
	_ResourceTypeName[2506:2536],
	_ResourceTypeName[2536:2555],
	_ResourceTypeName[2555:2570],
	_ResourceTypeName[2570:2592],
	_ResourceTypeName[2592:2612],
	_ResourceTypeName[2612:2638],
	_ResourceTypeName[2638:2662],
	_ResourceTypeName[2662:2683],
	_ResourceTypeName[2683:2701],
	_ResourceTypeName[2701:2730],
	_ResourceTypeName[2730:2767],
	_ResourceTypeName[2767:2783],
	_ResourceTypeName[2783:2811],
	_ResourceTypeName[2811:2826],
	_ResourceTypeName[2826:2839],
	_ResourceTypeName[2839:2857],
	_ResourceTypeName[2857:2888],
	_ResourceTypeName[2888:2913],
	_ResourceTypeName[2913:2932],
	_ResourceTypeName[2932:2955],
	_ResourceTypeName[2955:2979],
	_ResourceTypeName[2979:3014],
	_ResourceTypeName[3014:3036],
	_ResourceTypeName[3036:3056],
	_ResourceTypeName[3056:3080],
	_ResourceTypeName[3080:3096],
	_ResourceTypeName[3096:3109],
	_ResourceTypeName[3109:3135],
	_ResourceTypeName[3135:3145],
	_ResourceTypeName[3145:3166],
	_ResourceTypeName[3166:3173],
	_ResourceTypeName[3173:3193],
	_ResourceTypeName[3193:3209],
	_ResourceTypeName[3209:3221],
	_ResourceTypeName[3221:3238],
	_ResourceTypeName[3238:3256],
	_ResourceTypeName[3256:3282],
	_ResourceTypeName[3282:3297],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ResourceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
