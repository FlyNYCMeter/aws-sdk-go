// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package waf

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/aws/defaults"
	"github.com/upstartmobile/aws-sdk-go/aws/request"
	"github.com/upstartmobile/aws-sdk-go/aws/service"
	"github.com/upstartmobile/aws-sdk-go/aws/service/serviceinfo"
	"github.com/upstartmobile/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/upstartmobile/aws-sdk-go/private/signer/v4"
)

// This is the AWS WAF API Reference. This guide is for developers who need
// detailed information about the AWS WAF API actions, data types, and errors.
// For detailed information about AWS WAF features and an overview of how to
// use the AWS WAF API, see the AWS WAF Developer Guide (http://docs.aws.amazon.com/waf/latest/dev/).
type WAF struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new WAF client.
func New(config *aws.Config) *WAF {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "waf",
			APIVersion:   "2015-08-24",
			JSONVersion:  "1.1",
			TargetPrefix: "AWSWAF_20150824",
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

	return &WAF{service}
}

// newRequest creates a new request for a WAF operation and runs any
// custom request initialization.
func (c *WAF) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
