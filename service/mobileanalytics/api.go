// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mobileanalytics provides a client for Amazon Mobile Analytics.
package mobileanalytics

import (
	"github.com/upstartmobile/aws-sdk-go/aws/awsutil"
	"github.com/upstartmobile/aws-sdk-go/aws/request"
)

const opPutEvents = "PutEvents"

// PutEventsRequest generates a request for the PutEvents operation.
func (c *MobileAnalytics) PutEventsRequest(input *PutEventsInput) (req *request.Request, output *PutEventsOutput) {
	op := &request.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-06-05/events",
	}

	if input == nil {
		input = &PutEventsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &PutEventsOutput{}
	req.Data = output
	return
}

// The PutEvents operation records one or more events. You can have up to 1,500
// unique custom events per app, any combination of up to 40 attributes and
// metrics per custom event, and any number of attribute or metric values.
func (c *MobileAnalytics) PutEvents(input *PutEventsInput) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	err := req.Send()
	return out, err
}

// A JSON object representing a batch of unique event occurrences in your app.
type Event struct {
	// A collection of key-value pairs that give additional context to the event.
	// The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Attributes map[string]*string `locationName:"attributes" type:"map"`

	// A name signifying an event that occurred in your app. This is used for grouping
	// and aggregating like events together for reporting purposes.
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// A collection of key-value pairs that gives additional, measurable context
	// to the event. The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Metrics map[string]*float64 `locationName:"metrics" type:"map"`

	// The session the event occured within.
	Session *Session `locationName:"session" type:"structure"`

	// The time the event occurred in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	Timestamp *string `locationName:"timestamp" type:"string" required:"true"`

	// The version of the event.
	Version *string `locationName:"version" min:"1" type:"string"`

	metadataEvent `json:"-" xml:"-"`
}

type metadataEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Event) GoString() string {
	return s.String()
}

// A container for the data needed for a PutEvent operation
type PutEventsInput struct {
	// The client context including the client ID, app title, app version and package
	// name.
	ClientContext *string `location:"header" locationName:"x-amz-Client-Context" type:"string" required:"true"`

	// The encoding used for the client context.
	ClientContextEncoding *string `location:"header" locationName:"x-amz-Client-Context-Encoding" type:"string"`

	// An array of Event JSON objects
	Events []*Event `locationName:"events" type:"list" required:"true"`

	metadataPutEventsInput `json:"-" xml:"-"`
}

type metadataPutEventsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsInput) GoString() string {
	return s.String()
}

type PutEventsOutput struct {
	metadataPutEventsOutput `json:"-" xml:"-"`
}

type metadataPutEventsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsOutput) GoString() string {
	return s.String()
}

// Describes the session. Session information is required on ALL events.
type Session struct {
	// The duration of the session.
	Duration *int64 `locationName:"duration" type:"long"`

	// A unique identifier for the session
	Id *string `locationName:"id" min:"1" type:"string"`

	// The time the event started in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	StartTimestamp *string `locationName:"startTimestamp" type:"string"`

	// The time the event terminated in ISO 8601 standard date time format. For
	// example, 2014-06-30T19:07:47.885Z
	StopTimestamp *string `locationName:"stopTimestamp" type:"string"`

	metadataSession `json:"-" xml:"-"`
}

type metadataSession struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Session) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Session) GoString() string {
	return s.String()
}
