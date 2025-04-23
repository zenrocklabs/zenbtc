package keeper

import "github.com/zenrocklabs/zenbtc/x/zenbtc/types"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *types.GenesisState {
	return &types.GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:    *DefaultParams(),
		NoFeeMsgs: DefaultNoFeeMsgs(),
	}
}
