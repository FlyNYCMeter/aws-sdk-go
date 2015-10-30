// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package redshiftiface provides an interface for the Amazon Redshift.
package redshiftiface

import (
	"github.com/upstartmobile/aws-sdk-go/aws/request"
	"github.com/upstartmobile/aws-sdk-go/service/redshift"
)

// RedshiftAPI is the interface type for redshift.Redshift.
type RedshiftAPI interface {
	AuthorizeClusterSecurityGroupIngressRequest(*redshift.AuthorizeClusterSecurityGroupIngressInput) (*request.Request, *redshift.AuthorizeClusterSecurityGroupIngressOutput)

	AuthorizeClusterSecurityGroupIngress(*redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error)

	AuthorizeSnapshotAccessRequest(*redshift.AuthorizeSnapshotAccessInput) (*request.Request, *redshift.AuthorizeSnapshotAccessOutput)

	AuthorizeSnapshotAccess(*redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error)

	CopyClusterSnapshotRequest(*redshift.CopyClusterSnapshotInput) (*request.Request, *redshift.CopyClusterSnapshotOutput)

	CopyClusterSnapshot(*redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error)

	CreateClusterRequest(*redshift.CreateClusterInput) (*request.Request, *redshift.CreateClusterOutput)

	CreateCluster(*redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error)

	CreateClusterParameterGroupRequest(*redshift.CreateClusterParameterGroupInput) (*request.Request, *redshift.CreateClusterParameterGroupOutput)

	CreateClusterParameterGroup(*redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error)

	CreateClusterSecurityGroupRequest(*redshift.CreateClusterSecurityGroupInput) (*request.Request, *redshift.CreateClusterSecurityGroupOutput)

	CreateClusterSecurityGroup(*redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error)

	CreateClusterSnapshotRequest(*redshift.CreateClusterSnapshotInput) (*request.Request, *redshift.CreateClusterSnapshotOutput)

	CreateClusterSnapshot(*redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error)

	CreateClusterSubnetGroupRequest(*redshift.CreateClusterSubnetGroupInput) (*request.Request, *redshift.CreateClusterSubnetGroupOutput)

	CreateClusterSubnetGroup(*redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error)

	CreateEventSubscriptionRequest(*redshift.CreateEventSubscriptionInput) (*request.Request, *redshift.CreateEventSubscriptionOutput)

	CreateEventSubscription(*redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error)

	CreateHsmClientCertificateRequest(*redshift.CreateHsmClientCertificateInput) (*request.Request, *redshift.CreateHsmClientCertificateOutput)

	CreateHsmClientCertificate(*redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error)

	CreateHsmConfigurationRequest(*redshift.CreateHsmConfigurationInput) (*request.Request, *redshift.CreateHsmConfigurationOutput)

	CreateHsmConfiguration(*redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error)

	CreateSnapshotCopyGrantRequest(*redshift.CreateSnapshotCopyGrantInput) (*request.Request, *redshift.CreateSnapshotCopyGrantOutput)

	CreateSnapshotCopyGrant(*redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error)

	CreateTagsRequest(*redshift.CreateTagsInput) (*request.Request, *redshift.CreateTagsOutput)

	CreateTags(*redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error)

	DeleteClusterRequest(*redshift.DeleteClusterInput) (*request.Request, *redshift.DeleteClusterOutput)

	DeleteCluster(*redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error)

	DeleteClusterParameterGroupRequest(*redshift.DeleteClusterParameterGroupInput) (*request.Request, *redshift.DeleteClusterParameterGroupOutput)

	DeleteClusterParameterGroup(*redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error)

	DeleteClusterSecurityGroupRequest(*redshift.DeleteClusterSecurityGroupInput) (*request.Request, *redshift.DeleteClusterSecurityGroupOutput)

	DeleteClusterSecurityGroup(*redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error)

	DeleteClusterSnapshotRequest(*redshift.DeleteClusterSnapshotInput) (*request.Request, *redshift.DeleteClusterSnapshotOutput)

	DeleteClusterSnapshot(*redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error)

	DeleteClusterSubnetGroupRequest(*redshift.DeleteClusterSubnetGroupInput) (*request.Request, *redshift.DeleteClusterSubnetGroupOutput)

	DeleteClusterSubnetGroup(*redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error)

	DeleteEventSubscriptionRequest(*redshift.DeleteEventSubscriptionInput) (*request.Request, *redshift.DeleteEventSubscriptionOutput)

	DeleteEventSubscription(*redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error)

	DeleteHsmClientCertificateRequest(*redshift.DeleteHsmClientCertificateInput) (*request.Request, *redshift.DeleteHsmClientCertificateOutput)

	DeleteHsmClientCertificate(*redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error)

	DeleteHsmConfigurationRequest(*redshift.DeleteHsmConfigurationInput) (*request.Request, *redshift.DeleteHsmConfigurationOutput)

	DeleteHsmConfiguration(*redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error)

	DeleteSnapshotCopyGrantRequest(*redshift.DeleteSnapshotCopyGrantInput) (*request.Request, *redshift.DeleteSnapshotCopyGrantOutput)

	DeleteSnapshotCopyGrant(*redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error)

	DeleteTagsRequest(*redshift.DeleteTagsInput) (*request.Request, *redshift.DeleteTagsOutput)

	DeleteTags(*redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error)

	DescribeClusterParameterGroupsRequest(*redshift.DescribeClusterParameterGroupsInput) (*request.Request, *redshift.DescribeClusterParameterGroupsOutput)

	DescribeClusterParameterGroups(*redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error)

	DescribeClusterParameterGroupsPages(*redshift.DescribeClusterParameterGroupsInput, func(*redshift.DescribeClusterParameterGroupsOutput, bool) bool) error

	DescribeClusterParametersRequest(*redshift.DescribeClusterParametersInput) (*request.Request, *redshift.DescribeClusterParametersOutput)

	DescribeClusterParameters(*redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error)

	DescribeClusterParametersPages(*redshift.DescribeClusterParametersInput, func(*redshift.DescribeClusterParametersOutput, bool) bool) error

	DescribeClusterSecurityGroupsRequest(*redshift.DescribeClusterSecurityGroupsInput) (*request.Request, *redshift.DescribeClusterSecurityGroupsOutput)

	DescribeClusterSecurityGroups(*redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error)

	DescribeClusterSecurityGroupsPages(*redshift.DescribeClusterSecurityGroupsInput, func(*redshift.DescribeClusterSecurityGroupsOutput, bool) bool) error

	DescribeClusterSnapshotsRequest(*redshift.DescribeClusterSnapshotsInput) (*request.Request, *redshift.DescribeClusterSnapshotsOutput)

	DescribeClusterSnapshots(*redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error)

	DescribeClusterSnapshotsPages(*redshift.DescribeClusterSnapshotsInput, func(*redshift.DescribeClusterSnapshotsOutput, bool) bool) error

	DescribeClusterSubnetGroupsRequest(*redshift.DescribeClusterSubnetGroupsInput) (*request.Request, *redshift.DescribeClusterSubnetGroupsOutput)

	DescribeClusterSubnetGroups(*redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error)

	DescribeClusterSubnetGroupsPages(*redshift.DescribeClusterSubnetGroupsInput, func(*redshift.DescribeClusterSubnetGroupsOutput, bool) bool) error

	DescribeClusterVersionsRequest(*redshift.DescribeClusterVersionsInput) (*request.Request, *redshift.DescribeClusterVersionsOutput)

	DescribeClusterVersions(*redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error)

	DescribeClusterVersionsPages(*redshift.DescribeClusterVersionsInput, func(*redshift.DescribeClusterVersionsOutput, bool) bool) error

	DescribeClustersRequest(*redshift.DescribeClustersInput) (*request.Request, *redshift.DescribeClustersOutput)

	DescribeClusters(*redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error)

	DescribeClustersPages(*redshift.DescribeClustersInput, func(*redshift.DescribeClustersOutput, bool) bool) error

	DescribeDefaultClusterParametersRequest(*redshift.DescribeDefaultClusterParametersInput) (*request.Request, *redshift.DescribeDefaultClusterParametersOutput)

	DescribeDefaultClusterParameters(*redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error)

	DescribeDefaultClusterParametersPages(*redshift.DescribeDefaultClusterParametersInput, func(*redshift.DescribeDefaultClusterParametersOutput, bool) bool) error

	DescribeEventCategoriesRequest(*redshift.DescribeEventCategoriesInput) (*request.Request, *redshift.DescribeEventCategoriesOutput)

	DescribeEventCategories(*redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error)

	DescribeEventSubscriptionsRequest(*redshift.DescribeEventSubscriptionsInput) (*request.Request, *redshift.DescribeEventSubscriptionsOutput)

	DescribeEventSubscriptions(*redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error)

	DescribeEventSubscriptionsPages(*redshift.DescribeEventSubscriptionsInput, func(*redshift.DescribeEventSubscriptionsOutput, bool) bool) error

	DescribeEventsRequest(*redshift.DescribeEventsInput) (*request.Request, *redshift.DescribeEventsOutput)

	DescribeEvents(*redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error)

	DescribeEventsPages(*redshift.DescribeEventsInput, func(*redshift.DescribeEventsOutput, bool) bool) error

	DescribeHsmClientCertificatesRequest(*redshift.DescribeHsmClientCertificatesInput) (*request.Request, *redshift.DescribeHsmClientCertificatesOutput)

	DescribeHsmClientCertificates(*redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error)

	DescribeHsmClientCertificatesPages(*redshift.DescribeHsmClientCertificatesInput, func(*redshift.DescribeHsmClientCertificatesOutput, bool) bool) error

	DescribeHsmConfigurationsRequest(*redshift.DescribeHsmConfigurationsInput) (*request.Request, *redshift.DescribeHsmConfigurationsOutput)

	DescribeHsmConfigurations(*redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error)

	DescribeHsmConfigurationsPages(*redshift.DescribeHsmConfigurationsInput, func(*redshift.DescribeHsmConfigurationsOutput, bool) bool) error

	DescribeLoggingStatusRequest(*redshift.DescribeLoggingStatusInput) (*request.Request, *redshift.LoggingStatus)

	DescribeLoggingStatus(*redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error)

	DescribeOrderableClusterOptionsRequest(*redshift.DescribeOrderableClusterOptionsInput) (*request.Request, *redshift.DescribeOrderableClusterOptionsOutput)

	DescribeOrderableClusterOptions(*redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error)

	DescribeOrderableClusterOptionsPages(*redshift.DescribeOrderableClusterOptionsInput, func(*redshift.DescribeOrderableClusterOptionsOutput, bool) bool) error

	DescribeReservedNodeOfferingsRequest(*redshift.DescribeReservedNodeOfferingsInput) (*request.Request, *redshift.DescribeReservedNodeOfferingsOutput)

	DescribeReservedNodeOfferings(*redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error)

	DescribeReservedNodeOfferingsPages(*redshift.DescribeReservedNodeOfferingsInput, func(*redshift.DescribeReservedNodeOfferingsOutput, bool) bool) error

	DescribeReservedNodesRequest(*redshift.DescribeReservedNodesInput) (*request.Request, *redshift.DescribeReservedNodesOutput)

	DescribeReservedNodes(*redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error)

	DescribeReservedNodesPages(*redshift.DescribeReservedNodesInput, func(*redshift.DescribeReservedNodesOutput, bool) bool) error

	DescribeResizeRequest(*redshift.DescribeResizeInput) (*request.Request, *redshift.DescribeResizeOutput)

	DescribeResize(*redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error)

	DescribeSnapshotCopyGrantsRequest(*redshift.DescribeSnapshotCopyGrantsInput) (*request.Request, *redshift.DescribeSnapshotCopyGrantsOutput)

	DescribeSnapshotCopyGrants(*redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error)

	DescribeTagsRequest(*redshift.DescribeTagsInput) (*request.Request, *redshift.DescribeTagsOutput)

	DescribeTags(*redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error)

	DisableLoggingRequest(*redshift.DisableLoggingInput) (*request.Request, *redshift.LoggingStatus)

	DisableLogging(*redshift.DisableLoggingInput) (*redshift.LoggingStatus, error)

	DisableSnapshotCopyRequest(*redshift.DisableSnapshotCopyInput) (*request.Request, *redshift.DisableSnapshotCopyOutput)

	DisableSnapshotCopy(*redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error)

	EnableLoggingRequest(*redshift.EnableLoggingInput) (*request.Request, *redshift.LoggingStatus)

	EnableLogging(*redshift.EnableLoggingInput) (*redshift.LoggingStatus, error)

	EnableSnapshotCopyRequest(*redshift.EnableSnapshotCopyInput) (*request.Request, *redshift.EnableSnapshotCopyOutput)

	EnableSnapshotCopy(*redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error)

	ModifyClusterRequest(*redshift.ModifyClusterInput) (*request.Request, *redshift.ModifyClusterOutput)

	ModifyCluster(*redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error)

	ModifyClusterParameterGroupRequest(*redshift.ModifyClusterParameterGroupInput) (*request.Request, *redshift.ClusterParameterGroupNameMessage)

	ModifyClusterParameterGroup(*redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)

	ModifyClusterSubnetGroupRequest(*redshift.ModifyClusterSubnetGroupInput) (*request.Request, *redshift.ModifyClusterSubnetGroupOutput)

	ModifyClusterSubnetGroup(*redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error)

	ModifyEventSubscriptionRequest(*redshift.ModifyEventSubscriptionInput) (*request.Request, *redshift.ModifyEventSubscriptionOutput)

	ModifyEventSubscription(*redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error)

	ModifySnapshotCopyRetentionPeriodRequest(*redshift.ModifySnapshotCopyRetentionPeriodInput) (*request.Request, *redshift.ModifySnapshotCopyRetentionPeriodOutput)

	ModifySnapshotCopyRetentionPeriod(*redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error)

	PurchaseReservedNodeOfferingRequest(*redshift.PurchaseReservedNodeOfferingInput) (*request.Request, *redshift.PurchaseReservedNodeOfferingOutput)

	PurchaseReservedNodeOffering(*redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error)

	RebootClusterRequest(*redshift.RebootClusterInput) (*request.Request, *redshift.RebootClusterOutput)

	RebootCluster(*redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error)

	ResetClusterParameterGroupRequest(*redshift.ResetClusterParameterGroupInput) (*request.Request, *redshift.ClusterParameterGroupNameMessage)

	ResetClusterParameterGroup(*redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)

	RestoreFromClusterSnapshotRequest(*redshift.RestoreFromClusterSnapshotInput) (*request.Request, *redshift.RestoreFromClusterSnapshotOutput)

	RestoreFromClusterSnapshot(*redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error)

	RevokeClusterSecurityGroupIngressRequest(*redshift.RevokeClusterSecurityGroupIngressInput) (*request.Request, *redshift.RevokeClusterSecurityGroupIngressOutput)

	RevokeClusterSecurityGroupIngress(*redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error)

	RevokeSnapshotAccessRequest(*redshift.RevokeSnapshotAccessInput) (*request.Request, *redshift.RevokeSnapshotAccessOutput)

	RevokeSnapshotAccess(*redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error)

	RotateEncryptionKeyRequest(*redshift.RotateEncryptionKeyInput) (*request.Request, *redshift.RotateEncryptionKeyOutput)

	RotateEncryptionKey(*redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error)
}
