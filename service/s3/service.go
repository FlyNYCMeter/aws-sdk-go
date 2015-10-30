// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package s3

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/aws/defaults"
	"github.com/upstartmobile/aws-sdk-go/aws/request"
	"github.com/upstartmobile/aws-sdk-go/aws/service"
	"github.com/upstartmobile/aws-sdk-go/aws/service/serviceinfo"
	"github.com/upstartmobile/aws-sdk-go/private/protocol/restxml"
	"github.com/upstartmobile/aws-sdk-go/private/signer/v4"
)

// S3 is a client for Amazon S3.
type S3 struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new S3 client.
func New(config *aws.Config) *S3 {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "s3",
			APIVersion:  "2006-03-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &S3{service}
}

// newRequest creates a new request for a S3 operation and runs any
// custom request initialization.
func (c *S3) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
