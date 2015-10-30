//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(nil)
	})
}
