// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/cloudwatch"
	"github.com/upstartmobile/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudwatchiface.CloudWatchAPI)(nil), cloudwatch.New(nil))
}
