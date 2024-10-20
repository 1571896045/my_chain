package types

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the mychain module.
type AppModuleBasic struct{}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx context.Context, mux *runtime.ServeMux) error {
	// 这里应该实现实际的路由注册逻辑
	// 例如：return types.RegisterQueryHandlerClient(clientCtx, mux, types.NewQueryClient(clientCtx))
	return nil
}

func (AppModuleBasic) Name() string {
	return ModuleName
}

// Module represents the AppModule for this module
type Module struct {
	Authority string
}

// QueryServer defines the gRPC query service for the module
type QueryServer interface {
	// Add query methods here
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// MsgServer defines the gRPC msg service for the module
type MsgServer interface {
	// Add message methods here
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

var (
	// ModuleName defines the module name
	ModuleName = "mychain"

	// QueryServiceDesc defines the gRPC query service descriptor
	QueryServiceDesc = grpc.ServiceDesc{
		ServiceName: "mychain.Query",
		HandlerType: (*QueryServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Params",
				Handler:    nil, // This will be set in the keeper
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "mychain/query.proto",
	}

	// MsgServiceDesc defines the gRPC msg service descriptor
	MsgServiceDesc = grpc.ServiceDesc{
		ServiceName: "mychain.Msg",
		HandlerType: (*MsgServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "UpdateParams",
				Handler:    nil, // This will be set in the keeper
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "mychain/tx.proto",
	}
)
