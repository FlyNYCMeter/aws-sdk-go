//Package directoryservice provides gucumber integration tests support.
package directoryservice

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@directoryservice", func() {
		World["client"] = directoryservice.New(nil)
	})
}
