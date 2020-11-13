// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a VPC endpoint service configuration to which service consumers (AWS
// accounts, IAM users, and IAM roles) can connect. Service consumers can create an
// interface VPC endpoint to connect to your service. To create an endpoint service
// configuration, you must first create a Network Load Balancer for your service.
// For more information, see VPC Endpoint Services
// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html) in the
// Amazon Virtual Private Cloud User Guide. If you set the private DNS name, you
// must prove that you own the private DNS domain name. For more information, see
// VPC Endpoint Service Private DNS Name Verification
// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateVpcEndpointServiceConfiguration(ctx context.Context, params *CreateVpcEndpointServiceConfigurationInput, optFns ...func(*Options)) (*CreateVpcEndpointServiceConfigurationOutput, error) {
	if params == nil {
		params = &CreateVpcEndpointServiceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcEndpointServiceConfiguration", params, optFns, addOperationCreateVpcEndpointServiceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcEndpointServiceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcEndpointServiceConfigurationInput struct {

	// The Amazon Resource Names (ARNs) of one or more Network Load Balancers for your
	// service.
	//
	// This member is required.
	NetworkLoadBalancerArns []string

	// Indicates whether requests from service consumers to create an endpoint to your
	// service must be accepted. To accept a request, use AcceptVpcEndpointConnections.
	AcceptanceRequired bool

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The private DNS name to assign to the VPC endpoint service.
	PrivateDnsName *string

	// The tags to associate with the service.
	TagSpecifications []types.TagSpecification
}

type CreateVpcEndpointServiceConfigurationOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// Information about the service configuration.
	ServiceConfiguration *types.ServiceConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVpcEndpointServiceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcEndpointServiceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcEndpointServiceConfiguration{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateVpcEndpointServiceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcEndpointServiceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcEndpointServiceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpcEndpointServiceConfiguration",
	}
}
