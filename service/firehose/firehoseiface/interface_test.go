// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package firehoseiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/firehose"
	"github.com/upstartmobile/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*firehoseiface.FirehoseAPI)(nil), firehose.New(nil))
}
