//Package ecs provides gucumber integration tests support.
package ecs

import (
	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/ecs"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@ecs", func() {
		// FIXME remove custom region
		World["client"] = ecs.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
