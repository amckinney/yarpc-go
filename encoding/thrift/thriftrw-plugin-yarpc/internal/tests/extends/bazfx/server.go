// Code generated by thriftrw-plugin-yarpc
// @generated

package bazfx

import (
	"go.uber.org/fx"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends/bazserver"
)

// ServerParams defines the dependencies for the Baz server.
type ServerParams struct {
	fx.In

	Handler bazserver.Interface
}

// ServerResult defines the output of Baz server module. It provides the
// procedures of a Baz handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// Server provides procedures for Baz to an Fx application. It expects a
// bazfx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyBazHandler) bazserver.Interface {
// 			return h
// 		},
// 		bazfx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := bazserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}
