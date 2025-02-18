package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
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

		authority string

		Schema collections.Schema
		Params collections.Item[types.Params]
		// LockTransactionStore - key: lock transaction rawTx + vout | value: lock transaction data
		LockTransactionStore collections.Map[collections.Pair[string, uint64], types.LockTransaction]
		// PendingMintTransactions - key: pending zenBTC mint transaction
		PendingMintTransactions collections.Item[types.PendingMintTransactions] // DEPRECATED
		// PendingMintTransactionsMap - key: pending zenBTC mint transaction id | value: pending zenBTC mint transaction
		PendingMintTransactionsMap collections.Map[uint64, types.PendingMintTransaction]
		// FirstPendingMintTransaction - value: lowest key of pending mint transaction
		FirstPendingMintTransaction collections.Item[uint64]
		// PendingMintTransactionCount - value: count of pending zenBTC mint transactions
		PendingMintTransactionCount collections.Item[uint64]
		// BurnEvents - key: burn event index | value: burn event data
		BurnEvents collections.Map[uint64, types.BurnEvent]
		// FirstPendingBurnEvent - value: lowest key of pending burn event
		FirstPendingBurnEvent collections.Item[uint64]
		// BurnEventCount - value: count of burn events
		BurnEventCount collections.Item[uint64]
		// Redemptions - key: redemption index | value: redemption data
		Redemptions collections.Map[uint64, types.Redemption]
		// FirstPendingRedemption - value: lowest key of pending redemption
		FirstPendingRedemption collections.Item[uint64]
		// Supply - value: zenBTC supply data
		Supply collections.Item[types.Supply]
		// FirstPendingStakeTransaction - value: lowest key of pending stake transaction
		FirstPendingStakeTransaction collections.Item[uint64]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	validationKeeper *validation.Keeper,
	treasuryKeeper *treasury.Keeper,
) *Keeper {

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:                          cdc,
		storeService:                 storeService,
		logger:                       logger,
		validationKeeper:             validationKeeper,
		treasuryKeeper:               treasuryKeeper,
		authority:                    authority,
		Params:                       collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc)),
		LockTransactionStore:         collections.NewMap(sb, types.LockTransactionsKey, types.LockTransactionsIndex, collections.PairKeyCodec(collections.StringKey, collections.Uint64Key), codec.CollValue[types.LockTransaction](cdc)),
		PendingMintTransactions:      collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc)),
		PendingMintTransactionsMap:   collections.NewMap(sb, types.PendingMintTransactionsMapKey, types.PendingMintTransactionsMapIndex, collections.Uint64Key, codec.CollValue[types.PendingMintTransaction](cdc)),
		FirstPendingMintTransaction:  collections.NewItem(sb, types.FirstPendingMintTransactionKey, types.FirstPendingMintTransactionIndex, collections.Uint64Value),
		PendingMintTransactionCount:  collections.NewItem(sb, types.PendingMintTransactionCountKey, types.PendingMintTransactionCountIndex, collections.Uint64Value),
		BurnEvents:                   collections.NewMap(sb, types.BurnEventsKey, types.BurnEventsIndex, collections.Uint64Key, codec.CollValue[types.BurnEvent](cdc)),
		FirstPendingBurnEvent:        collections.NewItem(sb, types.FirstPendingBurnEventKey, types.FirstPendingBurnEventIndex, collections.Uint64Value),
		BurnEventCount:               collections.NewItem(sb, types.BurnEventCountKey, types.BurnEventCountIndex, collections.Uint64Value),
		Redemptions:                  collections.NewMap(sb, types.RedemptionsKey, types.RedemptionsIndex, collections.Uint64Key, codec.CollValue[types.Redemption](cdc)),
		FirstPendingRedemption:       collections.NewItem(sb, types.FirstPendingRedemptionKey, types.FirstPendingRedemptionIndex, collections.Uint64Value),
		Supply:                       collections.NewItem(sb, types.SupplyKey, types.SupplyIndex, codec.CollValue[types.Supply](cdc)),
		FirstPendingStakeTransaction: collections.NewItem(sb, types.FirstPendingStakeTransactionKey, types.FirstPendingStakeTransactionIndex, collections.Uint64Value),
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

// GetZenBTCExchangeRate returns the current exchange rate between BTC and zenBTC
// Returns the number of BTC represented by 1 zenBTC
func (k Keeper) GetExchangeRate(ctx context.Context) (math.LegacyDec, error) {
	supply, err := k.Supply.Get(ctx)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return math.LegacyNewDec(0), err
		}
		return math.LegacyNewDec(1), nil // Initial exchange rate of 1:1
	}

	totalZenBTC := supply.MintedZenBTC + supply.PendingZenBTC
	if totalZenBTC == 0 {
		return math.LegacyNewDec(1), nil // If no mints/deposits yet, use 1:1 rate
	}

	return math.LegacyNewDecFromInt(math.NewIntFromUint64(supply.CustodiedBTC)).Quo(math.LegacyNewDecFromInt(math.NewIntFromUint64(totalZenBTC))), nil
}

