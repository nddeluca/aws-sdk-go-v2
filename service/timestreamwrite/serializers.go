// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsAwsjson10_serializeOpCreateDatabase struct {
}

func (*awsAwsjson10_serializeOpCreateDatabase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateDatabase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateDatabaseInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.CreateDatabase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateDatabaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpCreateTable struct {
}

func (*awsAwsjson10_serializeOpCreateTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.CreateTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteDatabase struct {
}

func (*awsAwsjson10_serializeOpDeleteDatabase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteDatabase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteDatabaseInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.DeleteDatabase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteDatabaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteTable struct {
}

func (*awsAwsjson10_serializeOpDeleteTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.DeleteTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeDatabase struct {
}

func (*awsAwsjson10_serializeOpDescribeDatabase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeDatabase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeDatabaseInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.DescribeDatabase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeDatabaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeEndpoints struct {
}

func (*awsAwsjson10_serializeOpDescribeEndpoints) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeEndpoints) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeEndpointsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.DescribeEndpoints")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeEndpointsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeTable struct {
}

func (*awsAwsjson10_serializeOpDescribeTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.DescribeTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListDatabases struct {
}

func (*awsAwsjson10_serializeOpListDatabases) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListDatabases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListDatabasesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.ListDatabases")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListDatabasesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListTables struct {
}

func (*awsAwsjson10_serializeOpListTables) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTablesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.ListTables")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListTablesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListTagsForResource struct {
}

func (*awsAwsjson10_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpTagResource struct {
}

func (*awsAwsjson10_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUntagResource struct {
}

func (*awsAwsjson10_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateDatabase struct {
}

func (*awsAwsjson10_serializeOpUpdateDatabase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateDatabase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateDatabaseInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.UpdateDatabase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateDatabaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateTable struct {
}

func (*awsAwsjson10_serializeOpUpdateTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.UpdateTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpWriteRecords struct {
}

func (*awsAwsjson10_serializeOpWriteRecords) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpWriteRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*WriteRecordsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("Timestream_20181101.WriteRecords")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentWriteRecordsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentDimension(v *types.Dimension, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.DimensionValueType) > 0 {
		ok := object.Key("DimensionValueType")
		ok.String(string(v.DimensionValueType))
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson10_serializeDocumentDimensions(v []types.Dimension, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentDimension(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeDocumentRecord(v *types.Record, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Dimensions != nil {
		ok := object.Key("Dimensions")
		if err := awsAwsjson10_serializeDocumentDimensions(v.Dimensions, ok); err != nil {
			return err
		}
	}

	if v.MeasureName != nil {
		ok := object.Key("MeasureName")
		ok.String(*v.MeasureName)
	}

	if v.MeasureValue != nil {
		ok := object.Key("MeasureValue")
		ok.String(*v.MeasureValue)
	}

	if len(v.MeasureValueType) > 0 {
		ok := object.Key("MeasureValueType")
		ok.String(string(v.MeasureValueType))
	}

	if v.Time != nil {
		ok := object.Key("Time")
		ok.String(*v.Time)
	}

	if len(v.TimeUnit) > 0 {
		ok := object.Key("TimeUnit")
		ok.String(string(v.TimeUnit))
	}

	if v.Version != 0 {
		ok := object.Key("Version")
		ok.Long(v.Version)
	}

	return nil
}

func awsAwsjson10_serializeDocumentRecords(v []types.Record, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentRecord(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeDocumentRetentionProperties(v *types.RetentionProperties, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MagneticStoreRetentionPeriodInDays != 0 {
		ok := object.Key("MagneticStoreRetentionPeriodInDays")
		ok.Long(v.MagneticStoreRetentionPeriodInDays)
	}

	if v.MemoryStoreRetentionPeriodInHours != 0 {
		ok := object.Key("MemoryStoreRetentionPeriodInHours")
		ok.Long(v.MemoryStoreRetentionPeriodInHours)
	}

	return nil
}

func awsAwsjson10_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson10_serializeDocumentTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10_serializeDocumentTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeOpDocumentCreateDatabaseInput(v *CreateDatabaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.KmsKeyId != nil {
		ok := object.Key("KmsKeyId")
		ok.String(*v.KmsKeyId)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentCreateTableInput(v *CreateTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.RetentionProperties != nil {
		ok := object.Key("RetentionProperties")
		if err := awsAwsjson10_serializeDocumentRetentionProperties(v.RetentionProperties, ok); err != nil {
			return err
		}
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteDatabaseInput(v *DeleteDatabaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteTableInput(v *DeleteTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeDatabaseInput(v *DescribeDatabaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeEndpointsInput(v *DescribeEndpointsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeTableInput(v *DescribeTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListDatabasesInput(v *ListDatabasesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListTablesInput(v *ListTablesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsAwsjson10_serializeDocumentTagList(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceARN != nil {
		ok := object.Key("ResourceARN")
		ok.String(*v.ResourceARN)
	}

	if v.TagKeys != nil {
		ok := object.Key("TagKeys")
		if err := awsAwsjson10_serializeDocumentTagKeyList(v.TagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateDatabaseInput(v *UpdateDatabaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.KmsKeyId != nil {
		ok := object.Key("KmsKeyId")
		ok.String(*v.KmsKeyId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateTableInput(v *UpdateTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.RetentionProperties != nil {
		ok := object.Key("RetentionProperties")
		if err := awsAwsjson10_serializeDocumentRetentionProperties(v.RetentionProperties, ok); err != nil {
			return err
		}
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentWriteRecordsInput(v *WriteRecordsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CommonAttributes != nil {
		ok := object.Key("CommonAttributes")
		if err := awsAwsjson10_serializeDocumentRecord(v.CommonAttributes, ok); err != nil {
			return err
		}
	}

	if v.DatabaseName != nil {
		ok := object.Key("DatabaseName")
		ok.String(*v.DatabaseName)
	}

	if v.Records != nil {
		ok := object.Key("Records")
		if err := awsAwsjson10_serializeDocumentRecords(v.Records, ok); err != nil {
			return err
		}
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}
