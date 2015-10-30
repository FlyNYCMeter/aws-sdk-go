//Package elasticloadbalancing provides gucumber integration tests support.
package elasticloadbalancing

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/elb"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@elasticloadbalancing", func() {
		World["client"] = elb.New(nil)
	})
}
