// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// MCOObjectReferenceApplyConfiguration represents a declarative configuration of the MCOObjectReference type for use
// with apply.
type MCOObjectReferenceApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// MCOObjectReferenceApplyConfiguration constructs a declarative configuration of the MCOObjectReference type for use with
// apply.
func MCOObjectReference() *MCOObjectReferenceApplyConfiguration {
	return &MCOObjectReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MCOObjectReferenceApplyConfiguration) WithName(value string) *MCOObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}
