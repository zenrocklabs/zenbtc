package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "zenbtc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zenbtc"
)

var (
	ParamsKey                       = collections.NewPrefix(0)
	LockTransactionsKey             = collections.NewPrefix(1)
	PendingMintTransactionsKey      = collections.NewPrefix(2)
	RedemptionsKey                  = collections.NewPrefix(3)
	SupplyKey                       = collections.NewPrefix(4)
	PendingMintTransactionCountKey  = collections.NewPrefix(5)
	BurnEventsKey                   = collections.NewPrefix(6)
	PendingMintTransactionsMapKey   = collections.NewPrefix(7)
	BurnEventCountKey               = collections.NewPrefix(8)
	FirstPendingMintTransactionKey  = collections.NewPrefix(9)
	FirstPendingBurnEventKey        = collections.NewPrefix(10)
	FirstPendingRedemptionKey       = collections.NewPrefix(11)
	FirstPendingStakeTransactionKey = collections.NewPrefix(12)
	UTXOSpentKey                    = collections.NewPrefix(13)

	ParamsIndex                       = "params"
	LockTransactionsIndex             = "lock_transactions"
	PendingMintTransactionsIndex      = "pending_mint_transactions"
	RedemptionsIndex                  = "redemptions"
	SupplyIndex                       = "supply"
	PendingMintTransactionCountIndex  = "pending_mint_transaction_count"
	BurnEventsIndex                   = "burn_events"
	PendingMintTransactionsMapIndex   = "pending_mint_transactions_map"
	BurnEventCountIndex               = "burn_event_count"
	FirstPendingMintTransactionIndex  = "first_pending_mint_transaction"
	FirstPendingBurnEventIndex        = "first_pending_burn_event"
	FirstPendingRedemptionIndex       = "first_pending_redemption"
	FirstPendingStakeTransactionIndex = "first_pending_stake_transaction"
	UTXOSpentIndex                    = "redemption_utxo_usage"
)
