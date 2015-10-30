//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cloudwatchlogs"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudwatchlogs", func() {
		World["client"] = cloudwatchlogs.New(nil)
	})
}
