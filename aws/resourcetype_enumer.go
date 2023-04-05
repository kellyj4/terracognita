// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "aws_instanceaws_amiaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_connectionaws_dx_gatewayaws_dx_private_virtual_interfaceaws_dx_public_virtual_interfaceaws_dx_lagaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_snapshotaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_managed_prefix_listaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_network_aclaws_networkfirewall_firewallaws_networkfirewall_firewall_policyaws_networkfirewall_rule_groupaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_dhcp_optionsaws_vpc_endpointaws_vpc_ipamaws_vpc_ipam_poolaws_vpc_peering_connectionaws_vpn_gateway"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 26, 42, 70, 91, 111, 142, 168, 192, 216, 237, 257, 278, 300, 324, 348, 375, 412, 437, 464, 479, 494, 516, 535, 566, 594, 611, 625, 657, 688, 698, 723, 741, 757, 771, 786, 801, 824, 851, 874, 912, 947, 987, 1029, 1080, 1125, 1154, 1201, 1248, 1295, 1314, 1321, 1336, 1359, 1392, 1425, 1449, 1480, 1487, 1502, 1528, 1553, 1575, 1593, 1614, 1645, 1658, 1682, 1702, 1733, 1757, 1788, 1802, 1814, 1833, 1863, 1884, 1910, 1922, 1951, 1970, 2000, 2020, 2040, 2052, 2070, 2089, 2113, 2132, 2138, 2169, 2184, 2211, 2231, 2250, 2280, 2302, 2327, 2340, 2355, 2370, 2398, 2433, 2463, 2482, 2497, 2519, 2539, 2565, 2589, 2610, 2628, 2657, 2694, 2710, 2738, 2753, 2766, 2784, 2815, 2840, 2859, 2882, 2906, 2941, 2963, 2983, 3007, 3023, 3036, 3062, 3072, 3093, 3100, 3120, 3136, 3148, 3165, 3191, 3206}

