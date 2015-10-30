//Package cloudformation provides gucumber integration tests support.
package cloudformation

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cloudformation"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudformation", func() {
		World["client"] = cloudformation.New(nil)
	})
}
