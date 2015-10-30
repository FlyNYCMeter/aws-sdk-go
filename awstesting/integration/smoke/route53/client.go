//Package route53 provides gucumber integration tests support.
package route53

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/route53"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@route53", func() {
		World["client"] = route53.New(nil)
	})
}
