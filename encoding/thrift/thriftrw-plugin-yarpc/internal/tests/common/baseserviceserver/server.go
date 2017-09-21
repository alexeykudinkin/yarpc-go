// Code generated by thriftrw-plugin-yarpc
// @generated

package baseserviceserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common"
)

// Interface is the server-side interface for the BaseService service.
type Interface interface {
	Healthy(
		ctx context.Context,
	) (bool, error)
}

// New prepares an implementation of the BaseService service for
// registration.
//
// 	handler := BaseServiceHandler{}
// 	dispatcher.Register(baseserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "BaseService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "healthy",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Healthy),
				},
				Signature:    "Healthy() (bool)",
				ThriftModule: common.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) Healthy(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args common.BaseService_Healthy_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, appErr := h.impl.Healthy(ctx)

	result, err := common.BaseService_Healthy_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.Body = result
		response.ApplicationError = appErr
		response.IsApplicationError = appErr != nil
	}
	return response, err
}
