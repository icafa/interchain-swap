package interquery

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/junkai121/interchain-swap/testutil/sample"
	interquerysimulation "github.com/junkai121/interchain-swap/x/interquery/simulation"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = interquerysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSendQueryOsmosisPrice = "op_weight_msg_send_query_all_balances"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendQueryOsmosisPrice int = 100

	opWeightMsgSendOsmosisSwap = "op_weight_msg_send_osmosis_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendOsmosisSwap int = 100

	opWeightMsgRegisterICA = "op_weight_msg_register_ica"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterICA int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	interqueryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&interqueryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSendQueryOsmosisPrice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendQueryOsmosisPrice, &weightMsgSendQueryOsmosisPrice, nil,
		func(_ *rand.Rand) {
			weightMsgSendQueryOsmosisPrice = defaultWeightMsgSendQueryOsmosisPrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendQueryOsmosisPrice,
		interquerysimulation.SimulateMsgSendQueryOsmosisPrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendOsmosisSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendOsmosisSwap, &weightMsgSendOsmosisSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSendOsmosisSwap = defaultWeightMsgSendOsmosisSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendOsmosisSwap,
		interquerysimulation.SimulateMsgSendOsmosisSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterICA int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterICA, &weightMsgRegisterICA, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterICA = defaultWeightMsgRegisterICA
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterICA,
		interquerysimulation.SimulateMsgRegisterICA(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
