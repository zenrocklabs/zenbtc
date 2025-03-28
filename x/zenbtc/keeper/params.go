package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var (
	DefaultControllerAddr                    = "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9"
	DefaultEthTokenAddr                      = "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3"
	DefaultDepositKeyringAddr                = "keyring1k6vc6vhp6e6l3rxalue9v4ux"
	DefaultEthMinterKeyID             uint64 = 2
	DefaultChangeAddressKeyIDs               = []uint64{3}
	DefaultUnstakerKeyID              uint64 = 4
	DefaultRewardsDepositKeyID        uint64 = 5
	DefaultStakerKeyID                uint64 = 6
	DefaultCompleterKeyID             uint64 = 7
	DefaultTestnetBitcoinProxyAddress        = "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty"
	DefaultMainnetBitcoinProxyAddress        = "zen1mgl98jt30nemuqtt5asldk49ju9lnx0pfke79q"
	// DefaultStrategyAddr               = "0x0000000000000000000000000000000000000000"
	// DefaultStakerKeyID = 0
	// DefaultBurnerKeyID = 0
	DefaultSolana = &types.Solana{
		SignerKeyId:       1,
		ProgramId:         "zTsNe6inwEA8Z6WdogiA1vPLLAkaheGdU6WNvxdaXUx",
		NonceAccountKey:   11,
		NonceAuthorityKey: 12,
		MintAddress:       "Ce1gvD5BG1gjEjuDKJXPcQgNfwDPJH9VvgfcSEzyNZn4",
		FeeWallet:         "FzqGcRG98v1KhKxatX2Abb2z1aJ2rViQwBK5GHByKCAd",
		Fee:               0,
	}
)

// NewParams creates a new Params instance
func NewParams(
	depositKeyringAddr string,
	stakerKeyID,
	ethMinterKeyID,
	unstakerKeyID,
	completerKeyID,
	rewardsDepositKeyID uint64,
	changeAddressKeyIDs []uint64,
	bitcoinProxyAddress,
	ethTokenAddr,
	controllerAddr string,
	solana *types.Solana,
) *types.Params {
	return &types.Params{
		DepositKeyringAddr:  depositKeyringAddr,
		StakerKeyID:         stakerKeyID,
		EthMinterKeyID:      ethMinterKeyID,
		Solana:              solana,
		UnstakerKeyID:       unstakerKeyID,
		CompleterKeyID:      completerKeyID,
		RewardsDepositKeyID: rewardsDepositKeyID,
		ChangeAddressKeyIDs: changeAddressKeyIDs,
		BitcoinProxyAddress: bitcoinProxyAddress,
		EthTokenAddr:        ethTokenAddr,
		ControllerAddr:      controllerAddr,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() *types.Params {
	return NewParams(
		DefaultDepositKeyringAddr,
		DefaultStakerKeyID,
		DefaultEthMinterKeyID,
		DefaultUnstakerKeyID,
		DefaultCompleterKeyID,
		DefaultRewardsDepositKeyID,
		DefaultChangeAddressKeyIDs,
		DefaultTestnetBitcoinProxyAddress,
		DefaultEthTokenAddr,
		DefaultControllerAddr,
		DefaultSolana,
	)
}

func (k Keeper) GetControllerAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultControllerAddr
	}
	return params.ControllerAddr
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

func (k Keeper) GetSolanaParams(ctx context.Context) *types.Solana {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultSolana
	}
	return params.Solana
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

// func (k Keeper) GetBitcoinProxyAddress(ctx context.Context) string {
// TODO: why does the commented line below panic?
// params, err := k.Params.Get(ctx)
// if err != nil {
// return DefaultProxyAddress
// }
// return params.BitcoinProxyAddress
// }

func (k Keeper) GetBitcoinProxyAddress(ctx context.Context) string {
	return GetDefaultBitcoinProxyAddress(ctx)
}

func GetDefaultBitcoinProxyAddress(ctx context.Context) string {
	if strings.HasPrefix(sdk.UnwrapSDKContext(ctx).ChainID(), "diamond") {
		return DefaultMainnetBitcoinProxyAddress
	}
	return DefaultTestnetBitcoinProxyAddress
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
