package rental

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "rental/api/rental/rental"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "RentalAll",
					Use:       "list-rental",
					Short:     "List all rental",
				},
				{
					RpcMethod:      "Rental",
					Use:            "show-rental [id]",
					Short:          "Shows a rental by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RequestRental",
					Use:            "request-rental [amount] [fee] [asset] [due-date]",
					Short:          "Send a request-rental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "fee"}, {ProtoField: "asset"}, {ProtoField: "dueDate"}},
				},
				{
					RpcMethod:      "ApproveRental",
					Use:            "approve-rental [id]",
					Short:          "Send a approve-rental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "RepayRental",
					Use:            "repay-rental [id]",
					Short:          "Send a repay-rental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "LiquidateRental",
					Use:            "liquidate-rental [id]",
					Short:          "Send a liquidate-rental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CancelRental",
					Use:            "cancel-rental [id]",
					Short:          "Send a cancel-rental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
