package rental

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"rental/testutil/sample"
	rentalsimulation "rental/x/rental/simulation"
	"rental/x/rental/types"
)

// avoid unused import issue
var (
	_ = rentalsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRequestRental = "op_weight_msg_request_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestRental int = 100

	opWeightMsgApproveRental = "op_weight_msg_approve_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveRental int = 100

	opWeightMsgRepayRental = "op_weight_msg_repay_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayRental int = 100

	opWeightMsgLiquidateRental = "op_weight_msg_liquidate_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateRental int = 100

	opWeightMsgCancelRental = "op_weight_msg_cancel_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelRental int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rentalGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rentalGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestRental int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestRental, &weightMsgRequestRental, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRental = defaultWeightMsgRequestRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestRental,
		rentalsimulation.SimulateMsgRequestRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveRental int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveRental, &weightMsgApproveRental, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRental = defaultWeightMsgApproveRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRental,
		rentalsimulation.SimulateMsgApproveRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayRental int
	simState.AppParams.GetOrGenerate(opWeightMsgRepayRental, &weightMsgRepayRental, nil,
		func(_ *rand.Rand) {
			weightMsgRepayRental = defaultWeightMsgRepayRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayRental,
		rentalsimulation.SimulateMsgRepayRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateRental int
	simState.AppParams.GetOrGenerate(opWeightMsgLiquidateRental, &weightMsgLiquidateRental, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateRental = defaultWeightMsgLiquidateRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateRental,
		rentalsimulation.SimulateMsgLiquidateRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelRental int
	simState.AppParams.GetOrGenerate(opWeightMsgCancelRental, &weightMsgCancelRental, nil,
		func(_ *rand.Rand) {
			weightMsgCancelRental = defaultWeightMsgCancelRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelRental,
		rentalsimulation.SimulateMsgCancelRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestRental,
			defaultWeightMsgRequestRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentalsimulation.SimulateMsgRequestRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveRental,
			defaultWeightMsgApproveRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentalsimulation.SimulateMsgApproveRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRepayRental,
			defaultWeightMsgRepayRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentalsimulation.SimulateMsgRepayRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLiquidateRental,
			defaultWeightMsgLiquidateRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentalsimulation.SimulateMsgLiquidateRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelRental,
			defaultWeightMsgCancelRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentalsimulation.SimulateMsgCancelRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
