//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/sts"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@sts", func() {
		World["client"] = sts.New(nil)
	})
}
