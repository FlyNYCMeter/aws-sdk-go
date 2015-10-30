package iotdataplane

import (
	//	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/aws/service"
)

func init() {
	initService = func(svc *service.Service) {
		//		// AWS IoT Dataplane requires a custom endpoint to be provided to the
		//		// to the service client when it is created. It is an error if no endpoint
		//		// is provided.
		//		if aws.StringValue(svc.Config.Endpoint) == "" {
		//			svc.Endpoint = ""
		//		}
	}
}
