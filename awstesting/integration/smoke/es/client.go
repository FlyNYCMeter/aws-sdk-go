//Package es provides gucumber integration tests support.
package es

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/elasticsearchservice"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@es", func() {
		World["client"] = elasticsearchservice.New(nil)
	})
}
