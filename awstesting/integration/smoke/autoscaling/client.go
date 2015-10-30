//Package autoscaling provides gucumber integration tests support.
package autoscaling

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/autoscaling"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@autoscaling", func() {
		World["client"] = autoscaling.New(nil)
	})
}
