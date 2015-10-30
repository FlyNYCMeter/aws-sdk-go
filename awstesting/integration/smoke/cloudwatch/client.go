//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cloudwatch"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudwatch", func() {
		World["client"] = cloudwatch.New(nil)
	})
}