const _ResourceTypeLowerName = "aws_instanceaws_amiaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_autoscaling_scheduleaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_connectionaws_dx_gatewayaws_dx_private_virtual_interfaceaws_dx_public_virtual_interfaceaws_dx_lagaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_snapshotaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_ecs_task_definitionaws_ec2_managed_prefix_listaws_ec2_transit_gatewayaws_ec2_transit_gateway_vpc_attachmentaws_ec2_transit_gateway_route_tableaws_ec2_transit_gateway_multicast_domainaws_ec2_transit_gateway_peering_attachmentaws_ec2_transit_gateway_peering_attachment_accepteraws_ec2_transit_gateway_prefix_list_referenceaws_ec2_transit_gateway_routeaws_ec2_transit_gateway_route_table_associationaws_ec2_transit_gateway_route_table_propagationaws_ec2_transit_gateway_vpc_attachment_accepteraws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_network_aclaws_networkfirewall_firewallaws_networkfirewall_firewall_policyaws_networkfirewall_rule_groupaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_route_tableaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_dhcp_optionsaws_vpc_endpointaws_vpc_ipamaws_vpc_ipam_poolaws_vpc_peering_connectionaws_vpn_gateway"

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
	_ = x[DaxCluster-(22)]
	_ = x[DBInstance-(23)]
	_ = x[DBParameterGroup-(24)]
	_ = x[DBSubnetGroup-(25)]
	_ = x[DirectoryServiceDirectory-(26)]
	_ = x[DmsReplicationInstance-(27)]
	_ = x[DXConnection-(28)]
	_ = x[DXGateway-(29)]
	_ = x[DXPrivateVirtualInterface-(30)]
	_ = x[DXPublicVirtualInterface-(31)]
	_ = x[DXLag-(32)]
	_ = x[DynamodbGlobalTable-(33)]
	_ = x[DynamodbTable-(34)]
	_ = x[EBSSnapshot-(35)]
	_ = x[EBSVolume-(36)]
	_ = x[ECSCluster-(37)]
	_ = x[ECSService-(38)]
	_ = x[ECSTaskDefinition-(39)]
	_ = x[EC2ManagedPrefixList-(40)]
	_ = x[EC2TransitGateway-(41)]
	_ = x[EC2TransitGatewayVPCAttachment-(42)]
	_ = x[EC2TransitGatewayRouteTable-(43)]
	_ = x[EC2TransitGatewayMulticastDomain-(44)]
	_ = x[EC2TransitGatewayPeeringAttachment-(45)]
	_ = x[EC2TransitGatewayPeeringAttachmentAccepter-(46)]
	_ = x[EC2TransitGatewayPrefixListReference-(47)]
	_ = x[EC2TransitGatewayRoute-(48)]
	_ = x[EC2TransitGatewayRouteTableAssociation-(49)]
	_ = x[EC2TransitGatewayRouteTablePropagation-(50)]
	_ = x[EC2TransitGatewayVPCAttachmentAccepter-(51)]
	_ = x[EFSFileSystem-(52)]
	_ = x[EIP-(53)]
	_ = x[EKSCluster-(54)]
	_ = x[ElasticacheCluster-(55)]
	_ = x[ElasticacheReplicationGroup-(56)]
	_ = x[ElasticBeanstalkApplication-(57)]
	_ = x[ElasticsearchDomain-(58)]
	_ = x[ElasticsearchDomainPolicy-(59)]
	_ = x[ELB-(60)]
	_ = x[EMRCluster-(61)]
	_ = x[FsxLustreFileSystem-(62)]
	_ = x[GlueCatalogDatabase-(63)]
	_ = x[GlueCatalogTable-(64)]
	_ = x[IAMAccessKey-(65)]
	_ = x[IAMAccountAlias-(66)]
	_ = x[IAMAccountPasswordPolicy-(67)]
	_ = x[IAMGroup-(68)]
	_ = x[IAMGroupMembership-(69)]
	_ = x[IAMGroupPolicy-(70)]
	_ = x[IAMGroupPolicyAttachment-(71)]
	_ = x[IAMInstanceProfile-(72)]
	_ = x[IAMOpenidConnectProvider-(73)]
	_ = x[IAMPolicy-(74)]
	_ = x[IAMRole-(75)]
	_ = x[IAMRolePolicy-(76)]
	_ = x[IAMRolePolicyAttachment-(77)]
	_ = x[IAMSAMLProvider-(78)]
	_ = x[IAMServerCertificate-(79)]
	_ = x[IAMUser-(80)]
	_ = x[IAMUserGroupMembership-(81)]
	_ = x[IAMUserPolicy-(82)]
	_ = x[IAMUserPolicyAttachment-(83)]
	_ = x[IAMUserSSHKey-(84)]
	_ = x[InternetGateway-(85)]
	_ = x[KeyPair-(86)]
	_ = x[KinesisStream-(87)]
	_ = x[LambdaFunction-(88)]
	_ = x[LaunchConfiguration-(89)]
	_ = x[LaunchTemplate-(90)]
	_ = x[LB-(91)]
	_ = x[LBCookieStickinessPolicy-(92)]
	_ = x[LBListener-(93)]
	_ = x[LBListenerCertificate-(94)]
	_ = x[LBListenerRule-(95)]
	_ = x[LBTargetGroup-(96)]
	_ = x[LBTargetGroupAttachment-(97)]
	_ = x[LightsailInstance-(98)]
	_ = x[MediaStoreContainer-(99)]
	_ = x[MQBroker-(100)]
	_ = x[NatGateway-(101)]
	_ = x[NetworkAcl-(102)]
	_ = x[NetworkfirewallFirewall-(103)]
	_ = x[NetworkfirewallFirewallPolicy-(104)]
	_ = x[NetworkfirewallRuleGroup-(105)]
	_ = x[NeptuneCluster-(106)]
	_ = x[RDSCluster-(107)]
	_ = x[RDSGlobalCluster-(108)]
	_ = x[RedshiftCluster-(109)]
	_ = x[Route53DelegationSet-(110)]
	_ = x[Route53HealthCheck-(111)]
	_ = x[Route53QueryLog-(112)]
	_ = x[Route53Record-(113)]
	_ = x[Route53ResolverEndpoint-(114)]
	_ = x[Route53ResolverRuleAssociation-(115)]
	_ = x[Route53Zone-(116)]
	_ = x[Route53ZoneAssociation-(117)]
	_ = x[RouteTable-(118)]
	_ = x[S3Bucket-(119)]
	_ = x[SecurityGroup-(120)]
	_ = x[SESActiveReceiptRuleSet-(121)]
	_ = x[SESConfigurationSet-(122)]
	_ = x[SESDomainDKIM-(123)]
	_ = x[SESDomainIdentity-(124)]
	_ = x[SESDomainMailFrom-(125)]
	_ = x[SESIdentityNotificationTopic-(126)]
	_ = x[SESReceiptFilter-(127)]
	_ = x[SESReceiptRule-(128)]
	_ = x[SESReceiptRuleSet-(129)]
	_ = x[SESTemplate-(130)]
	_ = x[SQSQueue-(131)]
	_ = x[StoragegatewayGateway-(132)]
	_ = x[Subnet-(133)]
	_ = x[VolumeAttachment-(134)]
	_ = x[VPC-(135)]
	_ = x[VPCDhcpOptions-(136)]
	_ = x[VPCEndpoint-(137)]
	_ = x[VPCIpam-(138)]
	_ = x[VPCIpamPool-(139)]
	_ = x[VPCPeeringConnection-(140)]
	_ = x[VPNGateway-(141)]
}

