// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The value associated with a feature.
type FeatureValue struct {

	// The name of a feature that a feature value corresponds to.
	//
	// This member is required.
	FeatureName *string

	// The value associated with a feature, in string format. Note that features types
	// can be String, Integral, or Fractional. This value represents all three types as
	// a string.
	//
	// This member is required.
	ValueAsString *string
}