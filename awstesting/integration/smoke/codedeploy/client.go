//Package codedeploy provides gucumber integration tests support.
package codedeploy

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/codedeploy"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@codedeploy", func() {
		World["client"] = codedeploy.New(nil)
	})
}
