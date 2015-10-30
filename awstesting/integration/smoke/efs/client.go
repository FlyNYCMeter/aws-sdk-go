//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
