package keeper

import (
	"context"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var (
	DefaultEthBatcherAddr             = "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9"
	DefaultEthTokenAddr               = "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3"
	DefaultDepositKeyringAddr         = "keyring1k6vc6vhp6e6l3rxalue9v4ux"
	DefaultEthMinterKeyID      uint64 = 2
	DefaultChangeAddressKeyIDs        = []uint64{3}
	DefaultUnstakerKeyID       uint64 = 4
	DefaultRewardsDepositKeyID uint64 = 5
	DefaultStakerKeyID         uint64 = 6
	DefaultCompleterKeyID      uint64 = 7
	DefaultProxyAddress               = "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty"
	// DefaultStrategyAddr               = "0x0000000000000000000000000000000000000000"
	// DefaultStakerKeyID = 0
	// DefaultBurnerKeyID = 0
)

// NewParams creates a new Params instance
func NewParams(ethBatcherAddr, ethTokenAddr, depositKeyringAddr string, stakerKeyID, ethMinterKeyID, unstakerKeyID, completerKeyID uint64, changeAddressKeyIDs []uint64, rewardsDepositKeyID uint64, bitcoinProxyAddress string) *types.Params {
	return &types.Params{
		EthBatcherAddr:      ethBatcherAddr,
		EthTokenAddr:        ethTokenAddr,
		DepositKeyringAddr:  depositKeyringAddr,
		StakerKeyID:         stakerKeyID,
		EthMinterKeyID:      ethMinterKeyID,
		UnstakerKeyID:       unstakerKeyID,
		CompleterKeyID:      completerKeyID,
		ChangeAddressKeyIDs: changeAddressKeyIDs,
		RewardsDepositKeyID: rewardsDepositKeyID,
		BitcoinProxyAddress: bitcoinProxyAddress,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() *types.Params {
	return NewParams(
		DefaultEthBatcherAddr,
		DefaultEthTokenAddr,
		DefaultDepositKeyringAddr,
		DefaultStakerKeyID,
		DefaultEthMinterKeyID,
		DefaultUnstakerKeyID,
		DefaultCompleterKeyID,
		DefaultChangeAddressKeyIDs,
		DefaultRewardsDepositKeyID,
		DefaultProxyAddress,
	)
}

func (k Keeper) GetEthBatcherAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultEthBatcherAddr
	}
	return params.EthBatcherAddr
}

func (k Keeper) GetEthTokenAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultEthTokenAddr
	}
	return params.EthTokenAddr
}

func (k Keeper) GetDepositKeyringAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultDepositKeyringAddr
	}
	return params.DepositKeyringAddr
}

func (k Keeper) GetStakerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultStakerKeyID
	}
	return params.StakerKeyID
}

func (k Keeper) GetEthMinterKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultEthMinterKeyID
	}
	return params.EthMinterKeyID
}

func (k Keeper) GetUnstakerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultUnstakerKeyID
	}
	return params.UnstakerKeyID
}

func (k Keeper) GetCompleterKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultCompleterKeyID
	}
	return params.CompleterKeyID
}

func (k Keeper) GetBitcoinProxyAddress(ctx context.Context) string {
	// params, err := k.Params.Get(ctx)
	// if err != nil {
	// 	return DefaultProxyAddress
	// }
	// return params.BitcoinProxyAddress

	// TODO: fix above block

	return DefaultProxyAddress
}

func (k Keeper) GetChangeAddressKeyIDs(ctx context.Context) []uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultChangeAddressKeyIDs
	}
	return params.ChangeAddressKeyIDs
}

func (k Keeper) GetRewardsDepositKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultRewardsDepositKeyID
	}
	return params.RewardsDepositKeyID
}
