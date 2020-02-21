// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type GetBucketAnalyticsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket from which an analytics configuration is retrieved.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID that identifies the analytics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketAnalyticsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketAnalyticsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketAnalyticsConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketAnalyticsConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAnalyticsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketAnalyticsConfigurationInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketAnalyticsConfigurationInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketAnalyticsConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"AnalyticsConfiguration"`

	// The configuration and any analyses for the analytics filter.
	AnalyticsConfiguration *AnalyticsConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketAnalyticsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketAnalyticsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AnalyticsConfiguration != nil {
		v := s.AnalyticsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "AnalyticsConfiguration", v, metadata)
	}
	return nil
}

const opGetBucketAnalyticsConfiguration = "GetBucketAnalyticsConfiguration"

// GetBucketAnalyticsConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This implementation of the GET operation returns an analytics configuration
// (identified by the analytics configuration ID) from the bucket.
//
// To use this operation, you must have permissions to perform the s3:GetAnalyticsConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// For information about Amazon S3 analytics feature, see Amazon S3 Analytics
// – Storage Class Analysis (https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Related Resources
//
//    *
//
//    *
//
//    *
//
//    // Example sending a request using GetBucketAnalyticsConfigurationRequest.
//    req := client.GetBucketAnalyticsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAnalyticsConfiguration
func (c *Client) GetBucketAnalyticsConfigurationRequest(input *GetBucketAnalyticsConfigurationInput) GetBucketAnalyticsConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketAnalyticsConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?analytics",
	}

	if input == nil {
		input = &GetBucketAnalyticsConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketAnalyticsConfigurationOutput{})
	return GetBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketAnalyticsConfigurationRequest}
}

// GetBucketAnalyticsConfigurationRequest is the request type for the
// GetBucketAnalyticsConfiguration API operation.
type GetBucketAnalyticsConfigurationRequest struct {
	*aws.Request
	Input *GetBucketAnalyticsConfigurationInput
	Copy  func(*GetBucketAnalyticsConfigurationInput) GetBucketAnalyticsConfigurationRequest
}

// Send marshals and sends the GetBucketAnalyticsConfiguration API request.
func (r GetBucketAnalyticsConfigurationRequest) Send(ctx context.Context) (*GetBucketAnalyticsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketAnalyticsConfigurationResponse{
		GetBucketAnalyticsConfigurationOutput: r.Request.Data.(*GetBucketAnalyticsConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketAnalyticsConfigurationResponse is the response type for the
// GetBucketAnalyticsConfiguration API operation.
type GetBucketAnalyticsConfigurationResponse struct {
	*GetBucketAnalyticsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketAnalyticsConfiguration request.
func (r *GetBucketAnalyticsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}