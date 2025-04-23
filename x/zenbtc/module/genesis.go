package zenbtc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/keeper"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		panic(err)
	}

	for _, url := range genState.NoFeeMsgs {
		// Set the disabled type urls
		if err := k.NoFeeMsgsList.Set(ctx, url); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := keeper.DefaultGenesis()
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Params = params

	if err := k.NoFeeMsgsList.Walk(ctx, nil, func(msgUrl string) (stop bool, err error) {
		genesis.NoFeeMsgs = append(genesis.NoFeeMsgs, msgUrl)
		return false, nil
	}); err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
