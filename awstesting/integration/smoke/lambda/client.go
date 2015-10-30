//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/lambda"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@lambda", func() {
		World["client"] = lambda.New(nil)
	})
}