var _ResourceTypeValues = []ResourceType{Instance, AMI, ALB, ALBListener, ALBListenerCertificate, ALBListenerRule, ALBTargetGroup, ALBTargetGroupAttachment, APIGatewayDeployment, APIGatewayResource, APIGatewayRestAPI, APIGatewayStage, AthenaWorkgroup, AutoscalingGroup, AutoscalingPolicy, AutoscalingSchedule, BatchJobDefinition, CloudfrontDistribution, CloudfrontOriginAccessIdentity, CloudfrontPublicKey, CloudwatchMetricAlarm, DaxCluster, DBInstance, DBParameterGroup, DBSubnetGroup, DirectoryServiceDirectory, DmsReplicationInstance, DXConnection, DXGateway, DXPrivateVirtualInterface, DXPublicVirtualInterface, DXLag, DynamodbGlobalTable, DynamodbTable, EBSSnapshot, EBSVolume, ECSCluster, ECSService, ECSTaskDefinition, EC2ManagedPrefixList, EC2TransitGateway, EC2TransitGatewayVPCAttachment, EC2TransitGatewayRouteTable, EC2TransitGatewayMulticastDomain, EC2TransitGatewayPeeringAttachment, EC2TransitGatewayPeeringAttachmentAccepter, EC2TransitGatewayPrefixListReference, EC2TransitGatewayRoute, EC2TransitGatewayRouteTableAssociation, EC2TransitGatewayRouteTablePropagation, EC2TransitGatewayVPCAttachmentAccepter, EFSFileSystem, EIP, EKSCluster, ElasticacheCluster, ElasticacheReplicationGroup, ElasticBeanstalkApplication, ElasticsearchDomain, ElasticsearchDomainPolicy, ELB, EMRCluster, FsxLustreFileSystem, GlueCatalogDatabase, GlueCatalogTable, IAMAccessKey, IAMAccountAlias, IAMAccountPasswordPolicy, IAMGroup, IAMGroupMembership, IAMGroupPolicy, IAMGroupPolicyAttachment, IAMInstanceProfile, IAMOpenidConnectProvider, IAMPolicy, IAMRole, IAMRolePolicy, IAMRolePolicyAttachment, IAMSAMLProvider, IAMServerCertificate, IAMUser, IAMUserGroupMembership, IAMUserPolicy, IAMUserPolicyAttachment, IAMUserSSHKey, InternetGateway, KeyPair, KinesisStream, LambdaFunction, LaunchConfiguration, LaunchTemplate, LB, LBCookieStickinessPolicy, LBListener, LBListenerCertificate, LBListenerRule, LBTargetGroup, LBTargetGroupAttachment, LightsailInstance, MediaStoreContainer, MQBroker, NatGateway, NetworkAcl, NetworkfirewallFirewall, NetworkfirewallFirewallPolicy, NetworkfirewallRuleGroup, NeptuneCluster, RDSCluster, RDSGlobalCluster, RedshiftCluster, Route53DelegationSet, Route53HealthCheck, Route53QueryLog, Route53Record, Route53ResolverEndpoint, Route53ResolverRuleAssociation, Route53Zone, Route53ZoneAssociation, RouteTable, S3Bucket, SecurityGroup, SESActiveReceiptRuleSet, SESConfigurationSet, SESDomainDKIM, SESDomainIdentity, SESDomainMailFrom, SESIdentityNotificationTopic, SESReceiptFilter, SESReceiptRule, SESReceiptRuleSet, SESTemplate, SQSQueue, StoragegatewayGateway, Subnet, VolumeAttachment, VPC, VPCDhcpOptions, VPCEndpoint, VPCIpam, VPCIpamPool, VPCPeeringConnection, VPNGateway}

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
	_ResourceTypeName[464:479]:        DaxCluster,
	_ResourceTypeLowerName[464:479]:   DaxCluster,
	_ResourceTypeName[479:494]:        DBInstance,
	_ResourceTypeLowerName[479:494]:   DBInstance,
	_ResourceTypeName[494:516]:        DBParameterGroup,
	_ResourceTypeLowerName[494:516]:   DBParameterGroup,
	_ResourceTypeName[516:535]:        DBSubnetGroup,
	_ResourceTypeLowerName[516:535]:   DBSubnetGroup,
	_ResourceTypeName[535:566]:        DirectoryServiceDirectory,
	_ResourceTypeLowerName[535:566]:   DirectoryServiceDirectory,
	_ResourceTypeName[566:594]:        DmsReplicationInstance,
	_ResourceTypeLowerName[566:594]:   DmsReplicationInstance,
	_ResourceTypeName[594:611]:        DXConnection,
	_ResourceTypeLowerName[594:611]:   DXConnection,
	_ResourceTypeName[611:625]:        DXGateway,
	_ResourceTypeLowerName[611:625]:   DXGateway,
	_ResourceTypeName[625:657]:        DXPrivateVirtualInterface,
	_ResourceTypeLowerName[625:657]:   DXPrivateVirtualInterface,
	_ResourceTypeName[657:688]:        DXPublicVirtualInterface,
	_ResourceTypeLowerName[657:688]:   DXPublicVirtualInterface,
	_ResourceTypeName[688:698]:        DXLag,
	_ResourceTypeLowerName[688:698]:   DXLag,
	_ResourceTypeName[698:723]:        DynamodbGlobalTable,
	_ResourceTypeLowerName[698:723]:   DynamodbGlobalTable,
	_ResourceTypeName[723:741]:        DynamodbTable,
	_ResourceTypeLowerName[723:741]:   DynamodbTable,
	_ResourceTypeName[741:757]:        EBSSnapshot,
	_ResourceTypeLowerName[741:757]:   EBSSnapshot,
	_ResourceTypeName[757:771]:        EBSVolume,
	_ResourceTypeLowerName[757:771]:   EBSVolume,
	_ResourceTypeName[771:786]:        ECSCluster,
	_ResourceTypeLowerName[771:786]:   ECSCluster,
	_ResourceTypeName[786:801]:        ECSService,
	_ResourceTypeLowerName[786:801]:   ECSService,
	_ResourceTypeName[801:824]:        ECSTaskDefinition,
	_ResourceTypeLowerName[801:824]:   ECSTaskDefinition,
	_ResourceTypeName[824:851]:        EC2ManagedPrefixList,
	_ResourceTypeLowerName[824:851]:   EC2ManagedPrefixList,
	_ResourceTypeName[851:874]:        EC2TransitGateway,
	_ResourceTypeLowerName[851:874]:   EC2TransitGateway,
	_ResourceTypeName[874:912]:        EC2TransitGatewayVPCAttachment,
	_ResourceTypeLowerName[874:912]:   EC2TransitGatewayVPCAttachment,
	_ResourceTypeName[912:947]:        EC2TransitGatewayRouteTable,
	_ResourceTypeLowerName[912:947]:   EC2TransitGatewayRouteTable,
	_ResourceTypeName[947:987]:        EC2TransitGatewayMulticastDomain,
	_ResourceTypeLowerName[947:987]:   EC2TransitGatewayMulticastDomain,
	_ResourceTypeName[987:1029]:       EC2TransitGatewayPeeringAttachment,
	_ResourceTypeLowerName[987:1029]:  EC2TransitGatewayPeeringAttachment,
	_ResourceTypeName[1029:1080]:      EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeLowerName[1029:1080]: EC2TransitGatewayPeeringAttachmentAccepter,
	_ResourceTypeName[1080:1125]:      EC2TransitGatewayPrefixListReference,
	_ResourceTypeLowerName[1080:1125]: EC2TransitGatewayPrefixListReference,
	_ResourceTypeName[1125:1154]:      EC2TransitGatewayRoute,
	_ResourceTypeLowerName[1125:1154]: EC2TransitGatewayRoute,
	_ResourceTypeName[1154:1201]:      EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeLowerName[1154:1201]: EC2TransitGatewayRouteTableAssociation,
	_ResourceTypeName[1201:1248]:      EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeLowerName[1201:1248]: EC2TransitGatewayRouteTablePropagation,
	_ResourceTypeName[1248:1295]:      EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeLowerName[1248:1295]: EC2TransitGatewayVPCAttachmentAccepter,
	_ResourceTypeName[1295:1314]:      EFSFileSystem,
	_ResourceTypeLowerName[1295:1314]: EFSFileSystem,
	_ResourceTypeName[1314:1321]:      EIP,
	_ResourceTypeLowerName[1314:1321]: EIP,
	_ResourceTypeName[1321:1336]:      EKSCluster,
	_ResourceTypeLowerName[1321:1336]: EKSCluster,
	_ResourceTypeName[1336:1359]:      ElasticacheCluster,
	_ResourceTypeLowerName[1336:1359]: ElasticacheCluster,
	_ResourceTypeName[1359:1392]:      ElasticacheReplicationGroup,
	_ResourceTypeLowerName[1359:1392]: ElasticacheReplicationGroup,
	_ResourceTypeName[1392:1425]:      ElasticBeanstalkApplication,
	_ResourceTypeLowerName[1392:1425]: ElasticBeanstalkApplication,
	_ResourceTypeName[1425:1449]:      ElasticsearchDomain,
	_ResourceTypeLowerName[1425:1449]: ElasticsearchDomain,
	_ResourceTypeName[1449:1480]:      ElasticsearchDomainPolicy,
	_ResourceTypeLowerName[1449:1480]: ElasticsearchDomainPolicy,
	_ResourceTypeName[1480:1487]:      ELB,
	_ResourceTypeLowerName[1480:1487]: ELB,
	_ResourceTypeName[1487:1502]:      EMRCluster,
	_ResourceTypeLowerName[1487:1502]: EMRCluster,
	_ResourceTypeName[1502:1528]:      FsxLustreFileSystem,
	_ResourceTypeLowerName[1502:1528]: FsxLustreFileSystem,
	_ResourceTypeName[1528:1553]:      GlueCatalogDatabase,
	_ResourceTypeLowerName[1528:1553]: GlueCatalogDatabase,
	_ResourceTypeName[1553:1575]:      GlueCatalogTable,
	_ResourceTypeLowerName[1553:1575]: GlueCatalogTable,
	_ResourceTypeName[1575:1593]:      IAMAccessKey,
	_ResourceTypeLowerName[1575:1593]: IAMAccessKey,
	_ResourceTypeName[1593:1614]:      IAMAccountAlias,
	_ResourceTypeLowerName[1593:1614]: IAMAccountAlias,
	_ResourceTypeName[1614:1645]:      IAMAccountPasswordPolicy,
	_ResourceTypeLowerName[1614:1645]: IAMAccountPasswordPolicy,
	_ResourceTypeName[1645:1658]:      IAMGroup,
	_ResourceTypeLowerName[1645:1658]: IAMGroup,
	_ResourceTypeName[1658:1682]:      IAMGroupMembership,
	_ResourceTypeLowerName[1658:1682]: IAMGroupMembership,
	_ResourceTypeName[1682:1702]:      IAMGroupPolicy,
	_ResourceTypeLowerName[1682:1702]: IAMGroupPolicy,
	_ResourceTypeName[1702:1733]:      IAMGroupPolicyAttachment,
	_ResourceTypeLowerName[1702:1733]: IAMGroupPolicyAttachment,
	_ResourceTypeName[1733:1757]:      IAMInstanceProfile,
	_ResourceTypeLowerName[1733:1757]: IAMInstanceProfile,
	_ResourceTypeName[1757:1788]:      IAMOpenidConnectProvider,
	_ResourceTypeLowerName[1757:1788]: IAMOpenidConnectProvider,
	_ResourceTypeName[1788:1802]:      IAMPolicy,
	_ResourceTypeLowerName[1788:1802]: IAMPolicy,
	_ResourceTypeName[1802:1814]:      IAMRole,
	_ResourceTypeLowerName[1802:1814]: IAMRole,
	_ResourceTypeName[1814:1833]:      IAMRolePolicy,
	_ResourceTypeLowerName[1814:1833]: IAMRolePolicy,
	_ResourceTypeName[1833:1863]:      IAMRolePolicyAttachment,
	_ResourceTypeLowerName[1833:1863]: IAMRolePolicyAttachment,
	_ResourceTypeName[1863:1884]:      IAMSAMLProvider,
	_ResourceTypeLowerName[1863:1884]: IAMSAMLProvider,
	_ResourceTypeName[1884:1910]:      IAMServerCertificate,
	_ResourceTypeLowerName[1884:1910]: IAMServerCertificate,
	_ResourceTypeName[1910:1922]:      IAMUser,
	_ResourceTypeLowerName[1910:1922]: IAMUser,
	_ResourceTypeName[1922:1951]:      IAMUserGroupMembership,
	_ResourceTypeLowerName[1922:1951]: IAMUserGroupMembership,
	_ResourceTypeName[1951:1970]:      IAMUserPolicy,
	_ResourceTypeLowerName[1951:1970]: IAMUserPolicy,
	_ResourceTypeName[1970:2000]:      IAMUserPolicyAttachment,
	_ResourceTypeLowerName[1970:2000]: IAMUserPolicyAttachment,
	_ResourceTypeName[2000:2020]:      IAMUserSSHKey,
	_ResourceTypeLowerName[2000:2020]: IAMUserSSHKey,
	_ResourceTypeName[2020:2040]:      InternetGateway,
	_ResourceTypeLowerName[2020:2040]: InternetGateway,
	_ResourceTypeName[2040:2052]:      KeyPair,
	_ResourceTypeLowerName[2040:2052]: KeyPair,
	_ResourceTypeName[2052:2070]:      KinesisStream,
	_ResourceTypeLowerName[2052:2070]: KinesisStream,
	_ResourceTypeName[2070:2089]:      LambdaFunction,
	_ResourceTypeLowerName[2070:2089]: LambdaFunction,
	_ResourceTypeName[2089:2113]:      LaunchConfiguration,
	_ResourceTypeLowerName[2089:2113]: LaunchConfiguration,
	_ResourceTypeName[2113:2132]:      LaunchTemplate,
	_ResourceTypeLowerName[2113:2132]: LaunchTemplate,
	_ResourceTypeName[2132:2138]:      LB,
	_ResourceTypeLowerName[2132:2138]: LB,
	_ResourceTypeName[2138:2169]:      LBCookieStickinessPolicy,
	_ResourceTypeLowerName[2138:2169]: LBCookieStickinessPolicy,
	_ResourceTypeName[2169:2184]:      LBListener,
	_ResourceTypeLowerName[2169:2184]: LBListener,
	_ResourceTypeName[2184:2211]:      LBListenerCertificate,
	_ResourceTypeLowerName[2184:2211]: LBListenerCertificate,
	_ResourceTypeName[2211:2231]:      LBListenerRule,
	_ResourceTypeLowerName[2211:2231]: LBListenerRule,
	_ResourceTypeName[2231:2250]:      LBTargetGroup,
	_ResourceTypeLowerName[2231:2250]: LBTargetGroup,
	_ResourceTypeName[2250:2280]:      LBTargetGroupAttachment,
	_ResourceTypeLowerName[2250:2280]: LBTargetGroupAttachment,
	_ResourceTypeName[2280:2302]:      LightsailInstance,
	_ResourceTypeLowerName[2280:2302]: LightsailInstance,
	_ResourceTypeName[2302:2327]:      MediaStoreContainer,
	_ResourceTypeLowerName[2302:2327]: MediaStoreContainer,
	_ResourceTypeName[2327:2340]:      MQBroker,
	_ResourceTypeLowerName[2327:2340]: MQBroker,
	_ResourceTypeName[2340:2355]:      NatGateway,
	_ResourceTypeLowerName[2340:2355]: NatGateway,
	_ResourceTypeName[2355:2370]:      NetworkAcl,
	_ResourceTypeLowerName[2355:2370]: NetworkAcl,
	_ResourceTypeName[2370:2398]:      NetworkfirewallFirewall,
	_ResourceTypeLowerName[2370:2398]: NetworkfirewallFirewall,
	_ResourceTypeName[2398:2433]:      NetworkfirewallFirewallPolicy,
	_ResourceTypeLowerName[2398:2433]: NetworkfirewallFirewallPolicy,
	_ResourceTypeName[2433:2463]:      NetworkfirewallRuleGroup,
	_ResourceTypeLowerName[2433:2463]: NetworkfirewallRuleGroup,
	_ResourceTypeName[2463:2482]:      NeptuneCluster,
	_ResourceTypeLowerName[2463:2482]: NeptuneCluster,
	_ResourceTypeName[2482:2497]:      RDSCluster,
	_ResourceTypeLowerName[2482:2497]: RDSCluster,
	_ResourceTypeName[2497:2519]:      RDSGlobalCluster,
	_ResourceTypeLowerName[2497:2519]: RDSGlobalCluster,
	_ResourceTypeName[2519:2539]:      RedshiftCluster,
	_ResourceTypeLowerName[2519:2539]: RedshiftCluster,
	_ResourceTypeName[2539:2565]:      Route53DelegationSet,
	_ResourceTypeLowerName[2539:2565]: Route53DelegationSet,
	_ResourceTypeName[2565:2589]:      Route53HealthCheck,
	_ResourceTypeLowerName[2565:2589]: Route53HealthCheck,
	_ResourceTypeName[2589:2610]:      Route53QueryLog,
	_ResourceTypeLowerName[2589:2610]: Route53QueryLog,
	_ResourceTypeName[2610:2628]:      Route53Record,
	_ResourceTypeLowerName[2610:2628]: Route53Record,
	_ResourceTypeName[2628:2657]:      Route53ResolverEndpoint,
	_ResourceTypeLowerName[2628:2657]: Route53ResolverEndpoint,
	_ResourceTypeName[2657:2694]:      Route53ResolverRuleAssociation,
	_ResourceTypeLowerName[2657:2694]: Route53ResolverRuleAssociation,
	_ResourceTypeName[2694:2710]:      Route53Zone,
	_ResourceTypeLowerName[2694:2710]: Route53Zone,
	_ResourceTypeName[2710:2738]:      Route53ZoneAssociation,
	_ResourceTypeLowerName[2710:2738]: Route53ZoneAssociation,
	_ResourceTypeName[2738:2753]:      RouteTable,
	_ResourceTypeLowerName[2738:2753]: RouteTable,
	_ResourceTypeName[2753:2766]:      S3Bucket,
	_ResourceTypeLowerName[2753:2766]: S3Bucket,
	_ResourceTypeName[2766:2784]:      SecurityGroup,
	_ResourceTypeLowerName[2766:2784]: SecurityGroup,
	_ResourceTypeName[2784:2815]:      SESActiveReceiptRuleSet,
	_ResourceTypeLowerName[2784:2815]: SESActiveReceiptRuleSet,
	_ResourceTypeName[2815:2840]:      SESConfigurationSet,
	_ResourceTypeLowerName[2815:2840]: SESConfigurationSet,
	_ResourceTypeName[2840:2859]:      SESDomainDKIM,
	_ResourceTypeLowerName[2840:2859]: SESDomainDKIM,
	_ResourceTypeName[2859:2882]:      SESDomainIdentity,
	_ResourceTypeLowerName[2859:2882]: SESDomainIdentity,
	_ResourceTypeName[2882:2906]:      SESDomainMailFrom,
	_ResourceTypeLowerName[2882:2906]: SESDomainMailFrom,
	_ResourceTypeName[2906:2941]:      SESIdentityNotificationTopic,
	_ResourceTypeLowerName[2906:2941]: SESIdentityNotificationTopic,
	_ResourceTypeName[2941:2963]:      SESReceiptFilter,
	_ResourceTypeLowerName[2941:2963]: SESReceiptFilter,
	_ResourceTypeName[2963:2983]:      SESReceiptRule,
	_ResourceTypeLowerName[2963:2983]: SESReceiptRule,
	_ResourceTypeName[2983:3007]:      SESReceiptRuleSet,
	_ResourceTypeLowerName[2983:3007]: SESReceiptRuleSet,
	_ResourceTypeName[3007:3023]:      SESTemplate,
	_ResourceTypeLowerName[3007:3023]: SESTemplate,
	_ResourceTypeName[3023:3036]:      SQSQueue,
	_ResourceTypeLowerName[3023:3036]: SQSQueue,
	_ResourceTypeName[3036:3062]:      StoragegatewayGateway,
	_ResourceTypeLowerName[3036:3062]: StoragegatewayGateway,
	_ResourceTypeName[3062:3072]:      Subnet,
	_ResourceTypeLowerName[3062:3072]: Subnet,
	_ResourceTypeName[3072:3093]:      VolumeAttachment,
	_ResourceTypeLowerName[3072:3093]: VolumeAttachment,
	_ResourceTypeName[3093:3100]:      VPC,
	_ResourceTypeLowerName[3093:3100]: VPC,
	_ResourceTypeName[3100:3120]:      VPCDhcpOptions,
	_ResourceTypeLowerName[3100:3120]: VPCDhcpOptions,
	_ResourceTypeName[3120:3136]:      VPCEndpoint,
	_ResourceTypeLowerName[3120:3136]: VPCEndpoint,
	_ResourceTypeName[3136:3148]:      VPCIpam,
	_ResourceTypeLowerName[3136:3148]: VPCIpam,
	_ResourceTypeName[3148:3165]:      VPCIpamPool,
	_ResourceTypeLowerName[3148:3165]: VPCIpamPool,
	_ResourceTypeName[3165:3191]:      VPCPeeringConnection,
	_ResourceTypeLowerName[3165:3191]: VPCPeeringConnection,
	_ResourceTypeName[3191:3206]:      VPNGateway,
	_ResourceTypeLowerName[3191:3206]: VPNGateway,
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
	_ResourceTypeName[464:479],
	_ResourceTypeName[479:494],
	_ResourceTypeName[494:516],
	_ResourceTypeName[516:535],
	_ResourceTypeName[535:566],
	_ResourceTypeName[566:594],
	_ResourceTypeName[594:611],
	_ResourceTypeName[611:625],
	_ResourceTypeName[625:657],
	_ResourceTypeName[657:688],
	_ResourceTypeName[688:698],
	_ResourceTypeName[698:723],
	_ResourceTypeName[723:741],
	_ResourceTypeName[741:757],
	_ResourceTypeName[757:771],
	_ResourceTypeName[771:786],
	_ResourceTypeName[786:801],
	_ResourceTypeName[801:824],
	_ResourceTypeName[824:851],
	_ResourceTypeName[851:874],
	_ResourceTypeName[874:912],
	_ResourceTypeName[912:947],
	_ResourceTypeName[947:987],
	_ResourceTypeName[987:1029],
	_ResourceTypeName[1029:1080],
	_ResourceTypeName[1080:1125],
	_ResourceTypeName[1125:1154],
	_ResourceTypeName[1154:1201],
	_ResourceTypeName[1201:1248],
	_ResourceTypeName[1248:1295],
	_ResourceTypeName[1295:1314],
	_ResourceTypeName[1314:1321],
	_ResourceTypeName[1321:1336],
	_ResourceTypeName[1336:1359],
	_ResourceTypeName[1359:1392],
	_ResourceTypeName[1392:1425],
	_ResourceTypeName[1425:1449],
	_ResourceTypeName[1449:1480],
	_ResourceTypeName[1480:1487],
	_ResourceTypeName[1487:1502],
	_ResourceTypeName[1502:1528],
	_ResourceTypeName[1528:1553],
	_ResourceTypeName[1553:1575],
	_ResourceTypeName[1575:1593],
	_ResourceTypeName[1593:1614],
	_ResourceTypeName[1614:1645],
	_ResourceTypeName[1645:1658],
	_ResourceTypeName[1658:1682],
	_ResourceTypeName[1682:1702],
	_ResourceTypeName[1702:1733],
	_ResourceTypeName[1733:1757],
	_ResourceTypeName[1757:1788],
	_ResourceTypeName[1788:1802],
	_ResourceTypeName[1802:1814],
	_ResourceTypeName[1814:1833],
	_ResourceTypeName[1833:1863],
	_ResourceTypeName[1863:1884],
	_ResourceTypeName[1884:1910],
	_ResourceTypeName[1910:1922],
	_ResourceTypeName[1922:1951],
	_ResourceTypeName[1951:1970],
	_ResourceTypeName[1970:2000],
	_ResourceTypeName[2000:2020],
	_ResourceTypeName[2020:2040],
	_ResourceTypeName[2040:2052],
	_ResourceTypeName[2052:2070],
	_ResourceTypeName[2070:2089],
	_ResourceTypeName[2089:2113],
	_ResourceTypeName[2113:2132],
	_ResourceTypeName[2132:2138],
	_ResourceTypeName[2138:2169],
	_ResourceTypeName[2169:2184],
	_ResourceTypeName[2184:2211],
	_ResourceTypeName[2211:2231],
	_ResourceTypeName[2231:2250],
	_ResourceTypeName[2250:2280],
	_ResourceTypeName[2280:2302],
	_ResourceTypeName[2302:2327],
	_ResourceTypeName[2327:2340],
	_ResourceTypeName[2340:2355],
	_ResourceTypeName[2355:2370],
	_ResourceTypeName[2370:2398],
	_ResourceTypeName[2398:2433],
	_ResourceTypeName[2433:2463],
	_ResourceTypeName[2463:2482],
	_ResourceTypeName[2482:2497],
	_ResourceTypeName[2497:2519],
	_ResourceTypeName[2519:2539],
	_ResourceTypeName[2539:2565],
	_ResourceTypeName[2565:2589],
	_ResourceTypeName[2589:2610],
	_ResourceTypeName[2610:2628],
	_ResourceTypeName[2628:2657],
	_ResourceTypeName[2657:2694],
	_ResourceTypeName[2694:2710],
	_ResourceTypeName[2710:2738],
	_ResourceTypeName[2738:2753],
	_ResourceTypeName[2753:2766],
	_ResourceTypeName[2766:2784],
	_ResourceTypeName[2784:2815],
	_ResourceTypeName[2815:2840],
	_ResourceTypeName[2840:2859],
	_ResourceTypeName[2859:2882],
	_ResourceTypeName[2882:2906],
	_ResourceTypeName[2906:2941],
	_ResourceTypeName[2941:2963],
	_ResourceTypeName[2963:2983],
	_ResourceTypeName[2983:3007],
	_ResourceTypeName[3007:3023],
	_ResourceTypeName[3023:3036],
	_ResourceTypeName[3036:3062],
	_ResourceTypeName[3062:3072],
	_ResourceTypeName[3072:3093],
	_ResourceTypeName[3093:3100],
	_ResourceTypeName[3100:3120],
	_ResourceTypeName[3120:3136],
	_ResourceTypeName[3136:3148],
	_ResourceTypeName[3148:3165],
	_ResourceTypeName[3165:3191],
	_ResourceTypeName[3191:3206],
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
