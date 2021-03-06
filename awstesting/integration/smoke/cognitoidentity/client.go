//Package cognitoidentity provides gucumber integration tests support.
package cognitoidentity

import (
	"github.com/upstartmobile/aws-sdk-go/awstesting/integration/smoke"
	"github.com/upstartmobile/aws-sdk-go/service/cognitoidentity"
	. "github.com/lsegal/gucumber"
)

var _ = smoke.Imported

func init() {
	Before("@cognitoidentity", func() {
		World["client"] = cognitoidentity.New(nil)
	})
}
