// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directoryserviceiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/directoryservice"
	"github.com/upstartmobile/aws-sdk-go/service/directoryservice/directoryserviceiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*directoryserviceiface.DirectoryServiceAPI)(nil), directoryservice.New(nil))
}
