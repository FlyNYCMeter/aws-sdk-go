// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package simpledbiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/simpledb"
	"github.com/upstartmobile/aws-sdk-go/service/simpledb/simpledbiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*simpledbiface.SimpleDBAPI)(nil), simpledb.New(nil))
}
