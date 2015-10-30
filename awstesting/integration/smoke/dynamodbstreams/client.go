//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/dynamodbstreams"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@dynamodbstreams", func() {
		World["client"] = dynamodbstreams.New(nil)
	})
}
