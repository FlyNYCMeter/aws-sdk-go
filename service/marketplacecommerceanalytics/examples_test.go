// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package marketplacecommerceanalytics_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/service/marketplacecommerceanalytics"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMarketplaceCommerceAnalytics_GenerateDataSet() {
	svc := marketplacecommerceanalytics.New(nil)

	params := &marketplacecommerceanalytics.GenerateDataSetInput{
		DataSetPublicationDate:  aws.Time(time.Now()),                  // Required
		DataSetType:             aws.String("DataSetType"),             // Required
		DestinationS3BucketName: aws.String("DestinationS3BucketName"), // Required
		RoleNameArn:             aws.String("RoleNameArn"),             // Required
		SnsTopicArn:             aws.String("SnsTopicArn"),             // Required
		DestinationS3Prefix:     aws.String("DestinationS3Prefix"),
	}
	resp, err := svc.GenerateDataSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
