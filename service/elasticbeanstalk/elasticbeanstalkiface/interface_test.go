// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticbeanstalkiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/elasticbeanstalk"
	"github.com/upstartmobile/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*elasticbeanstalkiface.ElasticBeanstalkAPI)(nil), elasticbeanstalk.New(nil))
}
