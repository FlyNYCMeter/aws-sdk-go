//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/codepipeline"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@codepipeline", func() {
		World["client"] = codepipeline.New(nil)
	})
}
