//Package datapipeline provides gucumber integration tests support.
package datapipeline

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/datapipeline"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@datapipeline", func() {
		World["client"] = datapipeline.New(nil)
	})
}
