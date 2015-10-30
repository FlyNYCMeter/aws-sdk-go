//Package cloudsearch provides gucumber integration tests support.
package cloudsearch

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cloudsearch"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cloudsearch", func() {
		World["client"] = cloudsearch.New(nil)
	})
}
