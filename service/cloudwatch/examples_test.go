// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatch_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/upstartmobile/aws-sdk-go/aws"
	"github.com/upstartmobile/aws-sdk-go/service/cloudwatch"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudWatch_DeleteAlarms() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.DeleteAlarmsInput{
		AlarmNames: []*string{ // Required
			aws.String("AlarmName"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteAlarms(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_DescribeAlarmHistory() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.DescribeAlarmHistoryInput{
		AlarmName:       aws.String("AlarmName"),
		EndDate:         aws.Time(time.Now()),
		HistoryItemType: aws.String("HistoryItemType"),
		MaxRecords:      aws.Int64(1),
		NextToken:       aws.String("NextToken"),
		StartDate:       aws.Time(time.Now()),
	}
	resp, err := svc.DescribeAlarmHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_DescribeAlarms() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.DescribeAlarmsInput{
		ActionPrefix:    aws.String("ActionPrefix"),
		AlarmNamePrefix: aws.String("AlarmNamePrefix"),
		AlarmNames: []*string{
			aws.String("AlarmName"), // Required
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		StateValue: aws.String("StateValue"),
	}
	resp, err := svc.DescribeAlarms(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_DescribeAlarmsForMetric() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.DescribeAlarmsForMetricInput{
		MetricName: aws.String("MetricName"), // Required
		Namespace:  aws.String("Namespace"),  // Required
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String("DimensionName"),  // Required
				Value: aws.String("DimensionValue"), // Required
			},
			// More values...
		},
		Period:    aws.Int64(1),
		Statistic: aws.String("Statistic"),
		Unit:      aws.String("StandardUnit"),
	}
	resp, err := svc.DescribeAlarmsForMetric(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_DisableAlarmActions() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.DisableAlarmActionsInput{
		AlarmNames: []*string{ // Required
			aws.String("AlarmName"), // Required
			// More values...
		},
	}
	resp, err := svc.DisableAlarmActions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_EnableAlarmActions() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.EnableAlarmActionsInput{
		AlarmNames: []*string{ // Required
			aws.String("AlarmName"), // Required
			// More values...
		},
	}
	resp, err := svc.EnableAlarmActions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_GetMetricStatistics() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(time.Now()),     // Required
		MetricName: aws.String("MetricName"), // Required
		Namespace:  aws.String("Namespace"),  // Required
		Period:     aws.Int64(1),             // Required
		StartTime:  aws.Time(time.Now()),     // Required
		Statistics: []*string{ // Required
			aws.String("Statistic"), // Required
			// More values...
		},
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String("DimensionName"),  // Required
				Value: aws.String("DimensionValue"), // Required
			},
			// More values...
		},
		Unit: aws.String("StandardUnit"),
	}
	resp, err := svc.GetMetricStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_ListMetrics() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.ListMetricsInput{
		Dimensions: []*cloudwatch.DimensionFilter{
			{ // Required
				Name:  aws.String("DimensionName"), // Required
				Value: aws.String("DimensionValue"),
			},
			// More values...
		},
		MetricName: aws.String("MetricName"),
		Namespace:  aws.String("Namespace"),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListMetrics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_PutMetricAlarm() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.PutMetricAlarmInput{
		AlarmName:          aws.String("AlarmName"),          // Required
		ComparisonOperator: aws.String("ComparisonOperator"), // Required
		EvaluationPeriods:  aws.Int64(1),                     // Required
		MetricName:         aws.String("MetricName"),         // Required
		Namespace:          aws.String("Namespace"),          // Required
		Period:             aws.Int64(1),                     // Required
		Statistic:          aws.String("Statistic"),          // Required
		Threshold:          aws.Float64(1.0),                 // Required
		ActionsEnabled:     aws.Bool(true),
		AlarmActions: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		AlarmDescription: aws.String("AlarmDescription"),
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String("DimensionName"),  // Required
				Value: aws.String("DimensionValue"), // Required
			},
			// More values...
		},
		InsufficientDataActions: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		OKActions: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		Unit: aws.String("StandardUnit"),
	}
	resp, err := svc.PutMetricAlarm(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_PutMetricData() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{ // Required
			{ // Required
				MetricName: aws.String("MetricName"), // Required
				Dimensions: []*cloudwatch.Dimension{
					{ // Required
						Name:  aws.String("DimensionName"),  // Required
						Value: aws.String("DimensionValue"), // Required
					},
					// More values...
				},
				StatisticValues: &cloudwatch.StatisticSet{
					Maximum:     aws.Float64(1.0), // Required
					Minimum:     aws.Float64(1.0), // Required
					SampleCount: aws.Float64(1.0), // Required
					Sum:         aws.Float64(1.0), // Required
				},
				Timestamp: aws.Time(time.Now()),
				Unit:      aws.String("StandardUnit"),
				Value:     aws.Float64(1.0),
			},
			// More values...
		},
		Namespace: aws.String("Namespace"), // Required
	}
	resp, err := svc.PutMetricData(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatch_SetAlarmState() {
	svc := cloudwatch.New(nil)

	params := &cloudwatch.SetAlarmStateInput{
		AlarmName:       aws.String("AlarmName"),   // Required
		StateReason:     aws.String("StateReason"), // Required
		StateValue:      aws.String("StateValue"),  // Required
		StateReasonData: aws.String("StateReasonData"),
	}
	resp, err := svc.SetAlarmState(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
