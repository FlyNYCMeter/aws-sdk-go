//Package cloudtrail provides gucumber integration tests support.
package cloudtrail

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cloudtrail"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudtrail", func() {
		World["client"] = cloudtrail.New(nil)
	})
}
