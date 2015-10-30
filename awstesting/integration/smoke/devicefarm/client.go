//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/devicefarm"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@devicefarm", func() {
		// FIXME remove custom region
		World["client"] = devicefarm.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
