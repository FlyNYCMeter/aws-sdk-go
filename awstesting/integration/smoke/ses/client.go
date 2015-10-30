//Package ses provides gucumber integration tests support.
package ses

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/ses"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@ses", func() {
		World["client"] = ses.New(nil)
	})
}
