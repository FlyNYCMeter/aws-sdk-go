// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearningiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/machinelearning"
	"github.com/upstartmobile/aws-sdk-go/service/machinelearning/machinelearningiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*machinelearningiface.MachineLearningAPI)(nil), machinelearning.New(nil))
}
