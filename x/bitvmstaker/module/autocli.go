package bitvmstaker

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "fiamma/api/fiamma/bitvmstaker"
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
					RpcMethod:      "ListStakerAddresses",
					Use:            "list-staker-addresses",
					Short:          "Query list-staker-addresses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "CommitteeAddress",
					Use:            "committee-address",
					Short:          "Query committee-address",
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
					RpcMethod:      "CreateStaker",
					Use:            "create-staker [staker-address]",
					Short:          "Send a create-staker tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stakerAddress"}},
				},
				{
					RpcMethod:      "SlashStaker",
					Use:            "slash-staker [staker-address]",
					Short:          "Send a slash-staker tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "stakerAddress"}},
				},
				{
					RpcMethod:      "UpdateCommitteeAddress",
					Use:            "update-committee-address [new-committee-address]",
					Short:          "Send a update-committee-address tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "newCommitteeAddress"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
