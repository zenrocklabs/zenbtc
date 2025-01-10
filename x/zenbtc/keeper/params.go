package keeper

import "context"

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

func (k Keeper) GetZenBTCEthBatcherAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultEthBatcherAddr
	}
	return params.EthBatcherAddr
}

func (k Keeper) GetZenBTCDepositKeyringAddr(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultDepositKeyringAddr
	}
	return params.DepositKeyringAddr
}

func (k Keeper) GetZenBTCMinterKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultMinterKeyID
	}
	return params.MinterKeyID
}

func (k Keeper) GetZenBTCUnstakerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultUnstakerKeyID
	}
	return params.UnstakerKeyID
}

func (k Keeper) GetZenBTCWithdrawerKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultWithdrawerKeyID
	}
	return params.WithdrawerKeyID
}

func (k Keeper) GetBitcoinProxyAddress(ctx context.Context) string {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultProxyAddress
	}
	return params.BitcoinProxyAddress
}

func (k Keeper) GetZenBTCChangeAddressKeyIDs(ctx context.Context) []uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultChangeAddressKeyIDs
	}
	return params.ChangeAddressKeyIDs
}

func (k Keeper) GetZenBTCRewardsDepositKeyID(ctx context.Context) uint64 {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return DefaultRewardsDepositKeyID
	}
	return params.RewardsDepositKeyID
}
