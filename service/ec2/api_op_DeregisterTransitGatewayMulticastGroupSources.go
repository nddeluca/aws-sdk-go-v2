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

// Deregisters the specified sources (network interfaces) from the transit gateway
// multicast group.
func (c *Client) DeregisterTransitGatewayMulticastGroupSources(ctx context.Context, params *DeregisterTransitGatewayMulticastGroupSourcesInput, optFns ...func(*Options)) (*DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	if params == nil {
		params = &DeregisterTransitGatewayMulticastGroupSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterTransitGatewayMulticastGroupSources", params, optFns, addOperationDeregisterTransitGatewayMulticastGroupSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterTransitGatewayMulticastGroupSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTransitGatewayMulticastGroupSourcesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string

	// The IDs of the group sources' network interfaces.
	NetworkInterfaceIds []string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string
}

type DeregisterTransitGatewayMulticastGroupSourcesOutput struct {

	// Information about the deregistered group sources.
	DeregisteredMulticastGroupSources *types.TransitGatewayMulticastDeregisteredGroupSources

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterTransitGatewayMulticastGroupSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeregisterTransitGatewayMulticastGroupSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterTransitGatewayMulticastGroupSources{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTransitGatewayMulticastGroupSources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterTransitGatewayMulticastGroupSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterTransitGatewayMulticastGroupSources",
	}
}
