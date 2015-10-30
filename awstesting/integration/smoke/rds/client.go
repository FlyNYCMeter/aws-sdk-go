//Package rds provides gucumber integration tests support.
package rds

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/rds"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@rds", func() {
		World["client"] = rds.New(nil)
	})
}
