// Code generated by thriftrw-plugin-yarpc
// @generated

package bazclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends/barclient"
	"reflect"
)

// Interface is a client for the Baz service.
type Interface interface {
	barclient.Interface

	Baz(
		ctx context.Context,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the Baz service.
//
// 	client := bazclient.New(dispatcher.ClientConfig("baz"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Baz",
			ClientConfig: c,
		}, opts...),
		Interface: barclient.New(c, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	barclient.Interface

	c thrift.Client
}

func (c client) Baz(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (err error) {

	args := extends.Baz_Baz_Helper.Args()

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result extends.Baz_Baz_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = extends.Baz_Baz_Helper.UnwrapResponse(&result)
	return
}
