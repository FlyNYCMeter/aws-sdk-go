package sts_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/service/sts"
	"github.com/stretchr/testify/assert"
)

var svc = sts.New(&aws.Config{
	Region: aws.String("mock-region"),
})

func TestUnsignedRequest_AssumeRoleWithSAML(t *testing.T) {
	req, _ := svc.AssumeRoleWithSAMLRequest(&sts.AssumeRoleWithSAMLInput{
		PrincipalArn:  aws.String("ARN01234567890123456789"),
		RoleArn:       aws.String("ARN01234567890123456789"),
		SAMLAssertion: aws.String("ASSERT"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_AssumeRoleWithWebIdentity(t *testing.T) {
	req, _ := svc.AssumeRoleWithWebIdentityRequest(&sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          aws.String("ARN01234567890123456789"),
		RoleSessionName:  aws.String("SESSION"),
		WebIdentityToken: aws.String("TOKEN"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
