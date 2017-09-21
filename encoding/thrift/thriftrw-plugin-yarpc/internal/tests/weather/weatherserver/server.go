// Code generated by thriftrw-plugin-yarpc
// @generated

package weatherserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/weather"
)

// Interface is the server-side interface for the Weather service.
type Interface interface {
	Check(
		ctx context.Context,
	) (string, error)
}

// New prepares an implementation of the Weather service for
// registration.
//
// 	handler := WeatherHandler{}
// 	dispatcher.Register(weatherserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Weather",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "check",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Check),
				},
				Signature:    "Check() (string)",
				ThriftModule: weather.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) Check(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args weather.Weather_Check_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, appErr := h.impl.Check(ctx)

	result, err := weather.Weather_Check_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.Body = result
		response.ApplicationError = appErr
		response.IsApplicationError = appErr != nil
	}
	return response, err
}