func (k Keeper) SetPendingMintTransaction(ctx context.Context, pendingMintTransaction types.PendingMintTransaction) error {
	return k.PendingMintTransactionsMap.Set(ctx, pendingMintTransaction.Id, pendingMintTransaction)
}

func (k Keeper) WalkPendingMintTransactions(ctx context.Context, fn func(id uint64, pendingMintTransaction types.PendingMintTransaction) (stop bool, err error)) error {
	return k.PendingMintTransactionsMap.Walk(ctx, nil, fn)
}

func (k Keeper) HasRedemption(ctx context.Context, id uint64) (bool, error) {
	return k.Redemptions.Has(ctx, id)
}

func (k Keeper) SetRedemption(ctx context.Context, id uint64, redemption types.Redemption) error {
	return k.Redemptions.Set(ctx, id, redemption)
}

func (k Keeper) WalkRedemptions(ctx context.Context, fn func(id uint64, redemption types.Redemption) (stop bool, err error)) error {
	return k.Redemptions.Walk(ctx, nil, fn)
}

func (k Keeper) GetSupply(ctx context.Context) (types.Supply, error) {
	return k.Supply.Get(ctx)
}

func (k Keeper) SetSupply(ctx context.Context, supply types.Supply) error {
	return k.Supply.Set(ctx, supply)
}

func (k Keeper) GetBurnEvent(ctx context.Context, id uint64) (types.BurnEvent, error) {
	return k.BurnEvents.Get(ctx, id)
}

func (k Keeper) SetBurnEvent(ctx context.Context, id uint64, burnEvent types.BurnEvent) error {
	return k.BurnEvents.Set(ctx, id, burnEvent)
}

func (k Keeper) WalkBurnEvents(ctx context.Context, fn func(id uint64, burnEvent types.BurnEvent) (stop bool, err error)) error {
	return k.BurnEvents.Walk(ctx, nil, fn)
}

func (k Keeper) GetPendingMintTransactionsStore() collections.Map[uint64, types.PendingMintTransaction] {
	return k.PendingMintTransactionsMap
}

func (k Keeper) GetBurnEventsStore() collections.Map[uint64, types.BurnEvent] {
	return k.BurnEvents
}

func (k Keeper) GetRedemptionsStore() collections.Map[uint64, types.Redemption] {
	return k.Redemptions
}

// GetFirstPendingMintTransaction returns the ID of the first pending mint transaction
func (k Keeper) GetFirstPendingMintTransaction(ctx context.Context) (uint64, error) {
	return k.FirstPendingMintTransaction.Get(ctx)
}

// SetFirstPendingMintTransaction sets the ID of the first pending mint transaction
func (k Keeper) SetFirstPendingMintTransaction(ctx context.Context, id uint64) error {
	return k.FirstPendingMintTransaction.Set(ctx, id)
}

// GetFirstPendingBurnEvent returns the ID of the first pending burn event
func (k Keeper) GetFirstPendingBurnEvent(ctx context.Context) (uint64, error) {
	return k.FirstPendingBurnEvent.Get(ctx)
}

// SetFirstPendingBurnEvent sets the ID of the first pending burn event
func (k Keeper) SetFirstPendingBurnEvent(ctx context.Context, id uint64) error {
	return k.FirstPendingBurnEvent.Set(ctx, id)
}

// GetFirstPendingRedemption returns the ID of the first pending redemption
func (k Keeper) GetFirstPendingRedemption(ctx context.Context) (uint64, error) {
	return k.FirstPendingRedemption.Get(ctx)
}

// SetFirstPendingRedemption sets the ID of the first pending redemption
func (k Keeper) SetFirstPendingRedemption(ctx context.Context, id uint64) error {
	return k.FirstPendingRedemption.Set(ctx, id)
}

// GetFirstPendingStakeTransaction returns the ID of the first pending stake transaction
func (k Keeper) GetFirstPendingStakeTransaction(ctx context.Context) (uint64, error) {
	return k.FirstPendingStakeTransaction.Get(ctx)
}

// SetFirstPendingStakeTransaction sets the ID of the first pending stake transaction
func (k Keeper) SetFirstPendingStakeTransaction(ctx context.Context, id uint64) error {
	return k.FirstPendingStakeTransaction.Set(ctx, id)
}

func (k Keeper) GetAuthority() string {
	return k.authority
}
