// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package storagegateway

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/aws/defaults"
	"github.com/upstartmobile/aws-sdk-go/aws/request"
	"github.com/upstartmobile/aws-sdk-go/aws/service"
	"github.com/upstartmobile/aws-sdk-go/aws/service/serviceinfo"
	"github.com/upstartmobile/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/upstartmobile/aws-sdk-go/private/signer/v4"
)

// AWS Storage Gateway is the service that connects an on-premises software
// appliance with cloud-based storage to provide seamless and secure integration
// between an organization's on-premises IT environment and AWS's storage infrastructure.
// The service enables you to securely upload data to the AWS cloud for cost
// effective backup and rapid disaster recovery.
//
// Use the following links to get started using the AWS Storage Gateway Service
// API Reference:
//
//   AWS Storage Gateway Required Request Headers (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayHTTPRequestsHeaders.html):
// Describes the required headers that you must send with every POST request
// to AWS Storage Gateway.  Signing Requests (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewaySigningRequests.html):
// AWS Storage Gateway requires that you authenticate every request you send;
// this topic describes how sign such a request.  Error Responses (http://docs.aws.amazon.com/storagegateway/latest/userguide/APIErrorResponses.html):
// Provides reference information about AWS Storage Gateway errors.  Operations
// in AWS Storage Gateway (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPIOperations.html):
// Contains detailed descriptions of all AWS Storage Gateway operations, their
// request parameters, response elements, possible errors, and examples of requests
// and responses.  AWS Storage Gateway Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/index.html?rande.html):
// Provides a list of each of the regions and endpoints available for use with
// AWS Storage Gateway.
type StorageGateway struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new StorageGateway client.
func New(config *aws.Config) *StorageGateway {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "storagegateway",
			APIVersion:   "2013-06-30",
			JSONVersion:  "1.1",
			TargetPrefix: "StorageGateway_20130630",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &StorageGateway{service}
}

// newRequest creates a new request for a StorageGateway operation and runs any
// custom request initialization.
func (c *StorageGateway) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
