package apigateway

import (
	"github.com/upstartmobile/aws-sdk-go/aws/request"
	"github.com/upstartmobile/aws-sdk-go/aws/service"
)

func init() {
	initService = func(s *service.Service) {
		s.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
