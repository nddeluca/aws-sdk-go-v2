// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a key group that you can use with CloudFront signed URLs and signed
// cookies
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html).
// To create a key group, you must specify at least one public key for the key
// group. After you create a key group, you can reference it from one or more cache
// behaviors. When you reference a key group in a cache behavior, CloudFront
// requires signed URLs or signed cookies for all requests that match the cache
// behavior. The URLs or cookies must be signed with a private key whose
// corresponding public key is in the key group. The signed URL or cookie contains
// information about which public key CloudFront should use to verify the
// signature. For more information, see Serving private content
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateKeyGroup(ctx context.Context, params *CreateKeyGroupInput, optFns ...func(*Options)) (*CreateKeyGroupOutput, error) {
	if params == nil {
		params = &CreateKeyGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKeyGroup", params, optFns, addOperationCreateKeyGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyGroupInput struct {

	// A key group configuration.
	//
	// This member is required.
	KeyGroupConfig *types.KeyGroupConfig
}

type CreateKeyGroupOutput struct {

	// The identifier for this version of the key group.
	ETag *string

	// The key group that was just created.
	KeyGroup *types.KeyGroup

	// The URL of the key group.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateKeyGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateKeyGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKeyGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateKeyGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateKeyGroup",
	}
}