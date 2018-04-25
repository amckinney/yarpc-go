// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package extends

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Foo_Foo_Args represents the arguments for the Foo.foo function.
//
// The arguments for foo are sent and received over the wire as this struct.
type Foo_Foo_Args struct {
}

// ToWire translates a Foo_Foo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Foo_Foo_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Foo_Foo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Foo_Foo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Foo_Foo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Foo_Foo_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Foo_Foo_Args
// struct.
func (v *Foo_Foo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Foo_Foo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Foo_Foo_Args match the
// provided Foo_Foo_Args.
//
// This function performs a deep comparison.
func (v *Foo_Foo_Args) Equals(rhs *Foo_Foo_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "foo" for this struct.
func (v *Foo_Foo_Args) MethodName() string {
	return "foo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Foo_Foo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Foo_Foo_Helper provides functions that aid in handling the
// parameters and return values of the Foo.foo
// function.
var Foo_Foo_Helper = struct {
	// Args accepts the parameters of foo in-order and returns
	// the arguments struct for the function.
	Args func() *Foo_Foo_Args

	// IsException returns true if the given error can be thrown
	// by foo.
	//
	// An error can be thrown by foo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for foo
	// given the error returned by it. The provided error may
	// be nil if foo did not fail.
	//
	// This allows mapping errors returned by foo into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// foo
	//
	//   err := foo(args)
	//   result, err := Foo_Foo_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from foo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Foo_Foo_Result, error)

	// UnwrapResponse takes the result struct for foo
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if foo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Foo_Foo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Foo_Foo_Result) error
}{}

func init() {
	Foo_Foo_Helper.Args = func() *Foo_Foo_Args {
		return &Foo_Foo_Args{}
	}

	Foo_Foo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Foo_Foo_Helper.WrapResponse = func(err error) (*Foo_Foo_Result, error) {
		if err == nil {
			return &Foo_Foo_Result{}, nil
		}

		return nil, err
	}
	Foo_Foo_Helper.UnwrapResponse = func(result *Foo_Foo_Result) (err error) {
		return
	}

}

// Foo_Foo_Result represents the result of a Foo.foo function call.
//
// The result of a foo execution is sent and received over the wire as this struct.
type Foo_Foo_Result struct {
}

// ToWire translates a Foo_Foo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Foo_Foo_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Foo_Foo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Foo_Foo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Foo_Foo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Foo_Foo_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Foo_Foo_Result
// struct.
func (v *Foo_Foo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Foo_Foo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Foo_Foo_Result match the
// provided Foo_Foo_Result.
//
// This function performs a deep comparison.
func (v *Foo_Foo_Result) Equals(rhs *Foo_Foo_Result) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "foo" for this struct.
func (v *Foo_Foo_Result) MethodName() string {
	return "foo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Foo_Foo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
