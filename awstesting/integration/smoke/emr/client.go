//Package emr provides gucumber integration tests support.
package emr

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/emr"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@emr", func() {
		World["client"] = emr.New(nil)
	})
}
