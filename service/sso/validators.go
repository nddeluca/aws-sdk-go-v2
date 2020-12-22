// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetRoleCredentials struct {
}

func (*validateOpGetRoleCredentials) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRoleCredentials) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRoleCredentialsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRoleCredentialsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAccountRoles struct {
}

func (*validateOpListAccountRoles) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAccountRoles) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAccountRolesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAccountRolesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAccounts struct {
}

func (*validateOpListAccounts) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAccounts) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAccountsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAccountsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpLogout struct {
}

func (*validateOpLogout) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpLogout) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*LogoutInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpLogoutInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetRoleCredentialsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRoleCredentials{}, middleware.After)
}

func addOpListAccountRolesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAccountRoles{}, middleware.After)
}

func addOpListAccountsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAccounts{}, middleware.After)
}

func addOpLogoutValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpLogout{}, middleware.After)
}

func validateOpGetRoleCredentialsInput(v *GetRoleCredentialsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRoleCredentialsInput"}
	if v.RoleName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleName"))
	}
	if v.AccessToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessToken"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAccountRolesInput(v *ListAccountRolesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAccountRolesInput"}
	if v.AccessToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessToken"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAccountsInput(v *ListAccountsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAccountsInput"}
	if v.AccessToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpLogoutInput(v *LogoutInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LogoutInput"}
	if v.AccessToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}