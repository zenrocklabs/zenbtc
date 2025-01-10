package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"

	treasury "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/keeper"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		validationKeeper *validation.Keeper
		treasuryKeeper   *treasury.Keeper

		Schema               collections.Schema
		Params               collections.Item[types.Params]
		LockTransactionStore collections.Map[collections.Pair[string, uint64], types.LockTransaction]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	validationKeeper *validation.Keeper,
	treasuryKeeper *treasury.Keeper,
) *Keeper {

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:              cdc,
		storeService:     storeService,
		logger:           logger,
		validationKeeper: validationKeeper,
		treasuryKeeper:   treasuryKeeper,

		Params:               collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc)),
		LockTransactionStore: collections.NewMap(sb, types.LockTransactionsKey, types.LockTransactionsIndex, collections.PairKeyCodec(collections.StringKey, collections.Uint64Key), codec.CollValue[types.LockTransaction](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return &k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
