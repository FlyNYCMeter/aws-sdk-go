//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(nil)
	})
}
