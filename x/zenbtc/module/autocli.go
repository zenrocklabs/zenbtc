package zenbtc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/zenrocklabs/zenbtc/api/zrchain/zenbtc"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetParams",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "LockTransactions",
					Use:            "lock-transactions",
					Short:          "Query LockTransactions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetRedemptions",
					Use:            "redemptions",
					Short:          "Query Redemptions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "QueryPendingMintTransactions",
					Use:            "pending-mint-transactions",
					Short:          "Query PendingMintTransactions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "QuerySupply",
					Use:            "supply",
					Short:          "Query Supply",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod: "VerifyDepositBlockInclusion",
					Use:       "verify-block-inclusion [chain_name] [block_height] [raw_tx] [index] [proof] [deposit_addr] [amount]",
					Short:     "Send a VerifyDepositBlockInclusion tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "chain_name"},
						{ProtoField: "block_height"},
						{ProtoField: "raw_tx"},
						{ProtoField: "index"},
						{ProtoField: "proof"},
						{ProtoField: "deposit_addr"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod:      "SubmitUnsignedRedemptionTx",
					Use:            "submit-unsigned-redemption-tx [outputs]",
					Short:          "Send a SubmitUnsignedRedemptionTx tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						//{ProtoField: "outputs"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
