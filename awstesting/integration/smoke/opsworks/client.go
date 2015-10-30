//Package opsworks provides gucumber integration tests support.
package opsworks

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/opsworks"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@opsworks", func() {
		World["client"] = opsworks.New(nil)
	})
}
