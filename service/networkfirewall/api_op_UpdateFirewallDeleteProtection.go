// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the flag, DeleteProtection, which indicates whether it is possible to
// delete the firewall. If the flag is set to TRUE, the firewall is protected
// against deletion. This setting helps protect against accidentally deleting a
// firewall that's in use.
func (c *Client) UpdateFirewallDeleteProtection(ctx context.Context, params *UpdateFirewallDeleteProtectionInput, optFns ...func(*Options)) (*UpdateFirewallDeleteProtectionOutput, error) {
	if params == nil {
		params = &UpdateFirewallDeleteProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallDeleteProtection", params, optFns, addOperationUpdateFirewallDeleteProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallDeleteProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallDeleteProtectionInput struct {

	// A flag indicating whether it is possible to delete the firewall. A setting of
	// TRUE indicates that the firewall is protected against deletion. Use this setting
	// to protect against accidentally deleting a firewall that is in use. When you
	// create a firewall, the operation initializes this flag to TRUE.
	//
	// This member is required.
	DeleteProtection bool

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string
}

type UpdateFirewallDeleteProtectionOutput struct {

	//
	DeleteProtection bool

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFirewallDeleteProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateFirewallDeleteProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateFirewallDeleteProtection{}, middleware.After)
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
	if err = addOpUpdateFirewallDeleteProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallDeleteProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallDeleteProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateFirewallDeleteProtection",
	}
}