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

// Describes the associations between virtual interface groups and local gateway
// route tables.
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx context.Context, params *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations", params, optFns, addOperationDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * local-gateway-id - The ID of a local gateway.
	//
	// *
	// local-gateway-route-table-id - The ID of the local gateway route table.
	//
	// *
	// local-gateway-route-table-virtual-interface-group-association-id - The ID of the
	// association.
	//
	// * local-gateway-route-table-virtual-interface-group-id - The ID of
	// the virtual interface group.
	//
	// * state - The state of the association.
	Filters []types.Filter

	// The IDs of the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput struct {

	// Information about the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociations []types.LocalGatewayRouteTableVirtualInterfaceGroupAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
	}
}
