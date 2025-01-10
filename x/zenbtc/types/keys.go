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
	ParamsKey                  = collections.NewPrefix(0)
	LockTransactionsKey        = collections.NewPrefix(1)
	PendingMintTransactionsKey = collections.NewPrefix(2)
	RedemptionsKey             = collections.NewPrefix(3)
	SupplyKey                  = collections.NewPrefix(4)

	ParamsIndex                  = "params"
	LockTransactionsIndex        = "lock_transactions"
	PendingMintTransactionsIndex = "pending_mint_transactions"
	RedemptionsIndex             = "redemptions"
	SupplyIndex                  = "supply"
)
