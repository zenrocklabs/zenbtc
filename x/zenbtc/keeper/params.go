package keeper

import (
	"context"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var (
	DefaultEthBatcherAddr             = "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA"
	DefaultDepositKeyringAddr         = "keyring1k6vc6vhp6e6l3rxalue9v4ux"
	DefaultWithdrawerKeyID     uint64 = 1
	DefaultMinterKeyID         uint64 = 2
	DefaultChangeAddressKeyIDs        = []uint64{3}
	DefaultUnstakerKeyID       uint64 = 4
	DefaultRewardsDepositKeyID uint64 = 5
	DefaultProxyAddress               = "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty"
	DefaultParamsAuthority            = "zen1sd3fwcpw2mdw3pxexmlg34gsd78r0sxrk5weh3"
	// DefaultStrategyAddr               = "0x0000000000000000000000000000000000000000"
	// DefaultStakerKeyID = 0
	// DefaultBurnerKeyID = 0
)

// NewParams creates a new Params instance
func NewParams(ethBatcherAddr, depositKeyringAddr string, minterKeyID, unstakerKeyID, withdrawerKeyID uint64, changeAddressKeyIDs []uint64, rewardsDepositKeyID uint64, bitcoinProxyAddress, authority string) *types.Params {
	return &types.Params{
		EthBatcherAddr:      ethBatcherAddr,
		DepositKeyringAddr:  depositKeyringAddr,
		MinterKeyID:         minterKeyID,
		UnstakerKeyID:       unstakerKeyID,
		WithdrawerKeyID:     withdrawerKeyID,
		ChangeAddressKeyIDs: changeAddressKeyIDs,
		RewardsDepositKeyID: rewardsDepositKeyID,
		BitcoinProxyAddress: bitcoinProxyAddress,
		Authority:           authority,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() *types.Params {
	return NewParams(
		DefaultEthBatcherAddr,
		DefaultDepositKeyringAddr,
		DefaultMinterKeyID,
		DefaultUnstakerKeyID,
		DefaultWithdrawerKeyID,
		DefaultChangeAddressKeyIDs,
		DefaultRewardsDepositKeyID,
		DefaultProxyAddress,
		DefaultParamsAuthority,
	)
}

func (k Keeper) GetEthBatcherAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultEthBatcherAddr
	}
	return params.EthBatcherAddr
}

func (k Keeper) GetDepositKeyringAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultDepositKeyringAddr
	}
	return params.DepositKeyringAddr
}

func (k Keeper) GetMinterKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultMinterKeyID
	}
	return params.MinterKeyID
}

func (k Keeper) GetUnstakerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultUnstakerKeyID
	}
	return params.UnstakerKeyID
}

func (k Keeper) GetWithdrawerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultWithdrawerKeyID
	}
	return params.WithdrawerKeyID
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

func (k Keeper) GetParamsAuthority(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultParamsAuthority
	}
	return params.Authority
}
