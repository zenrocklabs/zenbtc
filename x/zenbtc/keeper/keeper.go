package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"

	treasury "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/keeper"
	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		validationKeeper *validation.Keeper
		treasuryKeeper   *treasury.Keeper

		Schema collections.Schema
		Params collections.Item[types.Params]
		// LockTransactionStore - key: lock transaction rawTx + vout | value: lock transaction data
		LockTransactionStore collections.Map[collections.Pair[string, uint64], types.LockTransaction]
		// PendingMintTransactions - key: pending zenBTC mint transaction
		PendingMintTransactions collections.Item[treasurytypes.PendingMintTransactions]
		// ZenBTCRedemptions - key: redemption index | value: redemption data
		Redemptions collections.Map[uint64, types.Redemption]
		// ZenBTCSupply - value: zenBTC supply data
		Supply collections.Item[types.Supply]
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

		Params:                  collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc)),
		LockTransactionStore:    collections.NewMap(sb, types.LockTransactionsKey, types.LockTransactionsIndex, collections.PairKeyCodec(collections.StringKey, collections.Uint64Key), codec.CollValue[types.LockTransaction](cdc)),
		PendingMintTransactions: collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[treasurytypes.PendingMintTransactions](cdc)),
		Redemptions:             collections.NewMap(sb, types.RedemptionsKey, types.RedemptionsIndex, collections.Uint64Key, codec.CollValue[types.Redemption](cdc)),
		Supply:                  collections.NewItem(sb, types.SupplyKey, types.SupplyIndex, codec.CollValue[types.Supply](cdc)),
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
